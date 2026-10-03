package users

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/ViitoJooj/go-sdk/internal"
)

// PasswordRules follows NIST SP 800-63B: no required character classes by
// default (composition rules are discouraged), length counted in runes
// (not bytes, so accented passwords aren't cut short), and the common
// password / keyboard / sequence checks stay on by default.
type PasswordRules struct {
	Min, Max       int
	RequireSpecial bool
	CheckCommon    bool
	CheckSequences bool
}

var DefaultPassword = PasswordRules{Min: 8, Max: 128, CheckCommon: true, CheckSequences: true}

func Password(password string) error { return DefaultPassword.Validate(password) }

func (r PasswordRules) Validate(password string) error {
	var errs []error

	length := utf8.RuneCountInString(password)

	if length == 0 {
		return errors.New("the password cannot be empty")
	}

	if length > r.Max {
		errs = append(errs, fmt.Errorf("password cannot exceed %d characters (current %d)", r.Max, length))
	}

	if length < r.Min {
		errs = append(errs, fmt.Errorf("password cannot be shorter than %d characters (current %d)", r.Min, length))
	}

	if r.RequireSpecial && !internal.HasSpecialCharacter(password) {
		errs = append(errs, errors.New("the password needs a special character"))
	}

	if internal.ContainsControlChars(password) {
		errs = append(errs, errors.New("the password cannot contain control characters"))
	}

	if internal.ContainsNullByte(password) {
		errs = append(errs, errors.New("the password cannot contain null bytes"))
	}

	if internal.ContainsInvalidUTF8(password) {
		errs = append(errs, errors.New("the password contains invalid UTF-8 characters"))
	}

	if internal.StartsWithWhitespace(password) {
		errs = append(errs, errors.New("the password cannot start with whitespace"))
	}

	if internal.EndsWithWhitespace(password) {
		errs = append(errs, errors.New("the password cannot end with whitespace"))
	}

	if r.CheckSequences && internal.HasSequentialNumbers(password) {
		errs = append(errs, errors.New("the password cannot contain sequential numbers"))
	}

	if r.CheckSequences && internal.HasSequentialLetters(password) {
		errs = append(errs, errors.New("the password cannot contain sequential letters"))
	}

	if r.CheckSequences && internal.HasKeyboardPatterns(password) {
		errs = append(errs, errors.New("the password cannot contain common keyboard patterns"))
	}

	if r.CheckCommon && internal.IsCommonPassword(password) {
		errs = append(errs, errors.New("the password is too common"))
	}

	return errors.Join(errs...)
}
