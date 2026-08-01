package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ViitoJooj/go-sdk/internal"
)

func Email(email string) error {
	var errs []error
	email = strings.TrimSpace(email)

	if len(email) == 0 {
		errs = append(errs, errors.New("The email address cannot be empty."))
	}

	if len(email) > 150 {
		errs = append(errs, fmt.Errorf("mails cannot exceed %d characters. (current %d)", 150, len(email)))
	}

	if len(email) < 6 {
		errs = append(errs, fmt.Errorf("mail addresses cannot be shorter than %d characters. (current %d)", 6, len(email)))
	}

	if !strings.Contains(email, "@") {
		errs = append(errs, fmt.Errorf("This is not an mail, because it is missing \"%s\"", "@"))
	}

	if !strings.Contains(email, ".") {
		errs = append(errs, fmt.Errorf("This is not an mail, because it is missing \"%s\"", "."))
	}

	if strings.Contains(email, " ") {
		errs = append(errs, errors.New("mails cannot contain spaces."))
	}

	if email != strings.ToLower(email) {
		errs = append(errs, errors.New("mail addresses cannot be in uppercase."))
	}

	if internal.HasAccent(email) {
		errs = append(errs, errors.New("mail cannot contain accents"))
	}

	invalidChars := []string{"<", ">", "(", ")", "[", "]", ",", ";", ":", "\\", "/", "\"", "'", "!", "#", "$", "%", "^", "&", "*", "=", "+", "{", "}", "|", "?", "~", "`"}
	for _, c := range invalidChars {
		if strings.Contains(email, c) {
			errs = append(errs, fmt.Errorf("mails cannot contain \"%s\"", c))
		}
	}

	sqlPatterns := []string{"' #", "'/*", "' or true --", "\") or (\"1\"=\"1"}
	for _, p := range sqlPatterns {
		if strings.Contains(email, p) {
			errs = append(errs, errors.New("mail format is invalid"))
			break
		}
	}

	if disposable, err := internal.IsDisposable(email); err == nil && disposable {
		errs = append(errs, errors.New("temporary mails are not allowed"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
