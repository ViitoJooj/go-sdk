package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ViitoJooj/go-sdk/internal"
)

type brasilAPIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func CNPJWithBrasilAPI(cnpj string) error {
	digits := internal.StripNonDigits(cnpj)

	if err := CNPJ(digits); err != nil {
		return fmt.Errorf("formato invalido: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	return consultaBrasilAPI(digits, client)
}

func CNPJWithBrasilAPIClient(cnpj string, client *http.Client) error {
	digits := internal.StripNonDigits(cnpj)

	if err := CNPJ(digits); err != nil {
		return fmt.Errorf("formato invalido: %w", err)
	}

	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	return consultaBrasilAPI(digits, client)
}

func consultaBrasilAPI(digits string, client *http.Client) error {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%s", digits)

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao consultar Brasil API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {
		var apiErr brasilAPIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = "CNPJ nao encontrado na base da Receita Federal"
		}
		return errors.New(msg)
	}

	return fmt.Errorf("Brasil API retornou status %d", resp.StatusCode)
}
