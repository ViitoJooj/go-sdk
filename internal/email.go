package internal

import (
	"bufio"
	"net/http"
	"strings"
	"sync"
)

var (
	disposableDomains map[string]struct{}
	domainsOnce       sync.Once
	domainsLoadErr    error
)

func loadDisposableDomains() {
	disposableDomains = make(map[string]struct{})

	resp, err := http.Get("https://disposable.github.io/disposable-email-domains/domains.txt")
	if err != nil {
		domainsLoadErr = err
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		domain := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if domain != "" {
			disposableDomains[domain] = struct{}{}
		}
	}

	domainsLoadErr = scanner.Err()
}

func init() {
	go domainsOnce.Do(loadDisposableDomains)
}

func IsDisposable(mail string) (bool, error) {
	domainsOnce.Do(loadDisposableDomains)

	if domainsLoadErr != nil {
		return false, domainsLoadErr
	}

	parts := strings.Split(mail, "@")
	if len(parts) != 2 {
		return false, nil
	}

	domain := strings.ToLower(strings.TrimSpace(parts[1]))
	_, exists := disposableDomains[domain]
	return exists, nil
}
