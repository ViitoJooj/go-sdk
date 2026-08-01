package company

import (
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func CNPJ(cnpj string) string {
	digits := internal.StripNonDigits(cnpj)
	if len(digits) != 14 {
		return cnpj
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", digits[0:2], digits[2:5], digits[5:8], digits[8:12], digits[12:14])
}
