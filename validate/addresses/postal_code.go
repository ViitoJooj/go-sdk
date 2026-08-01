package addresses

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/ViitoJooj/go-sdk/internal"
)

func PostalCode(postal string) error {
	postal = strings.TrimSpace(postal)
	if postal == "" {
		return errors.New("postal code cannot be empty.")
	}

	var errs []error
	n := len(postal)

	if n > 12 {
		errs = append(errs, fmt.Errorf("postal code cannot exceed %d characters. (current %d)", 12, n))
	}
	if n < 2 {
		errs = append(errs, fmt.Errorf("postal code cannot be shorter than %d characters. (current %d)", 2, n))
	}
	if internal.ContainsNullByte(postal) || internal.ContainsControlChars(postal) {
		errs = append(errs, errors.New("postal code cannot contain control characters."))
	}
	for _, r := range postal {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || unicode.IsDigit(r) ||
			r == ' ' || r == '-' {
			continue
		}
		errs = append(errs, fmt.Errorf("postal code contains an invalid character: '%c'.", r))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
