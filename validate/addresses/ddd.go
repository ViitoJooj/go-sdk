package addresses

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ViitoJooj/go-sdk/internal"
)

type dddError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func DDD(ddd string) error {
	digits := internal.StripNonDigits(ddd)

	if len(digits) == 0 {
		return errors.New("DDD cannot be empty.")
	}
	if len(digits) != 2 {
		return fmt.Errorf("DDD must have 2 digits. (current %d)", len(digits))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://brasilapi.com.br/api/ddd/v1/%s", digits)

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao consultar Brasil API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {
		var apiErr dddError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = "DDD nao encontrado"
		}
		return errors.New(msg)
	}

	return fmt.Errorf("Brasil API retornou status %d", resp.StatusCode)
}
