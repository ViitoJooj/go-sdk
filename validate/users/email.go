package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ViitoJooj/go-sdk/internal"
	"github.com/ViitoJooj/go-sdk/validate/networks"
)

// EmailRules are the length bounds applied to an email address. The local
// part and the domain are always checked (RFC 5322 local-part characters,
// domain labels and TLD via networks.Domain); only the length is
// configurable.
type EmailRules struct {
	MaxLen int
	MinLen int
}

// DefaultEmail: max 254 (RFC 5321 path length), min 6 ("a@b.co").
var DefaultEmail = EmailRules{MaxLen: 254, MinLen: 6}

// Email validates with the default rules. The SDK only validates: it never
// trims or lowercases the input for you.
func Email(email string) error { return DefaultEmail.Validate(email) }

func (r EmailRules) Validate(email string) error {
	var errs []error

	errs = append(errs, validateLen(email, r.MaxLen, r.MinLen))
	errs = append(errs, validateInvalidChars(email))

	local, domain, ok := splitEmail(email)
	switch {
	case !ok:
		errs = append(errs, errors.New(`email must contain exactly one "@"`))
	case local == "":
		errs = append(errs, errors.New("email local part cannot be empty"))
	default:
		if err := networks.Domain(domain); err != nil {
			errs = append(errs, fmt.Errorf("invalid email domain: %w", err))
		}
		if disposable, err := internal.IsDisposable(email); err == nil && disposable {
			errs = append(errs, errors.New("temporary mails are not allowed"))
		}
	}

	return errors.Join(errs...)
}

func splitEmail(email string) (local, domain string, ok bool) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func validateLen(email string, max, min int) error {
	if len(email) == 0 {
		return errors.New("the email address cannot be empty")
	}

	var errs []error

	if len(email) > max {
		errs = append(errs, fmt.Errorf("mails cannot exceed %d characters (current %d)", max, len(email)))
	}

	if len(email) < min {
		errs = append(errs, fmt.Errorf("mail addresses cannot be shorter than %d characters (current %d)", min, len(email)))
	}

	return errors.Join(errs...)
}

// validateInvalidChars rejects whitespace, accents, uppercase and the
// characters RFC 5322 does not allow unquoted in the local part. Everything
// else in atext (! # $ % & ' * + / = ? ^ _ ` { | } ~) is accepted.
func validateInvalidChars(email string) error {
	var errs []error

	if strings.Contains(email, " ") {
		errs = append(errs, errors.New("mails cannot contain spaces"))
	}

	if email != strings.ToLower(email) {
		errs = append(errs, errors.New("mail addresses cannot be in uppercase"))
	}

	if internal.HasAccent(email) {
		errs = append(errs, errors.New("mail cannot contain accents"))
	}

	invalidChars := []string{"<", ">", "(", ")", "[", "]", ",", ";", ":", "\\", "\""}
	for _, c := range invalidChars {
		if strings.Contains(email, c) {
			errs = append(errs, fmt.Errorf("mails cannot contain %q", c))
		}
	}

	return errors.Join(errs...)
}
