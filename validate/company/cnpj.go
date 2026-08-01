package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ViitoJooj/go-sdk/internal"
)

func CNPJ(cnpj string) error {
	var errs []error
	digits := internal.StripNonDigits(cnpj)

	if len(digits) == 0 {
		return errors.New("CNPJ cannot be empty.")
	}

	if len(digits) != 14 {
		errs = append(errs, fmt.Errorf("CNPJ must have 14 digits. (current %d)", len(digits)))
	}

	if len(digits) == 14 && internal.AllSameDigit(digits) {
		errs = append(errs, errors.New("CNPJ cannot have all identical digits."))
	}

	if len(digits) == 14 && !internal.AllSameDigit(digits) && !internal.CNPJCheckDigitsValid(digits) {
		errs = append(errs, errors.New("CNPJ check digits are invalid."))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	client := &http.Client{Timeout: 10 * time.Second}
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
		var apiErr internal.BrasilAPIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = "CNPJ nao encontrado na base da Receita Federal"
		}
		return errors.New(msg)
	}

	return fmt.Errorf("Brasil API retornou status %d", resp.StatusCode)
}
