package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func District(district string) error {
	district = strings.TrimSpace(district)
	if district == "" {
		return errors.New("district cannot be empty.")
	}

	var errs []error

	if len(district) > 100 {
		errs = append(errs, fmt.Errorf("district cannot exceed %d characters. (current %d)", 100, len(district)))
	}
	if internal.ContainsNullByte(district) || internal.ContainsControlChars(district) {
		errs = append(errs, errors.New("district cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(district) {
		errs = append(errs, errors.New("district contains invalid UTF-8 characters."))
	}
	for _, r := range district {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(" -'.,", r) {
			continue
		}
		errs = append(errs, fmt.Errorf("district contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
