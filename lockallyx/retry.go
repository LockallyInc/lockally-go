package lockallyx

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math"
	mathrand "math/rand"
	"net/http"
	"strconv"
	"time"
)

// Defaults for the retrying transport.
const (
	DefaultMaxRetries = 3
	DefaultBaseDelay  = 200 * time.Millisecond
	DefaultMaxDelay   = 20 * time.Second
)

// idempotentMethods get an auto-generated Idempotency-Key (so a retried write
// is not applied twice). GET/HEAD are safe and need no key.
var mutatingMethods = map[string]bool{
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodPatch:  true,
	http.MethodDelete: true,
}

// Options configure the retrying transport.
type Options struct {
	MaxRetries int           // additional attempts after the first (default 3)
	BaseDelay  time.Duration // initial backoff (default 200ms)
	MaxDelay   time.Duration // backoff cap (default 20s)
	Base       http.RoundTripper
	// DisableAutoIdempotencyKey turns OFF the automatic Idempotency-Key header on
	// mutating requests that don't already carry one. The zero value keeps it ON.
	DisableAutoIdempotencyKey bool
}

func (o *Options) withDefaults() {
	if o.MaxRetries == 0 {
		o.MaxRetries = DefaultMaxRetries
	}
	if o.BaseDelay == 0 {
		o.BaseDelay = DefaultBaseDelay
	}
	if o.MaxDelay == 0 {
		o.MaxDelay = DefaultMaxDelay
	}
	if o.Base == nil {
		o.Base = http.DefaultTransport
	}
}

// Transport is an http.RoundTripper that retries transient failures with
// exponential backoff + jitter, honors Retry-After (429/503), and injects an
// Idempotency-Key on mutating requests.
type Transport struct {
	opts Options
}

// NewTransport builds a retrying RoundTripper. A nil *Options uses defaults.
func NewTransport(o *Options) *Transport {
	var opts Options
	if o != nil {
		opts = *o
	}
	opts.withDefaults()
	return &Transport{opts: opts}
}

// NewClient returns an *http.Client using the retrying transport, suitable for
// the generated Configuration.HTTPClient. A nil *Options uses defaults.
func NewClient(o *Options) *http.Client {
	return &http.Client{Transport: NewTransport(o)}
}

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the body once so it can be replayed on each attempt.
	var bodyBytes []byte
	if req.Body != nil && req.Body != http.NoBody {
		b, err := io.ReadAll(req.Body)
		_ = req.Body.Close()
		if err != nil {
			return nil, err
		}
		bodyBytes = b
	}

	// Generate one idempotency key up front so it is stable across retries.
	idemKey := req.Header.Get("Idempotency-Key")
	if !t.opts.DisableAutoIdempotencyKey && idemKey == "" && mutatingMethods[req.Method] {
		idemKey = newUUID()
	}

	ctx := req.Context()
	var lastErr error

	for attempt := 0; ; attempt++ {
		// Clone the request per attempt (RoundTrippers must not mutate the input).
		r := req.Clone(ctx)
		if bodyBytes != nil {
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			r.ContentLength = int64(len(bodyBytes))
			r.GetBody = func() (io.ReadCloser, error) {
				return io.NopCloser(bytes.NewReader(bodyBytes)), nil
			}
		}
		if idemKey != "" {
			r.Header.Set("Idempotency-Key", idemKey)
		}

		resp, err := t.opts.Base.RoundTrip(r)

		// Decide whether to retry.
		retryable := false
		var wait time.Duration
		if err != nil {
			retryable = true
			wait = t.backoff(attempt)
			lastErr = err
		} else if isRetryableStatus(resp.StatusCode) {
			retryable = true
			wait = t.retryDelay(attempt, resp)
			lastErr = &APIError{StatusCode: resp.StatusCode}
		}

		if !retryable || attempt >= t.opts.MaxRetries {
			return resp, err
		}

		// Drain+close the discarded response so the connection can be reused.
		if resp != nil {
			_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4<<10))
			_ = resp.Body.Close()
		}

		// Sleep, respecting context cancellation.
		select {
		case <-ctx.Done():
			if lastErr == nil {
				lastErr = ctx.Err()
			}
			return nil, ctx.Err()
		case <-time.After(wait):
		}
	}
}

func isRetryableStatus(code int) bool {
	switch code {
	case http.StatusTooManyRequests, // 429 — rate limited
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	}
	return false
}

// retryDelay honors Retry-After for 429/503, else falls back to backoff.
func (t *Transport) retryDelay(attempt int, resp *http.Response) time.Duration {
	if resp != nil {
		if d, ok := parseRetryAfter(resp.Header.Get("Retry-After")); ok {
			if d > t.opts.MaxDelay {
				return t.opts.MaxDelay
			}
			return d
		}
	}
	return t.backoff(attempt)
}

// backoff = min(base * 2^attempt, max) with full jitter.
func (t *Transport) backoff(attempt int) time.Duration {
	exp := float64(t.opts.BaseDelay) * math.Pow(2, float64(attempt))
	if exp > float64(t.opts.MaxDelay) {
		exp = float64(t.opts.MaxDelay)
	}
	// Full jitter: random in [0, exp].
	return time.Duration(mathrand.Int63n(int64(exp) + 1))
}

// parseRetryAfter handles both "delta-seconds" and HTTP-date forms.
func parseRetryAfter(v string) (time.Duration, bool) {
	if v == "" {
		return 0, false
	}
	if secs, err := strconv.Atoi(v); err == nil {
		if secs < 0 {
			secs = 0
		}
		return time.Duration(secs) * time.Second, true
	}
	if tm, err := http.ParseTime(v); err == nil {
		d := time.Until(tm)
		if d < 0 {
			d = 0
		}
		return d, true
	}
	return 0, false
}

// newUUID returns a RFC 4122 v4 UUID string (crypto/rand, no dependencies).
func newUUID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
