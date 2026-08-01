package shared

import (
	"math/rand"
	"time"
)

var Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var ValidEmailDomains = []string{"gmail.com", "outlook.com", "yahoo.com", "hotmail.com", "proton.me"}
var UsernameAdj = []string{"cool", "fast", "happy", "lazy", "quiet", "brave", "shy", "wild", "calm", "bold"}
var UsernameNoun = []string{"panda", "tiger", "eagle", "wolf", "fox", "bear", "hawk", "lion", "deer", "owl"}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[Rng.Intn(len(letters))]
	}
	return string(b)
}
