package users

import (
	"errors"
	"fmt"
	"strings"
)

func Phone(phone string) error {
	var errs []error
	phone = strings.TrimSpace(phone)

	if len(phone) == 0 {
		errs = append(errs, errors.New("phone cannot be empty."))
		return errs[0]
	}

	digits := countDigits(phone)

	if digits > 20 {
		errs = append(errs, fmt.Errorf("phone cannot exceed %d digits. (current %d)", 20, digits))
	}

	if digits < 8 {
		errs = append(errs, fmt.Errorf("phone cannot be shorter than %d digits. (current %d)", 8, digits))
	}

	for i := 0; i < len(phone); i++ {
		c := phone[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '+' && i == 0 {
			continue
		}
		errs = append(errs, fmt.Errorf("phone contains an invalid character: '%c'.", c))
		break
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func countDigits(phoneNumber string) int {
	count := 0
	for i := 0; i < len(phoneNumber); i++ {
		if phoneNumber[i] >= '0' && phoneNumber[i] <= '9' {
			count++
		}
	}
	return count
}
