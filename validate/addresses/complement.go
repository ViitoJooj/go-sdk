package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Complement(complement string) error {
	complement = strings.TrimSpace(complement)
	if complement == "" {
		return errors.New("complement cannot be empty.")
	}

	var errs []error

	if len(complement) > 255 {
		errs = append(errs, fmt.Errorf("complement cannot exceed %d characters. (current %d)", 255, len(complement)))
	}
	if internal.ContainsNullByte(complement) || internal.ContainsControlChars(complement) {
		errs = append(errs, errors.New("complement cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(complement) {
		errs = append(errs, errors.New("complement contains invalid UTF-8 characters."))
	}
	for _, r := range complement {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(streetAllowedPunct, r) {
			continue
		}
		errs = append(errs, fmt.Errorf("complement contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
