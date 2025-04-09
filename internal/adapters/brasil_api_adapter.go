package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fmantinossi/weather-app/internal/domain"
)

type BrasilApiAdapter struct {
	BaseURL string
	Client  *http.Client
}

func NewBrasilApiAdapter() *BrasilApiAdapter {
	return &BrasilApiAdapter{
		BaseURL: "https://brasilapi.com.br",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (b *BrasilApiAdapter) GetAddress(cep string) (*domain.Address, error) {
	url := fmt.Sprintf("%s/api/cep/v2/%s", b.BaseURL, cep)

	resp, err := b.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("BrasilAPI request error: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, domain.ErrNotFound
	case http.StatusUnprocessableEntity:
		return nil, domain.ErrUnprocessableEntity
	case http.StatusBadRequest:
		return nil, domain.ErrUnprocessableEntity
	default:
		return nil, fmt.Errorf("BrasilAPI unexpected status code: %d", resp.StatusCode)
	}

	var adr domain.Address
	if err := json.NewDecoder(resp.Body).Decode(&adr); err != nil {
		return nil, fmt.Errorf("BrasilAPI decode error: %w", err)
	}

	return &adr, nil
}
