package internal

import (
	"strings"
	"unicode"
)

var (
	keyboardPatterns = []string{
		"qwerty",
		"asdfgh",
		"zxcvbn",
		"123456",
		"654321",
	}

	commonPasswords = map[string]struct{}{
		"password": {},
		"123456":   {},
		"12345678": {},
		"qwerty":   {},
		"admin":    {},
		"senha123": {},
	}
)

func HasNumber(password string) bool {
	for _, r := range password {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func HasLetter(password string) bool {
	for i := 0; i < len(password); i++ {
		c := password[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			return true
		}
	}
	return false
}

func HasSpecialCharacter(password string) bool {
	for i := 0; i < len(password); i++ {
		c := password[i]
		if !(c >= 'a' && c <= 'z') &&
			!(c >= 'A' && c <= 'Z') &&
			!(c >= '0' && c <= '9') {
			return true
		}
	}
	return false
}

func HasSequentialNumbers(s string) bool {
	count := 1
	for i := 1; i < len(s); i++ {
		prev, curr := s[i-1], s[i]
		if curr >= '0' && curr <= '9' && prev >= '0' && prev <= '9' && curr == prev+1 {
			count++
			if count >= 3 {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}

func HasSequentialLetters(s string) bool {
	count := 1
	for i := 1; i < len(s); i++ {
		prev, curr := s[i-1], s[i]
		isLetter := func(c byte) bool {
			return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		}
		if isLetter(curr) && isLetter(prev) && curr == prev+1 {
			count++
			if count >= 3 {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}

func HasKeyboardPatterns(s string) bool {
	s = strings.ToLower(s)
	for _, p := range keyboardPatterns {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

func IsCommonPassword(password string) bool {
	_, exists := commonPasswords[strings.ToLower(password)]
	return exists
}
