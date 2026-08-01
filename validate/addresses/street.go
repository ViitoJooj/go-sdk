package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Street(street string) error {
	street = strings.TrimSpace(street)
	if street == "" {
		return errors.New("street cannot be empty.")
	}

	var errs []error

	if len(street) > 255 {
		errs = append(errs, fmt.Errorf("street cannot exceed %d characters. (current %d)", 255, len(street)))
	}
	if internal.ContainsNullByte(street) || internal.ContainsControlChars(street) {
		errs = append(errs, errors.New("street cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(street) {
		errs = append(errs, errors.New("street contains invalid UTF-8 characters."))
	}
	for _, r := range street {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(streetAllowedPunct, r) {
			continue
		}
		errs = append(errs, fmt.Errorf("street contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
