package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fmantinossi/weather-app/internal/domain"
)

// WeatherApiAdapter implementa domain.WeatherProvider
type WeatherApiAdapter struct {
	BaseURL string
	Client  *http.Client
}

func NewWeatherApiAdapter() *WeatherApiAdapter {
	return &WeatherApiAdapter{
		BaseURL: "https://api.weatherapi.com/v1",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetWeather busca informações climáticas via WeatherAPI
func (w *WeatherApiAdapter) GetWeather(apiKey, lat, lon string) (*domain.WeatherResponse, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s,%s", w.BaseURL, apiKey, lat, lon)

	resp, err := w.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("weatherapi: error in request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weatherapi: error status code: %d", resp.StatusCode)
	}

	var result domain.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("weatherapi: error decoding JSON: %w", err)
	}

	return &result, nil
}
