package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func StateRegion(stateRegion string) error {
	stateRegion = strings.TrimSpace(stateRegion)
	if stateRegion == "" {
		return errors.New("state_region cannot be empty.")
	}

	var errs []error

	if len(stateRegion) > 100 {
		errs = append(errs, fmt.Errorf("state_region cannot exceed %d characters. (current %d)", 100, len(stateRegion)))
	}
	if internal.ContainsNullByte(stateRegion) || internal.ContainsControlChars(stateRegion) {
		errs = append(errs, errors.New("state_region cannot contain control characters."))
	}
	if internal.ContainsInvalidUTF8(stateRegion) {
		errs = append(errs, errors.New("state_region contains invalid UTF-8 characters."))
	}
	for _, r := range stateRegion {
		if unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(" -'.,", r) {
			continue
		}
		errs = append(errs, fmt.Errorf("state_region contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
