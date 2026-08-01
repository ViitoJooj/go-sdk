package networks

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func URL(raw string) error {
	if strings.TrimSpace(raw) == "" {
		return errors.New("URL cannot be empty.")
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("URL scheme must be http or https.")
	}
	if u.Host == "" {
		return errors.New("URL must contain a host.")
	}
	return nil
}
