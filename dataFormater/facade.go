package dataFormater

import (
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func CPF(cpf string) string {
	digits := internal.StripNonDigits(cpf)
	if len(digits) != 11 {
		return cpf
	}
	return fmt.Sprintf("%s.%s.%s-%s", digits[0:3], digits[3:6], digits[6:9], digits[9:11])
}

func CNPJ(cnpj string) string {
	digits := internal.StripNonDigits(cnpj)
	if len(digits) != 14 {
		return cnpj
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", digits[0:2], digits[2:5], digits[5:8], digits[8:12], digits[12:14])
}

func Phone(phone string) string {
	digits := internal.StripNonDigits(phone)
	l := len(digits)
	switch {
	case l == 11:
		return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:7], digits[7:11])
	case l == 10:
		return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:6], digits[6:10])
	case l > 11:
		return fmt.Sprintf("+%s (%s) %s-%s", digits[0:l-11], digits[l-11:l-9], digits[l-9:l-4], digits[l-4:l])
	default:
		return phone
	}
}

func PostalCode(postal string) string {
	digits := internal.StripNonDigits(postal)
	if len(digits) == 8 {
		return fmt.Sprintf("%s-%s", digits[0:5], digits[5:8])
	}
	return postal
}
