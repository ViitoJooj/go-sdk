package internal

import "unicode/utf8"

func StripNonDigits(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			b = append(b, s[i])
		}
	}
	return string(b)
}

func AllSameDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func HasAccent(s string) bool {
	for _, r := range s {
		if r > 127 {
			return true
		}
	}
	return false
}

func ContainsControlChars(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] == 127 {
			return true
		}
	}
	return false
}

func ContainsNullByte(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return true
		}
	}
	return false
}

func ContainsInvalidUTF8(s string) bool {
	return !utf8.ValidString(s)
}

func StartsWithWhitespace(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

func EndsWithWhitespace(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[len(s)-1] {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}
