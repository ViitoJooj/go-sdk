package fakeData

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var validEmailDomains = []string{"gmail.com", "outlook.com", "yahoo.com", "hotmail.com", "proton.me"}
var usernameAdj = []string{"cool", "fast", "happy", "lazy", "quiet", "brave", "shy", "wild", "calm", "bold"}
var usernameNoun = []string{"panda", "tiger", "eagle", "wolf", "fox", "bear", "hawk", "lion", "deer", "owl"}

func GenValidEmail() string {
	user := randomString(8)
	domain := validEmailDomains[rng.Intn(len(validEmailDomains))]
	return fmt.Sprintf("%s@%s", strings.ToLower(user), domain)
}

func GenInvalidEmail() string {
	patterns := []func() string{
		func() string { return "no-arroba.com" },
		func() string { return "@no-local-part.com" },
		func() string { return "spaces in@mail.com" },
		func() string { return "UPPERCASE@mail.com" },
		func() string { return "" },
		func() string { return "a@b" },
	}
	return patterns[rng.Intn(len(patterns))]()
}

func GenValidUsername() string {
	adj := usernameAdj[rng.Intn(len(usernameAdj))]
	noun := usernameNoun[rng.Intn(len(usernameNoun))]
	num := rng.Intn(1000)
	return fmt.Sprintf("%s_%s%d", adj, noun, num)
}

func GenInvalidUsername() string {
	patterns := []func() string{
		func() string { return "" },
		func() string { return "ab" },
		func() string { return "123456" },
		func() string { return "user name" },
		func() string { return "invalid@user" },
		func() string { return strings.Repeat("a", 60) },
	}
	return patterns[rng.Intn(len(patterns))]()
}

func GenValidCPF() string {
	digits := make([]int, 9)
	for i := 0; i < 9; i++ {
		digits[i] = rng.Intn(10)
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}

	sum = 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (11 - i)
	}
	sum += d1 * 2
	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d",
		digits[0], digits[1], digits[2], digits[3], digits[4],
		digits[5], digits[6], digits[7], digits[8], d1, d2)
}

func GenValidCNPJ() string {
	digits := make([]int, 12)
	for i := 0; i < 12; i++ {
		digits[i] = rng.Intn(10)
	}

	w1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digits[i] * w1[i]
	}
	r := sum % 11
	dv1 := 0
	if r >= 2 {
		dv1 = 11 - r
	}

	w2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum = 0
	for i := 0; i < 12; i++ {
		sum += digits[i] * w2[i]
	}
	sum += dv1 * w2[12]
	r = sum % 11
	dv2 := 0
	if r >= 2 {
		dv2 = 11 - r
	}

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
		digits[0], digits[1], digits[2], digits[3], digits[4],
		digits[5], digits[6], digits[7], digits[8], digits[9],
		digits[10], digits[11], dv1, dv2)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}
