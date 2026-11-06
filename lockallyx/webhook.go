package lockallyx

// Webhook signature verification.
//
// Every delivery carries:
//
//	X-Lockally-Signature: t=<unix-seconds>,v1=<hex(hmac_sha256(secret, t + "." + body))>
//
// Verifying it by hand is an HMAC, a constant-time compare and a clock
// window, in that order, and getting any of the three wrong fails quietly:
// a == on the digest leaks timing, a missing window accepts a replay, and
// re-serialising the body before hashing rejects every real delivery. It is
// exactly the code an official SDK should remove, and until now none of ours
// had it.
//
// The rules below are pinned against the server that signs the deliveries by
// a shared fixture, sdk/ergonomics/testdata/webhook-cases.json, generated
// from internal/webhook. Three of them are not what you would guess:
//
//   - The window is ONE-SIDED. Only a timestamp too far in the PAST is
//     refused. A future one is accepted, so a verifier that rejects it would
//     drop deliveries the server considers valid.
//   - Unknown pairs are IGNORED, so a later v2= scheme cannot break a
//     verifier written today.
//   - The hex is decoded before comparison, so its case does not matter.
//
// And the one that breaks most first attempts: THE SECRET IS HEX TEXT AND
// THE KEY IS THE BYTES IT DECODES TO. `signing_secret` is BYTEA, so the
// dispatcher signs with 32 raw bytes while the API shows you the hex of
// them, once, at creation. Signing with the text gives a valid-looking
// digest that never matches. VerifyWebhook takes the string you were shown
// and decodes it for you; VerifyWebhookKey takes the raw bytes if you have
// already decoded them yourself.

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// WebhookHeader is the header carrying the signature.
const WebhookHeader = "X-Lockally-Signature"

// WebhookMaxSkew is how old a timestamp may be before it is refused as a
// replay. It matches the server.
const WebhookMaxSkew = 5 * time.Minute

// Errors a caller can branch on. Use errors.Is.
var (
	// ErrWebhookMalformed means the header is not a readable signature.
	ErrWebhookMalformed = errors.New("lockally: malformed webhook signature header")
	// ErrWebhookExpired means the signature is genuine but too old.
	ErrWebhookExpired = errors.New("lockally: webhook timestamp outside the skew window")
	// ErrWebhookMismatch means the body does not match the signature. Treat
	// this as hostile, not as a bug.
	ErrWebhookMismatch = errors.New("lockally: webhook HMAC does not match")
	// ErrWebhookSecret means the secret is not the hex string the dashboard
	// showed you. Named separately from a mismatch because it is your
	// configuration that is wrong, not the caller's signature.
	ErrWebhookSecret = errors.New("lockally: webhook signing secret is not valid hex")
)

// VerifyWebhook reports whether body was signed with secret, as described by
// header. It returns nil when the delivery is genuine and recent.
//
// body MUST be the raw bytes as received. Decoding JSON and re-encoding it
// changes the bytes, and the signature is over the bytes.
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	        body, _ := io.ReadAll(r.Body)
//	        if err := lockallyx.VerifyWebhook(secretFromDashboard, body, r.Header.Get(lockallyx.WebhookHeader)); err != nil {
//	                http.Error(w, "bad signature", http.StatusBadRequest)
//	                return
//	        }
//	        // ... trust body from here
//	}
func VerifyWebhook(secretHex string, body []byte, header string) error {
	key, err := hex.DecodeString(strings.TrimSpace(secretHex))
	if err != nil {
		return fmt.Errorf("%w: %q", ErrWebhookSecret, truncate(secretHex))
	}
	return verifyWebhookAt(key, body, header, time.Now())
}

// VerifyWebhookKey is VerifyWebhook for a caller holding the already decoded
// 32 bytes rather than the hex string.
func VerifyWebhookKey(key, body []byte, header string) error {
	return verifyWebhookAt(key, body, header, time.Now())
}

// truncate keeps a wrong-secret error readable without printing the whole
// value into a log.
func truncate(s string) string {
	if len(s) <= 8 {
		return s
	}
	return s[:8] + "..."
}

// VerifyWebhookRequest is VerifyWebhook for an *http.Request whose body you
// have already read. It saves looking the header name up.
func VerifyWebhookRequest(secretHex string, body []byte, r *http.Request) error {
	if r == nil {
		return fmt.Errorf("%w: nil request", ErrWebhookMalformed)
	}
	return VerifyWebhook(secretHex, body, r.Header.Get(WebhookHeader))
}

// verifyWebhookAt is the testable core: `now` is injected so the fixture can
// pin the clock.
func verifyWebhookAt(secret, body []byte, header string, now time.Time) error {
	ts, sigHex, err := parseWebhookSignature(header)
	if err != nil {
		return err
	}

	tsInt, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: t= is not an integer", ErrWebhookMalformed)
	}
	// One-sided on purpose: see the note at the top of this file.
	if now.Sub(time.Unix(tsInt, 0)) > WebhookMaxSkew {
		return ErrWebhookExpired
	}

	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return fmt.Errorf("%w: v1= is not hex", ErrWebhookMalformed)
	}

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(ts))
	mac.Write([]byte{'.'})
	mac.Write(body)

	// hmac.Equal, not bytes.Equal: the comparison has to take the same time
	// whether the first byte is wrong or only the last.
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return ErrWebhookMismatch
	}
	return nil
}

func parseWebhookSignature(header string) (ts, sig string, err error) {
	for _, part := range strings.Split(header, ",") {
		k, v, ok := strings.Cut(strings.TrimSpace(part), "=")
		if !ok {
			return "", "", fmt.Errorf("%w: expected k=v pairs", ErrWebhookMalformed)
		}
		// Unknown keys are skipped rather than refused, so a later scheme
		// does not break a verifier written today.
		switch strings.TrimSpace(k) {
		case "t":
			ts = strings.TrimSpace(v)
		case "v1":
			sig = strings.TrimSpace(v)
		}
	}
	if ts == "" || sig == "" {
		return "", "", fmt.Errorf("%w: missing t= or v1=", ErrWebhookMalformed)
	}
	return ts, sig, nil
}
