package company

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

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
