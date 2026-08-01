package users

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ViitoJooj/go-sdk/internal"
)

func FullName(name string) error {
	var errs []error

	if len(name) == 0 {
		errs = append(errs, errors.New("full name cannot be empty."))
		return errs[0]
	}

	if utf8.RuneCountInString(name) > 150 {
		errs = append(errs, fmt.Errorf("full name cannot exceed %d characters. (current %d)", 150, utf8.RuneCountInString(name)))
	}

	if utf8.RuneCountInString(name) < 5 {
		errs = append(errs, fmt.Errorf("full name cannot be shorter than %d characters. (current %d)", 5, utf8.RuneCountInString(name)))
	}

	parts := strings.Fields(name)
	if len(parts) < 2 {
		errs = append(errs, errors.New("full name must contain at least first and last name."))
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && r != '-' && r != '\'' && r != ' ' {
			errs = append(errs, fmt.Errorf("full name contains an invalid character: '%c'.", r))
			break
		}
	}

	if strings.Contains(name, "  ") {
		errs = append(errs, errors.New("full name cannot contain consecutive spaces."))
	}

	if len(name) > 0 && (name[0] == ' ' || internal.StartsWithInvalidNameChar(name)) {
		errs = append(errs, errors.New("full name cannot start with '-', \"'\" or space."))
	}

	if len(name) > 0 && (name[len(name)-1] == ' ' || internal.EndsWithInvalidNameChar(name)) {
		errs = append(errs, errors.New("full name cannot end with '-', \"'\" or space."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
