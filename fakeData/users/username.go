package users

import (
	"fmt"
	"strings"

	"github.com/ViitoJooj/go-sdk/fakeData/shared"
)

func GenValidUsername() string {
	adj := shared.UsernameAdj[shared.Rng.Intn(len(shared.UsernameAdj))]
	noun := shared.UsernameNoun[shared.Rng.Intn(len(shared.UsernameNoun))]
	num := shared.Rng.Intn(1000)
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
	return patterns[shared.Rng.Intn(len(patterns))]()
}
