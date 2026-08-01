package users

import (
	"fmt"

	"github.com/ViitoJooj/go-sdk/fakeData/shared"
)

func GenValidCPF() string {
	digits := make([]int, 9)
	for i := 0; i < 9; i++ {
		digits[i] = shared.Rng.Intn(10)
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
