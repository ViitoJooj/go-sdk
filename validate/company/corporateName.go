package company

import (
	"errors"
	"fmt"
)

func CorporateName(name string) error {
	if len(name) == 0 {
		return errors.New("corporate name cannot be empty.")
	}
	if len(name) > 255 {
		return fmt.Errorf("corporate name cannot exceed %d characters. (current %d)", 255, len(name))
	}
	return nil
}
