package internal

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func NotEmpty(s, label string) error {
	if s == "" {
		return fmt.Errorf("%s cannot be empty.", label)
	}
	return nil
}

func MaxLen(s string, max int, label string) error {
	if len(s) > max {
		return fmt.Errorf("%s cannot exceed %d characters. (current %d)", label, max, len(s))
	}
	return nil
}

func MinLen(s string, min int, label string) error {
	if len(s) < min {
		return fmt.Errorf("%s cannot be shorter than %d characters. (current %d)", label, min, len(s))
	}
	return nil
}

func LenBetween(s string, min, max int, label string) error {
	if len(s) < min {
		return fmt.Errorf("%s cannot be shorter than %d characters. (current %d)", label, min, len(s))
	}
	if len(s) > max {
		return fmt.Errorf("%s cannot exceed %d characters. (current %d)", label, max, len(s))
	}
	return nil
}

func RuneMaxLen(s string, max int, label string) error {
	n := utf8.RuneCountInString(s)
	if n > max {
		return fmt.Errorf("%s cannot exceed %d characters. (current %d)", label, max, n)
	}
	return nil
}

func RuneMinLen(s string, min int, label string) error {
	n := utf8.RuneCountInString(s)
	if n < min {
		return fmt.Errorf("%s cannot be shorter than %d characters. (current %d)", label, min, n)
	}
	return nil
}

func RuneLenBetween(s string, min, max int, label string) error {
	n := utf8.RuneCountInString(s)
	if n < min {
		return fmt.Errorf("%s cannot be shorter than %d characters. (current %d)", label, min, n)
	}
	if n > max {
		return fmt.Errorf("%s cannot exceed %d characters. (current %d)", label, max, n)
	}
	return nil
}

func SafeString(s, label string) error {
	var errs []error
	if ContainsNullByte(s) {
		errs = append(errs, fmt.Errorf("%s cannot contain null bytes.", label))
	}
	if ContainsControlChars(s) {
		errs = append(errs, fmt.Errorf("%s cannot contain control characters.", label))
	}
	if ContainsInvalidUTF8(s) {
		errs = append(errs, fmt.Errorf("%s contains invalid UTF-8 characters.", label))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func OnlyChars(s, allowed, label string) error {
	for _, r := range s {
		if !stringsContainRuneOrFunc(allowed, func(r2 rune) bool {
			return unicode.IsLetter(r2) || unicode.IsMark(r2) || unicode.IsDigit(r2)
		}, r) {
			return fmt.Errorf("%s contains an invalid character: '%c'.", label, r)
		}
	}
	return nil
}

func stringsContainRuneOrFunc(allowed string, fn func(rune) bool, r rune) bool {
	if fn(r) {
		return true
	}
	for _, a := range allowed {
		if a == r {
			return true
		}
	}
	return false
}
