package company

import (
	"fmt"

	"github.com/ViitoJooj/go-sdk/fakeData/shared"
)

func GenValidCNPJ() string {
	digits := make([]int, 12)
	for i := 0; i < 12; i++ {
		digits[i] = shared.Rng.Intn(10)
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
