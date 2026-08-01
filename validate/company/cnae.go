package company

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

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
