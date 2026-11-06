// Package lockallyx is the hand-written ergonomics layer for the generated
// Lockally Go SDK (github.com/lockallyinc/lockally-go). It works at the HTTP
// transport level, so it is independent of the generated code and survives
// regeneration: retries with backoff, 429/rate-limit handling, automatic
// idempotency keys, typed errors, and cursor pagination helpers.
//
// Wire it into the generated client via its Configuration.HTTPClient:
//
//	cfg := lockally.NewConfiguration()
//	cfg.HTTPClient = lockallyx.NewClient(nil) // retry+idempotency transport
//	client := lockally.NewAPIClient(cfg)
package lockallyx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIError is a typed error mapped from a non-2xx Lockally API response.
// The generated client returns its own error type; use FromResponse when you
// want the structured fields (status, code, request id) in your own flows.
type APIError struct {
	StatusCode int    // HTTP status
	Code       string // machine-readable error code from the body, if present
	Message    string // human-readable message from the body, if present
	RequestID  string // X-Request-Id, for support/correlation
	Body       []byte // raw response body (truncated for very large bodies)
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	msg := e.Message
	if msg == "" {
		msg = http.StatusText(e.StatusCode)
	}
	if e.Code != "" {
		return fmt.Sprintf("lockally: %d %s (%s)", e.StatusCode, msg, e.Code)
	}
	return fmt.Sprintf("lockally: %d %s", e.StatusCode, msg)
}

// Retryable reports whether the server indicated a transient condition.
func (e *APIError) Retryable() bool { return isRetryableStatus(e.StatusCode) }

// FromResponse builds an *APIError from an HTTP response. It reads and closes
// the body. Returns nil for 2xx responses.
func FromResponse(resp *http.Response) *APIError {
	if resp == nil {
		return nil
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	const maxBody = 64 << 10
	body, _ := io.ReadAll(io.LimitReader(resp.Body, maxBody))
	_ = resp.Body.Close()

	e := &APIError{
		StatusCode: resp.StatusCode,
		RequestID:  resp.Header.Get("X-Request-Id"),
		Body:       body,
	}
	// Best-effort parse of the standard {error:{code,message}} / {message} shapes.
	if strings.Contains(resp.Header.Get("Content-Type"), "json") || looksJSON(body) {
		var p struct {
			Error struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if json.Unmarshal(body, &p) == nil {
			e.Code = firstNonEmpty(p.Error.Code, p.Code)
			e.Message = firstNonEmpty(p.Error.Message, p.Message)
		}
	}
	return e
}

func looksJSON(b []byte) bool {
	b = trimSpace(b)
	return len(b) > 0 && (b[0] == '{' || b[0] == '[')
}

func trimSpace(b []byte) []byte {
	for len(b) > 0 && (b[0] == ' ' || b[0] == '\n' || b[0] == '\t' || b[0] == '\r') {
		b = b[1:]
	}
	return b
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
