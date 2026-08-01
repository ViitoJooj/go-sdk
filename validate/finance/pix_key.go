package finance

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/ViitoJooj/go-sdk/internal"
)

var pixKeyRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func PixKey(key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("PIX key cannot be empty.")
	}
	if len(key) > 77 {
		return fmt.Errorf("PIX key cannot exceed %d characters. (current %d)", 77, len(key))
	}

	digits := internal.StripNonDigits(key)

	switch {
	case len(digits) == 11:
		return nil
	case len(digits) == 14:
		return nil
	case strings.HasPrefix(key, "+"):
		for i := 1; i < len(key); i++ {
			c := key[i]
			if (c >= '0' && c <= '9') || c == ' ' || c == '-' {
				continue
			}
			return fmt.Errorf("PIX key (phone) contains invalid character: '%c'.", c)
		}
		return nil
	case pixKeyRegex.MatchString(key):
		return nil
	case len(digits) == 36 || isValidUUID(key):
		return nil
	default:
		return errors.New("invalid PIX key format. Must be CPF, CNPJ, phone, email, or random UUID.")
	}
}

func isValidUUID(s string) bool {
	return len(s) == 36 && s[8] == '-' && s[13] == '-' && s[18] == '-' && s[23] == '-'
}
