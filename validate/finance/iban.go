package finance

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ibanRegex = regexp.MustCompile(`^[A-Z]{2}\d{2}[A-Z0-9]{1,30}$`)

func IBAN(iban string) error {
	iban = strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(iban), " ", ""))
	if iban == "" {
		return errors.New("IBAN cannot be empty.")
	}
	if len(iban) < 5 || len(iban) > 34 {
		return fmt.Errorf("IBAN must be between %d and %d characters. (current %d)", 5, 34, len(iban))
	}
	if !ibanRegex.MatchString(iban) {
		return errors.New("invalid IBAN format.")
	}
	return nil
}
