package networks

import (
	"errors"
	"net"
	"strings"
)

func Ipv4(ip string) error {
	ip = strings.TrimSpace(ip)
	if ip == "" {
		return errors.New("IPv4 cannot be empty.")
	}
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return errors.New("invalid IPv4 address.")
	}
	if parsed.To4() == nil {
		return errors.New("not a valid IPv4 address.")
	}
	return nil
}
