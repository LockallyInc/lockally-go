package lockallyx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// The one call almost every integration makes, without the ceremony.
//
// The generated client is a faithful translation of the spec, which means
// sending looks like this:
//
//	cfg := lockally.NewConfiguration()
//	cfg.AddDefaultHeader("Authorization", "Bearer "+key)
//	api := lockally.NewAPIClient(cfg)
//	req := lockally.NewSendEmailRequest(from, to)
//	req.SetText("hello")
//	resp, _, err := api.SendAPI.SendEmail(ctx).
//		IdempotencyKey(uuid.NewString()).
//		SendEmailRequest(*req).
//		Execute()
//
// Every part of that is correct and none of it is what anybody wants to read
// in a quickstart. This wraps it, and only it.

const defaultBaseURL = "https://api.lockally.com"

// Message mirrors the spec's SendEmailRequest. Zero values are omitted rather
// than sent as null, so an unset field behaves as absent.
type Message struct {
	From        string            `json:"from"`
	To          []string          `json:"to"`
	Subject     string            `json:"subject,omitempty"`
	Text        string            `json:"text,omitempty"`
	HTML        string            `json:"html,omitempty"`
	CC          []string          `json:"cc,omitempty"`
	BCC         []string          `json:"bcc,omitempty"`
	ReplyTo     []string          `json:"reply_to,omitempty"`
	ReplyToName string            `json:"reply_to_name,omitempty"`
	TemplateID  string            `json:"template_id,omitempty"`
	Variables   map[string]string `json:"variables,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Attachments []Attachment      `json:"attachments,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	SendAt      string            `json:"send_at,omitempty"`
	Unsubscribe bool              `json:"unsubscribe,omitempty"`
}

// Attachment is inline on the send; there is no upload endpoint.
type Attachment struct {
	Filename      string `json:"filename"`
	ContentBase64 string `json:"content_base64"`
	ContentType   string `json:"content_type,omitempty"`
}

// SendResult is what came back. Status is the field to branch on.
type SendResult struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
	Warning   string `json:"warning,omitempty"`
}

// Client sends through POST /v1/send.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithBaseURL points the client somewhere other than api.lockally.com.
func WithBaseURL(u string) ClientOption {
	return func(c *Client) { c.baseURL = strings.TrimRight(u, "/") }
}

// WithHTTPClient supplies the transport. Pass NewClient(nil) to get retries,
// 429 handling and backoff underneath the convenience.
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *Client) { c.http = h }
}

// New returns a Client for the given API key.
func New(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("lockally: an API key is required")
	}
	c := &Client{apiKey: apiKey, baseURL: defaultBaseURL, http: http.DefaultClient}
	for _, o := range opts {
		o(c)
	}
	return c, nil
}

type sendOptions struct{ idempotencyKey string }

// SendOption configures one send.
type SendOption func(*sendOptions)

// WithIdempotencyKey sets the key. Supply one.
//
// It is REQUIRED by the API, so when it is omitted Send generates a UUID to
// keep the call working. That is a convenience, not a default to rely on: a
// generated key is new on every call, so retrying through Send delivers a
// SECOND message rather than replaying the first. The key has to identify the
// message, which only the caller knows.
func WithIdempotencyKey(k string) SendOption {
	return func(o *sendOptions) { o.idempotencyKey = k }
}

// Send submits one message.
//
// It returns a result on 200 as well as 202. A 200 means nothing was queued,
// either because every recipient was suppressed or because the key was a test
// key, and both are successes. Branch on Status, never on the status code.
func (c *Client) Send(ctx context.Context, m Message, opts ...SendOption) (*SendResult, error) {
	var o sendOptions
	for _, f := range opts {
		f(&o)
	}
	if o.idempotencyKey == "" {
		// newUUID is the layer's own, in retry.go: no third-party
		// dependency enters a package that gets folded into the SDK.
		o.idempotencyKey = newUUID()
	}

	body, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("lockally: encode message: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/send", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lockally: send: %w", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		// A non-JSON error body is itself the diagnostic: nginx answers an
		// oversized payload as HTML before the application sees it.
		var p struct {
			Title     string `json:"title"`
			Code      string `json:"code"`
			RequestID string `json:"request_id"`
		}
		_ = json.Unmarshal(raw, &p)
		msg := p.Title
		if msg == "" {
			msg = fmt.Sprintf("lockally: HTTP %d", resp.StatusCode)
		}
		reqID := p.RequestID
		if reqID == "" {
			reqID = resp.Header.Get("X-Request-Id")
		}
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Code:       p.Code,
			Message:    msg,
			RequestID:  reqID,
			Body:       raw,
		}
	}

	var out SendResult
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &out); err != nil {
			return nil, fmt.Errorf("lockally: decode response: %w", err)
		}
	}
	return &out, nil
}
