package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fmantinossi/weather-app/internal/domain"
)

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Location     struct {
		Type        string `json:"type"`
		Coordinates struct {
			Longitude string `json:"longitude"`
			Latitude  string `json:"latitude"`
		} `json:"coordinates"`
	} `json:"location"`
}

func NewBrasilApiAdapter() *Address {
	return &Address{}
}

func (a *Address) GetAddress(cep string) (*Address, error) {
	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v2/%s", cep)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("BrasilApi returns error: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// Continua normalmente

	case http.StatusNotFound: // 404
		return nil, domain.ErrNotFound

	case http.StatusUnprocessableEntity: // 422
		return nil, domain.ErrUnprocessableEntity

	default:
		return nil, fmt.Errorf("An error occurred while processing the request: status %d", resp.StatusCode)
	}

	var adr Address
	if err := json.NewDecoder(resp.Body).Decode(&adr); err != nil {
		return nil, fmt.Errorf("Error decoding BrasilApi JSON: %v", err)
	}
	log.Printf("[LOG] BrasilApi response: %s", adr)
	return &adr, nil
}
