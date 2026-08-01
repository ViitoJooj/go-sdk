package company

import (
	"errors"
	"fmt"
)

func TradeName(name string) error {
	if len(name) == 0 {
		return errors.New("trade name cannot be empty.")
	}
	if len(name) > 255 {
		return fmt.Errorf("trade name cannot exceed %d characters. (current %d)", 255, len(name))
	}
	return nil
}
