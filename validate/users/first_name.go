package users

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ViitoJooj/go-sdk/internal"
)

func FirstName(name string) error {
	var errs []error

	if len(name) == 0 {
		errs = append(errs, errors.New("first name cannot be empty."))
		return errs[0]
	}

	if utf8.RuneCountInString(name) > 50 {
		errs = append(errs, fmt.Errorf("first name cannot exceed %d characters. (current %d)", 50, utf8.RuneCountInString(name)))
	}

	if utf8.RuneCountInString(name) < 2 {
		errs = append(errs, fmt.Errorf("first name cannot be shorter than %d characters. (current %d)", 2, utf8.RuneCountInString(name)))
	}

	if strings.Contains(name, " ") {
		errs = append(errs, errors.New("first name cannot contain spaces."))
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && r != '-' && r != '\'' {
			errs = append(errs, fmt.Errorf("first name contains an invalid character: '%c'.", r))
			break
		}
	}

	if internal.StartsWithInvalidNameChar(name) {
		errs = append(errs, errors.New("first name cannot start with '-' or \"'\"."))
	}

	if internal.EndsWithInvalidNameChar(name) {
		errs = append(errs, errors.New("first name cannot end with '-' or \"'\"."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
