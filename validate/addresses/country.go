package addresses

import (
	"errors"
	"fmt"
	"strings"
)

func Country(code string) error {
	var errs []error
	code = strings.TrimSpace(code)

	if len(code) == 0 {
		return errors.New("country code cannot be empty.")
	}

	if len(code) != 2 {
		errs = append(errs, fmt.Errorf("country code must be exactly 2 letters (ISO 3166-1 alpha-2). (current %d)", len(code)))
	}

	for i := 0; i < len(code); i++ {
		c := code[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			continue
		}
		errs = append(errs, fmt.Errorf("country code contains an invalid character: '%c'.", c))
		break
	}

	if code != strings.ToUpper(code) {
		errs = append(errs, errors.New("country code must be uppercase."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
