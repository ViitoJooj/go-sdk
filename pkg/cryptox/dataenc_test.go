package cryptox

import (
	"encoding/base64"
	"testing"
)

func TestEncryptRoundTripAndBlindIndex(t *testing.T) {
	if err := InitDataKey(base64.StdEncoding.EncodeToString(make([]byte, 32))); err != nil {
		t.Fatal(err)
	}

	const plain = "user@example.com"

	ct, err := Encrypt(plain)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Decrypt(ct)
	if err != nil {
		t.Fatal(err)
	}
	if got != plain {
		t.Fatalf("round trip: got %q want %q", got, plain)
	}

	if ct2, _ := Encrypt(plain); ct == ct2 {
		t.Fatal("ciphertext is not randomized")
	}

	bidx, bidxAgain := BlindIndex(plain), BlindIndex(plain)
	if bidx != bidxAgain {
		t.Fatal("blind index not deterministic")
	}
	if BlindIndex("a") == BlindIndex("b") {
		t.Fatal("blind index collision")
	}
}
