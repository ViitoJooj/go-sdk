package company

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func CNPJ(cnpj string) error {
	var errs []error
	digits := internal.StripNonDigits(cnpj)

	if len(digits) == 0 {
		return errors.New("CNPJ cannot be empty.")
	}

	if len(digits) != 14 {
		errs = append(errs, fmt.Errorf("CNPJ must have 14 digits. (current %d)", len(digits)))
	}

	if len(digits) == 14 && internal.AllSameDigit(digits) {
		errs = append(errs, errors.New("CNPJ cannot have all identical digits."))
	}

	if len(digits) == 14 && !internal.AllSameDigit(digits) && !internal.CNPJCheckDigitsValid(digits) {
		errs = append(errs, errors.New("CNPJ check digits are invalid."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func CorporateName(name string) error {
	if len(name) == 0 {
		return errors.New("corporate name cannot be empty.")
	}
	if len(name) > 255 {
		return fmt.Errorf("corporate name cannot exceed %d characters. (current %d)", 255, len(name))
	}
	return nil
}

func TradeName(name string) error {
	if len(name) == 0 {
		return errors.New("trade name cannot be empty.")
	}
	if len(name) > 255 {
		return fmt.Errorf("trade name cannot exceed %d characters. (current %d)", 255, len(name))
	}
	return nil
}

func IM(im string) error {
	im = internal.StripNonDigits(im)
	if len(im) == 0 {
		return errors.New("IM (Inscricao Municipal) cannot be empty.")
	}
	if len(im) > 20 {
		return fmt.Errorf("IM cannot exceed %d digits. (current %d)", 20, len(im))
	}
	return nil
}

func IE(ie string) error {
	ie = internal.StripNonDigits(ie)
	if len(ie) == 0 {
		return errors.New("IE (Inscricao Estadual) cannot be empty.")
	}
	if len(ie) > 20 {
		return fmt.Errorf("IE cannot exceed %d digits. (current %d)", 20, len(ie))
	}
	return nil
}

func CNAE(cnae string) error {
	var errs []error
	digits := internal.StripNonDigits(cnae)

	if len(digits) == 0 {
		return errors.New("CNAE cannot be empty.")
	}
	if len(digits) != 7 {
		errs = append(errs, fmt.Errorf("CNAE must have 7 digits. (current %d)", len(digits)))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
