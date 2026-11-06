package lockallyx

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"
)

// These vectors are generated from internal/webhook, the package that signs
// the deliveries customers receive. Reading them here is the whole point: a
// verifier that drifts from the signer either rejects real webhooks or
// accepts forged ones, and a test written against this file's own idea of
// the algorithm would pass while doing both.
const casesPath = "../testdata/webhook-cases.json"

type cases struct {
	SecretHex string `json:"secret_hex"`
	NowUnix   int64  `json:"now_unix"`
	MaxSkewS  int    `json:"max_skew_seconds"`
	Sign      []struct {
		Name   string `json:"name"`
		Body   string `json:"body"`
		Header string `json:"header"`
	} `json:"sign"`
	Verify []struct {
		Name    string `json:"name"`
		Body    string `json:"body"`
		Header  string `json:"header"`
		Outcome string `json:"outcome"`
		Why     string `json:"why"`
	} `json:"verify"`
}

func load(t *testing.T) cases {
	t.Helper()
	raw, err := os.ReadFile(casesPath)
	if err != nil {
		t.Fatal(err)
	}
	var c cases
	if err := json.Unmarshal(raw, &c); err != nil {
		t.Fatal(err)
	}
	return c
}

func outcome(err error) string {
	switch {
	case err == nil:
		return "ok"
	case errors.Is(err, ErrWebhookExpired):
		return "expired"
	case errors.Is(err, ErrWebhookMismatch):
		return "mismatch"
	case errors.Is(err, ErrWebhookMalformed):
		return "malformed"
	}
	return "unknown"
}

func TestVerifyWebhook_AgainstSharedCases(t *testing.T) {
	c := load(t)
	now := time.Unix(c.NowUnix, 0)
	secret := key(t, c)

	if got := int(WebhookMaxSkew / time.Second); got != c.MaxSkewS {
		t.Errorf("skew window is %ds here and %ds on the server", got, c.MaxSkewS)
	}

	for _, v := range c.Verify {
		got := outcome(verifyWebhookAt(secret, []byte(v.Body), v.Header, now))
		if got != v.Outcome {
			t.Errorf("%s: got %q, server says %q\n  %s", v.Name, got, v.Outcome, v.Why)
		}
	}
}

// Every signature the server produces must verify here. The sign vectors are
// the server's actual output, so this is the end-to-end direction.
func TestVerifyWebhook_AcceptsEverySignedBody(t *testing.T) {
	c := load(t)
	now := time.Unix(c.NowUnix, 0)
	for _, s := range c.Sign {
		if err := verifyWebhookAt(key(t, c), []byte(s.Body), s.Header, now); err != nil {
			t.Errorf("%s: a signature the server produced was refused: %v", s.Name, err)
		}
	}
}

// A wrong secret must be a mismatch, not an accept. Worth its own test
// because it is the failure that matters and the fixture holds only one
// secret.
func TestVerifyWebhook_WrongSecretIsAMismatch(t *testing.T) {
	c := load(t)
	now := time.Unix(c.NowUnix, 0)
	s := c.Sign[0]
	other, _ := hex.DecodeString("00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff")
	err := verifyWebhookAt(other, []byte(s.Body), s.Header, now)
	if !errors.Is(err, ErrWebhookMismatch) {
		t.Fatalf("wrong secret gave %v, want ErrWebhookMismatch", err)
	}
}

// The public entry point takes the hex secret the dashboard shows and uses
// the real clock. Both are things the fixture-driven tests above cannot
// prove, and both are what a caller actually touches.
func TestVerifyWebhook_PublicEntryPointTakesTheHexSecret(t *testing.T) {
	c := load(t)
	body := []byte(`{"type":"delivery.delivered"}`)
	now := time.Now()
	hdr := signForTest(key(t, c), body, now)

	// The string straight out of the dashboard.
	if err := VerifyWebhook(c.SecretHex, body, hdr); err != nil {
		t.Fatalf("the hex secret a customer is given was refused: %v", err)
	}
	// And the decoded bytes, for a caller who did it themselves.
	if err := VerifyWebhookKey(key(t, c), body, hdr); err != nil {
		t.Fatalf("the decoded key was refused: %v", err)
	}

	stale := signForTest(key(t, c), body, now.Add(-WebhookMaxSkew-time.Minute))
	if err := VerifyWebhook(c.SecretHex, body, stale); !errors.Is(err, ErrWebhookExpired) {
		t.Fatalf("a stale signature gave %v, want ErrWebhookExpired", err)
	}
}

// Passing the hex TEXT as the key is the documented first-attempt failure.
// It must not read as a forged payload, or the caller spends the afternoon
// looking at the sender.
func TestVerifyWebhook_HexTextAsKeyIsNotMistakenForAMismatch(t *testing.T) {
	c := load(t)
	body := []byte(`{"type":"delivery.delivered"}`)
	hdr := signForTest(key(t, c), body, time.Now())

	if err := VerifyWebhookKey([]byte(c.SecretHex), body, hdr); !errors.Is(err, ErrWebhookMismatch) {
		t.Fatalf("got %v; using the hex text as the key should fail", err)
	}
	if err := VerifyWebhook("not hex at all", body, hdr); !errors.Is(err, ErrWebhookSecret) {
		t.Fatalf("a non-hex secret gave %v, want ErrWebhookSecret", err)
	}
}

func key(t *testing.T, c cases) []byte {
	t.Helper()
	k, err := hex.DecodeString(c.SecretHex)
	if err != nil {
		t.Fatal(err)
	}
	return k
}
