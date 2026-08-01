package networks

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func Domain(domain string) error {
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return errors.New("domain cannot be empty.")
	}
	if len(domain) > 253 {
		return fmt.Errorf("domain cannot exceed %d characters. (current %d)", 253, len(domain))
	}
	if !domainRegex.MatchString(domain) {
		return errors.New("invalid domain format.")
	}
	return nil
}
