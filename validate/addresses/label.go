package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Label(label string) error {
	label = strings.TrimSpace(label)
	if label == "" {
		return errors.New("label cannot be empty.")
	}

	var errs []error

	if len(label) > 50 {
		errs = append(errs, fmt.Errorf("label cannot exceed %d characters. (current %d)", 50, len(label)))
	}
	if internal.ContainsNullByte(label) || internal.ContainsControlChars(label) {
		errs = append(errs, errors.New("label cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(label) {
		errs = append(errs, errors.New("label contains invalid UTF-8 characters."))
	}
	for _, r := range label {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(" .,'-_#", r) {
			continue
		}
		errs = append(errs, fmt.Errorf("label contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
