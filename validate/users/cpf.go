package users

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func CPF(cpf string) error {
	var errs []error
	digits := internal.StripNonDigits(cpf)

	if len(digits) == 0 {
		return errors.New("CPF cannot be empty.")
	}

	if len(digits) != 11 {
		errs = append(errs, fmt.Errorf("CPF must have 11 digits. (current %d)", len(digits)))
	}

	if len(digits) == 11 && internal.AllSameDigit(digits) {
		errs = append(errs, errors.New("CPF cannot have all identical digits."))
	}

	if len(digits) == 11 && !internal.AllSameDigit(digits) && !internal.CPFCheckDigitsValid(digits) {
		errs = append(errs, errors.New("CPF check digits are invalid."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
