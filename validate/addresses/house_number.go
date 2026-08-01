package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func HouseNumber(number string) error {
	number = strings.TrimSpace(number)
	if number == "" {
		return errors.New("house number cannot be empty.")
	}

	var errs []error

	if len(number) > 50 {
		errs = append(errs, fmt.Errorf("house number cannot exceed %d characters. (current %d)", 50, len(number)))
	}
	if internal.ContainsNullByte(number) || internal.ContainsControlChars(number) {
		errs = append(errs, errors.New("house number cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(number) {
		errs = append(errs, errors.New("house number contains invalid UTF-8 characters."))
	}
	for _, r := range number {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || strings.ContainsRune(" /-.°ºª", r) {
			continue
		}
		errs = append(errs, fmt.Errorf("house number contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
