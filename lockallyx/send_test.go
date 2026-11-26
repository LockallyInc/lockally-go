package lockallyx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func stubServer(t *testing.T, status int, payload any, seen *http.Request, body *map[string]any) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if seen != nil {
			*seen = *r.Clone(r.Context())
		}
		if body != nil {
			_ = json.NewDecoder(r.Body).Decode(body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-Id", "req-1")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(payload)
	}))
}

func TestSend_PostsWithKeyAndIdempotencyHeader(t *testing.T) {
	var seen http.Request
	srv := stubServer(t, 202, map[string]any{"id": "m1", "message_id": "<m1@test.lockally.com>", "status": "delivered"}, &seen, nil)
	defer srv.Close()

	c, err := New("lk_test_abc", WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.Send(context.Background(),
		Message{From: "hello@acme.test", To: []string{"user@example.com"}, Subject: "Hi", Text: "hello"},
		WithIdempotencyKey("welcome-user-42"))
	if err != nil {
		t.Fatalf("send: %v", err)
	}
	if seen.URL.Path != "/v1/send" {
		t.Errorf("path = %q, want /v1/send", seen.URL.Path)
	}
	if got := seen.Header.Get("Authorization"); got != "Bearer lk_test_abc" {
		t.Errorf("authorization = %q", got)
	}
	if got := seen.Header.Get("Idempotency-Key"); got != "welcome-user-42" {
		t.Errorf("idempotency key = %q", got)
	}
	if res.Status != "delivered" || res.ID != "m1" {
		t.Errorf("result = %+v", res)
	}
}

// Required by the API, so omitting it has to still work. The doc comment says
// why relying on the generated one is wrong; this pins that it IS sent.
func TestSend_OmittedKeyIsGenerated(t *testing.T) {
	var seen http.Request
	srv := stubServer(t, 202, map[string]any{"id": "m", "status": "queued"}, &seen, nil)
	defer srv.Close()

	c, _ := New("k", WithBaseURL(srv.URL))
	if _, err := c.Send(context.Background(), Message{From: "a@b.test", To: []string{"c@d.test"}, Text: "x"}); err != nil {
		t.Fatal(err)
	}
	if got := seen.Header.Get("Idempotency-Key"); len(got) < 16 {
		t.Errorf("expected a generated key, got %q", got)
	}
}

// An unset field must be absent from the wire, not sent as null: the API
// distinguishes the two on several fields.
func TestSend_UnsetFieldsAreOmitted(t *testing.T) {
	var body map[string]any
	srv := stubServer(t, 202, map[string]any{"id": "m", "status": "queued"}, nil, &body)
	defer srv.Close()

	c, _ := New("k", WithBaseURL(srv.URL))
	if _, err := c.Send(context.Background(), Message{From: "a@b.test", To: []string{"c@d.test"}, Text: "x"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := body["html"]; ok {
		t.Error("html must be omitted when unset")
	}
	if _, ok := body["attachments"]; ok {
		t.Error("attachments must be omitted when unset")
	}
	if body["from"] != "a@b.test" {
		t.Errorf("from = %v", body["from"])
	}
}

func TestSend_200IsASuccess(t *testing.T) {
	srv := stubServer(t, 200, map[string]any{"status": "suppressed", "warning": "All recipients are on your suppression list; nothing was sent."}, nil, nil)
	defer srv.Close()

	c, _ := New("k", WithBaseURL(srv.URL))
	res, err := c.Send(context.Background(), Message{From: "a@b.test", To: []string{"c@d.test"}, Text: "x"})
	if err != nil {
		t.Fatalf("a 200 is a success, got error: %v", err)
	}
	if res.Status != "suppressed" {
		t.Errorf("status = %q", res.Status)
	}
}

func TestSend_ErrorCarriesStatusCodeAndRequestID(t *testing.T) {
	srv := stubServer(t, 403, map[string]any{"title": "Sender not authorised", "code": "sender_not_authorised", "request_id": "req-1"}, nil, nil)
	defer srv.Close()

	c, _ := New("k", WithBaseURL(srv.URL))
	_, err := c.Send(context.Background(), Message{From: "a@b.test", To: []string{"c@d.test"}, Text: "x"})
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 403 || apiErr.Code != "sender_not_authorised" || apiErr.RequestID != "req-1" {
		t.Errorf("error = %+v", apiErr)
	}
	if apiErr.Retryable() {
		t.Error("a 403 is not retryable")
	}
}

func TestNew_RefusesAnEmptyKey(t *testing.T) {
	if _, err := New(""); err == nil {
		t.Error("an empty API key must be refused up front")
	}
}

func TestSend_FromNameGoesOutAsFromName(t *testing.T) {
	var body map[string]any
	srv := stubServer(t, 202, map[string]any{"id": "m", "status": "queued"}, nil, &body)
	defer srv.Close()

	c, _ := New("k", WithBaseURL(srv.URL))
	if _, err := c.Send(context.Background(), Message{From: "orders@acme.test", FromName: "Acme Orders", To: []string{"c@d.test"}, Text: "x"}); err != nil {
		t.Fatal(err)
	}
	if body["from"] != "orders@acme.test" || body["from_name"] != "Acme Orders" {
		t.Errorf("from = %v, from_name = %v", body["from"], body["from_name"])
	}

	// Unset, it is omitted like every other optional field. A fresh map, since
	// the stub decodes into whatever body already holds.
	body = nil
	if _, err := c.Send(context.Background(), Message{From: "a@b.test", To: []string{"c@d.test"}, Text: "x"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := body["from_name"]; ok {
		t.Error("from_name must be omitted when unset")
	}
}
