package finance

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var swiftRegex = regexp.MustCompile(`^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`)

func Swift(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))
	if code == "" {
		return errors.New("SWIFT/BIC code cannot be empty.")
	}
	if len(code) != 8 && len(code) != 11 {
		return fmt.Errorf("SWIFT/BIC must be 8 or 11 characters. (current %d)", len(code))
	}
	if !swiftRegex.MatchString(code) {
		return errors.New("invalid SWIFT/BIC format.")
	}
	return nil
}
