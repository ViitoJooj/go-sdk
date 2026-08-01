package fakeData

import (
	"github.com/ViitoJooj/go-sdk/fakeData/company"
	"github.com/ViitoJooj/go-sdk/fakeData/users"
)

// User - data generate
func GenValidEmail() string    { return users.GenValidEmail() }
func GenInvalidEmail() string  { return users.GenInvalidEmail() }
func GenValidUsername() string { return users.GenValidUsername() }
func GenInvalidUsername() string { return users.GenInvalidUsername() }
func GenValidCPF() string      { return users.GenValidCPF() }

// Company - data generate
func GenValidCNPJ() string { return company.GenValidCNPJ() }
