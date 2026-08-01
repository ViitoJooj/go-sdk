package users

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Password(password string) error {
	var errs []error

	if len(password) == 0 {
		errs = append(errs, errors.New("The password address cannot be empty."))
	}

	if len(password) > 50 {
		errs = append(errs, fmt.Errorf("password cannot exceed %d characters. (current %d)", 50, len(password)))
	}

	if len(password) < 6 {
		errs = append(errs, fmt.Errorf("password addresses cannot be shorter than %d characters. (current %d)", 6, len(password)))
	}

	if !internal.HasSpecialCharacter(password) {
		errs = append(errs, errors.New("The password needs a special character."))
	}

	if internal.ContainsControlChars(password) {
		errs = append(errs, errors.New("The password cannot contain control characters."))
	}

	if internal.ContainsNullByte(password) {
		errs = append(errs, errors.New("The password cannot contain null bytes."))
	}

	if internal.ContainsInvalidUTF8(password) {
		errs = append(errs, errors.New("The password contains invalid UTF-8 characters."))
	}

	if internal.StartsWithWhitespace(password) {
		errs = append(errs, errors.New("The password cannot start with whitespace."))
	}

	if internal.EndsWithWhitespace(password) {
		errs = append(errs, errors.New("The password cannot end with whitespace."))
	}

	if internal.HasSequentialNumbers(password) {
		errs = append(errs, errors.New("The password cannot contain sequential numbers."))
	}

	if internal.HasSequentialLetters(password) {
		errs = append(errs, errors.New("The password cannot contain sequential letters."))
	}

	if internal.HasKeyboardPatterns(password) {
		errs = append(errs, errors.New("The password cannot contain common keyboard patterns."))
	}

	if internal.IsCommonPassword(password) {
		errs = append(errs, errors.New("The password is too common."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
