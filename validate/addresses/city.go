package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func City(city string) error {
	city = strings.TrimSpace(city)
	if city == "" {
		return errors.New("city cannot be empty.")
	}

	var errs []error

	if len(city) > 100 {
		errs = append(errs, fmt.Errorf("city cannot exceed %d characters. (current %d)", 100, len(city)))
	}
	if internal.ContainsNullByte(city) || internal.ContainsControlChars(city) {
		errs = append(errs, errors.New("city cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(city) {
		errs = append(errs, errors.New("city contains invalid UTF-8 characters."))
	}
	for _, r := range city {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(" -'.,", r) {
			continue
		}
		errs = append(errs, fmt.Errorf("city contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
