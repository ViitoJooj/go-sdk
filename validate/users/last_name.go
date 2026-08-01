package users

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ViitoJooj/go-sdk/internal"
)

func LastName(name string) error {
	var errs []error

	if len(name) == 0 {
		errs = append(errs, errors.New("last name cannot be empty."))
		return errs[0]
	}

	if utf8.RuneCountInString(name) > 100 {
		errs = append(errs, fmt.Errorf("last name cannot exceed %d characters. (current %d)", 100, utf8.RuneCountInString(name)))
	}

	if utf8.RuneCountInString(name) < 2 {
		errs = append(errs, fmt.Errorf("last name cannot be shorter than %d characters. (current %d)", 2, utf8.RuneCountInString(name)))
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && r != '-' && r != '\'' && r != ' ' {
			errs = append(errs, fmt.Errorf("last name contains an invalid character: '%c'.", r))
			break
		}
	}

	if strings.Contains(name, "  ") {
		errs = append(errs, errors.New("last name cannot contain consecutive spaces."))
	}

	if internal.StartsWithInvalidNameChar(name) || name[0] == ' ' {
		errs = append(errs, errors.New("last name cannot start with '-', \"'\" or space."))
	}

	if internal.EndsWithInvalidNameChar(name) || name[len(name)-1] == ' ' {
		errs = append(errs, errors.New("last name cannot end with '-', \"'\" or space."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
