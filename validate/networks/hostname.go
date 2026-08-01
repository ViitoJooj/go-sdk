package networks

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var hostnameRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)

func Hostname(hostname string) error {
	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return errors.New("hostname cannot be empty.")
	}
	if len(hostname) > 253 {
		return fmt.Errorf("hostname cannot exceed %d characters. (current %d)", 253, len(hostname))
	}
	if !hostnameRegex.MatchString(hostname) {
		return errors.New("invalid hostname format.")
	}
	return nil
}
