package validate

import "github.com/ViitoJooj/go-sdk/validate/users"

// Config carries the rules every configurable validator uses. The zero
// value is not ready to use (its PasswordRules and EmailRules are empty, so
// everything would be rejected) — build one with Default and override only
// the fields that need to change.
type Config struct {
	PasswordRules users.PasswordRules
	EmailRules    users.EmailRules
}

// Default returns a Config with the SDK's default rules.
func Default() *Config {
	return &Config{
		PasswordRules: users.DefaultPassword,
		EmailRules:    users.DefaultEmail,
	}
}
