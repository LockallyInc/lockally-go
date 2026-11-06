package lockallyx

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func newFastClient(o *Options) *http.Client {
	if o == nil {
		o = &Options{}
	}
	o.BaseDelay = time.Millisecond // keep tests fast
	o.MaxDelay = 5 * time.Millisecond
	return NewClient(o)
}

func TestRetriesOn429ThenSucceeds(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n < 3 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer srv.Close()

	resp, err := newFastClient(nil).Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("want 200, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&calls); got != 3 {
		t.Fatalf("want 3 attempts, got %d", got)
	}
}

func TestGivesUpAfterMaxRetries(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	resp, err := newFastClient(&Options{MaxRetries: 2}).Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 503 {
		t.Fatalf("want 503, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&calls); got != 3 { // 1 initial + 2 retries
		t.Fatalf("want 3 attempts, got %d", got)
	}
}

func TestNoRetryOn4xx(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer srv.Close()

	_, err := newFastClient(nil).Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("400 must not retry; got %d attempts", got)
	}
}

func TestIdempotencyKeyInjectedAndStableAcrossRetries(t *testing.T) {
	var keys []string
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Idempotency-Key"))
		if atomic.AddInt32(&calls, 1) < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	resp, err := newFastClient(nil).Post(srv.URL, "application/json", strings.NewReader(`{"a":1}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = resp.Body.Close()
	if len(keys) != 2 {
		t.Fatalf("want 2 attempts, got %d (keys=%v)", len(keys), keys)
	}
	if keys[0] == "" {
		t.Fatal("Idempotency-Key was not injected on POST")
	}
	if keys[0] != keys[1] {
		t.Fatalf("Idempotency-Key changed across retries: %q vs %q", keys[0], keys[1])
	}
}

func TestGetHasNoIdempotencyKey(t *testing.T) {
	var key string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key = r.Header.Get("Idempotency-Key")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	_, err := newFastClient(nil).Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if key != "" {
		t.Fatalf("GET must not get an Idempotency-Key, got %q", key)
	}
}

func TestCallerIdempotencyKeyPreserved(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("Idempotency-Key")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodPost, srv.URL, strings.NewReader("{}"))
	req.Header.Set("Idempotency-Key", "caller-supplied")
	resp, err := newFastClient(nil).Do(req)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	if got != "caller-supplied" {
		t.Fatalf("caller key must be preserved, got %q", got)
	}
}

func TestContextCancelStopsRetry(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60") // long, so cancel wins
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	// Cap high enough that the Retry-After-driven backoff outlasts the ctx
	// timeout, so cancellation wins during the sleep.
	client := NewClient(&Options{BaseDelay: time.Second, MaxDelay: time.Second})
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	_, err := client.Do(req)
	if err == nil {
		t.Fatal("expected context deadline error")
	}
}

func TestParseRetryAfter(t *testing.T) {
	if d, ok := parseRetryAfter("2"); !ok || d != 2*time.Second {
		t.Fatalf("seconds parse failed: %v %v", d, ok)
	}
	if _, ok := parseRetryAfter(""); ok {
		t.Fatal("empty should not parse")
	}
	future := time.Now().Add(3 * time.Second).UTC().Format(http.TimeFormat)
	if d, ok := parseRetryAfter(future); !ok || d <= 0 {
		t.Fatalf("http-date parse failed: %v %v", d, ok)
	}
}

func TestPaginate(t *testing.T) {
	pages := map[string]Page[int]{
		"":   {Items: []int{1, 2}, NextCursor: "c1"},
		"c1": {Items: []int{3, 4}, NextCursor: "c2"},
		"c2": {Items: []int{5}, NextCursor: ""},
	}
	got, err := Collect(context.Background(), func(_ context.Context, cursor string) (Page[int], error) {
		return pages[cursor], nil
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("want %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}

func TestPaginateEarlyStopAndLoopGuard(t *testing.T) {
	// yield stops after 2 items
	var count int
	err := Paginate(context.Background(),
		func(_ context.Context, cursor string) (Page[int], error) {
			return Page[int]{Items: []int{1, 2, 3}, NextCursor: "same"}, nil // would loop forever
		},
		func(int) bool { count++; return count < 2 },
	)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("want early stop at 2, got %d", count)
	}
}
