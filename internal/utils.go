package internal

import (
	"bufio"
	"net/http"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

var (
	disposableDomains map[string]struct{}
	domainsOnce       sync.Once
	domainsLoadErr    error

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

func loadDisposableDomains() {
	disposableDomains = make(map[string]struct{})

	resp, err := http.Get("https://disposable.github.io/disposable-email-domains/domains.txt")
	if err != nil {
		domainsLoadErr = err
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		domain := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if domain != "" {
			disposableDomains[domain] = struct{}{}
		}
	}

	domainsLoadErr = scanner.Err()
}

func init() {
	go domainsOnce.Do(loadDisposableDomains)
}

func IsDisposable(mail string) (bool, error) {
	domainsOnce.Do(loadDisposableDomains)

	if domainsLoadErr != nil {
		return false, domainsLoadErr
	}

	parts := strings.Split(mail, "@")
	if len(parts) != 2 {
		return false, nil
	}

	domain := strings.ToLower(strings.TrimSpace(parts[1]))
	_, exists := disposableDomains[domain]
	return exists, nil
}

func HasAccent(s string) bool {
	for _, r := range s {
		if r > 127 {
			return true
		}
	}
	return false
}

func StripNonDigits(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			b.WriteByte(s[i])
		}
	}
	return b.String()
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

func CPFCheckDigitsValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}

	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}

	if d1 != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}

	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}

	return d2 == int(cpf[10]-'0')
}

func CNPJCheckDigitsValid(d string) bool {
	w1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	w2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(d[i]-'0') * w1[i]
	}
	r := sum % 11
	dv1 := 0
	if r >= 2 {
		dv1 = 11 - r
	}
	if int(d[12]-'0') != dv1 {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(d[i]-'0') * w2[i]
	}
	r = sum % 11
	dv2 := 0
	if r >= 2 {
		dv2 = 11 - r
	}
	return int(d[13]-'0') == dv2
}

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

func StartsWithInvalidNameChar(name string) bool {
	if len(name) == 0 {
		return false
	}
	return name[0] == '-' || name[0] == '\''
}

func EndsWithInvalidNameChar(name string) bool {
	if len(name) == 0 {
		return false
	}
	return name[len(name)-1] == '-' || name[len(name)-1] == '\''
}

func IsNumericOnly(username string) bool {
	if len(username) == 0 {
		return false
	}
	for i := 0; i < len(username); i++ {
		if username[i] < '0' || username[i] > '9' {
			return false
		}
	}
	return true
}
