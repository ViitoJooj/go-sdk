package validate_test

import (
	"testing"

	"github.com/ViitoJooj/go-sdk/fakeData"
	"github.com/ViitoJooj/go-sdk/validate"
)

func TestValidEmail(t *testing.T) {
	cfg := validate.Default()
	email := fakeData.GenValidEmail()

	if err := cfg.Email(email); err != nil {
		t.Errorf("Email(%q): %v", email, err)
	}
}

func TestInvalidEmail(t *testing.T) {
	cfg := validate.Default()
	email := fakeData.GenInvalidEmail()

	if err := cfg.Email(email); err == nil {
		t.Errorf("Email(%q): want error, got nil", email)
	}
}
