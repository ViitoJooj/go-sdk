package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Field-level encryption for PII at rest. Two subkeys are derived from a single
// 32-byte master key with domain separation:
//   - an AES-256-GCM key for reversible, randomized encryption of stored values
//   - an HMAC-SHA256 key for deterministic "blind index" tokens, which let us
//     look up and enforce uniqueness on encrypted columns without storing or
//     exposing the plaintext.
var (
	dataAEAD cipher.AEAD
	bidxKey  []byte
)

// InitDataKey initializes the field-encryption keys from a base64-encoded
// 32-byte master key. Must be called once at startup before any Encrypt/Decrypt/
// BlindIndex call. Losing the key permanently destroys all encrypted data.
func InitDataKey(b64Master string) error {
	master, err := base64.StdEncoding.DecodeString(b64Master)
	if err != nil {
		return fmt.Errorf("decode data key: %w", err)
	}
	if len(master) != 32 {
		return fmt.Errorf("data key must be 32 bytes, got %d", len(master))
	}

	encKey := sha256.Sum256(append([]byte("tentaculum-enc:"), master...))
	block, err := aes.NewCipher(encKey[:])
	if err != nil {
		return err
	}
	dataAEAD, err = cipher.NewGCM(block)
	if err != nil {
		return err
	}

	macKey := sha256.Sum256(append([]byte("tentaculum-bidx:"), master...))
	bidxKey = macKey[:]
	return nil
}

// Encrypt encrypts plaintext with AES-256-GCM and returns base64(nonce||ciphertext).
// Each call produces a different ciphertext (random nonce), so encrypted columns
// cannot be used for lookups or UNIQUE constraints — use BlindIndex for that.
func Encrypt(plaintext string) (string, error) {
	if dataAEAD == nil {
		return "", errors.New("data encryption key not initialized")
	}
	nonce := make([]byte, dataAEAD.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := dataAEAD.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ct), nil
}

// Decrypt reverses Encrypt.
func Decrypt(encoded string) (string, error) {
	if dataAEAD == nil {
		return "", errors.New("data encryption key not initialized")
	}
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	ns := dataAEAD.NonceSize()
	if len(raw) < ns {
		return "", errors.New("ciphertext too short")
	}
	pt, err := dataAEAD.Open(nil, raw[:ns], raw[ns:], nil)
	if err != nil {
		return "", err
	}
	return string(pt), nil
}

// BlindIndex returns a deterministic HMAC-SHA256 token (hex) of value. Stored in
// a side column, it lets encrypted fields be looked up and kept unique without
// exposing the plaintext. The match is exact (no case folding).
func BlindIndex(value string) string {
	mac := hmac.New(sha256.New, bidxKey)
	mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))
}
