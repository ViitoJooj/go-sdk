package users

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Username(username string) error {
	var errs []error

	if len(username) == 0 {
		errs = append(errs, errors.New("The username address cannot be empty."))
	}

	if len(username) > 50 {
		errs = append(errs, fmt.Errorf("username cannot exceed %d characters. (current %d)", 50, len(username)))
	}

	if len(username) < 3 {
		errs = append(errs, fmt.Errorf("username addresses cannot be shorter than %d characters. (current %d)", 3, len(username)))
	}

	for i := 0; i < len(username); i++ {
		c := username[i]
		if (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') ||
			c == '_' ||
			c == '-' ||
			c == '.' {
			continue
		}
		errs = append(errs, fmt.Errorf("The username contains an invalid character: '%c'.", c))
	}

	if internal.IsNumericOnly(username) {
		errs = append(errs, errors.New("username cannot contain only numbers."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
