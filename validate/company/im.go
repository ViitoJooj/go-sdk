package company

import (
	"errors"
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

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
