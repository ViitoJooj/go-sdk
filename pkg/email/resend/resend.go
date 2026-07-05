// Package resend is a minimal HTTP client for the Resend transactional email API.
package resend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.resend.com"

type Client struct {
	apiKey string
	from   string
}

func New(apiKey, from string) *Client {
	return &Client{apiKey: apiKey, from: from}
}

func (c *Client) Send(to, subject, html string) error {
	payload := map[string]interface{}{
		"from":    c.from,
		"to":      []string{to},
		"subject": subject,
		"html":    html,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, baseURL+"/emails", bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("resend: status %d", resp.StatusCode)
	}
	return nil
}
