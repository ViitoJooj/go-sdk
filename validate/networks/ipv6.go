package networks

import (
	"errors"
	"net"
	"strings"
)

func Ipv6(ip string) error {
	ip = strings.TrimSpace(ip)
	if ip == "" {
		return errors.New("IPv6 cannot be empty.")
	}
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return errors.New("invalid IPv6 address.")
	}
	if parsed.To4() != nil {
		return errors.New("address is IPv4, not IPv6.")
	}
	return nil
}
