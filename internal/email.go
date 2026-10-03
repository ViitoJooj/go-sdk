package internal

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

//go:embed disposable_domains.txt
var embeddedDisposableDomains []byte

var (
	disposableDomainsMu sync.RWMutex
	disposableDomains   map[string]struct{}
)

func init() {
	domains, _ := parseDisposableDomains(bytes.NewReader(embeddedDisposableDomains))
	disposableDomains = domains
}

func parseDisposableDomains(r io.Reader) (map[string]struct{}, error) {
	domains := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		domain := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if domain != "" {
			domains[domain] = struct{}{}
		}
	}

	return domains, scanner.Err()
}

// IsDisposable reports whether mail's domain is a known disposable-email
// domain. The list ships embedded in the binary: importing this package
// never makes a network call. Use RefreshDisposableDomains to update the
// list at runtime, explicitly.
func IsDisposable(mail string) (bool, error) {
	parts := strings.Split(mail, "@")
	if len(parts) != 2 {
		return false, nil
	}

	domain := strings.ToLower(strings.TrimSpace(parts[1]))

	disposableDomainsMu.RLock()
	_, exists := disposableDomains[domain]
	disposableDomainsMu.RUnlock()

	return exists, nil
}

// DefaultDisposableDomainsURL is the source the embedded list was built
// from, for convenience when calling RefreshDisposableDomains.
const DefaultDisposableDomainsURL = "https://disposable.github.io/disposable-email-domains/domains.txt"

// RefreshDisposableDomains fetches an updated disposable-domains list from
// url and swaps it in atomically. It is never called implicitly — only an
// explicit call does network I/O, and it respects ctx's deadline.
func RefreshDisposableDomains(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status fetching disposable domains: %s", resp.Status)
	}

	domains, err := parseDisposableDomains(resp.Body)
	if err != nil {
		return err
	}

	disposableDomainsMu.Lock()
	disposableDomains = domains
	disposableDomainsMu.Unlock()

	return nil
}
