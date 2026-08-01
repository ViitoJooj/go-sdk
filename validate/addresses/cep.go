package addresses

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ViitoJooj/go-sdk/internal"
)

type cepError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func CEP(cep string) error {
	digits := internal.StripNonDigits(cep)

	if len(digits) == 0 {
		return errors.New("CEP cannot be empty.")
	}
	if len(digits) != 8 {
		return fmt.Errorf("CEP must have 8 digits. (current %d)", len(digits))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", digits)

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao consultar Brasil API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {
		var apiErr cepError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = "CEP nao encontrado"
		}
		return errors.New(msg)
	}

	return fmt.Errorf("Brasil API retornou status %d", resp.StatusCode)
}

func CEPv2(cep string) error {
	digits := internal.StripNonDigits(cep)

	if len(digits) == 0 {
		return errors.New("CEP cannot be empty.")
	}
	if len(digits) != 8 {
		return fmt.Errorf("CEP must have 8 digits. (current %d)", len(digits))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v2/%s", digits)

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao consultar Brasil API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {
		var apiErr cepError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = "CEP nao encontrado"
		}
		return errors.New(msg)
	}

	return fmt.Errorf("Brasil API retornou status %d", resp.StatusCode)
}
