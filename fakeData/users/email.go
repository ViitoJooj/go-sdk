package users

import (
	"fmt"
	"strings"

	"github.com/ViitoJooj/go-sdk/fakeData/shared"
)

func GenValidEmail() string {
	user := shared.RandomString(8)
	domain := shared.ValidEmailDomains[shared.Rng.Intn(len(shared.ValidEmailDomains))]
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
	return patterns[shared.Rng.Intn(len(patterns))]()
}
