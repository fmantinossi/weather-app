package service

import (
	"testing"

	"github.com/fmantinossi/weather-app/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockAddressProvider struct{}

func (m *mockAddressProvider) GetAddress(cep string) (*domain.Address, error) {
	return &domain.Address{
		Cep:   "01001-000",
		City:  "São Paulo",
		State: "SP",
		Location: struct {
			Type        string `json:"type"`
			Coordinates struct {
				Longitude string `json:"longitude"`
				Latitude  string `json:"latitude"`
			} `json:"coordinates"`
		}{
			Type: "Point",
			Coordinates: struct {
				Longitude string `json:"longitude"`
				Latitude  string `json:"latitude"`
			}{
				Latitude:  "-23.55052",
				Longitude: "-46.633308",
			},
		},
	}, nil
}

type mockWeatherProvider struct{}

func (m *mockWeatherProvider) GetWeather(apiKey, lat, lon string) (*domain.WeatherResponse, error) {
	return &domain.WeatherResponse{
		Current: struct {
			TempC     float64 `json:"temp_c"`
			TempF     float64 `json:"temp_f"`
			Condition struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
			} `json:"condition"`
			Humidity int     `json:"humidity"`
			WindKph  float64 `json:"wind_kph"`
		}{
			TempC: 25.0,
			TempF: 77.0,
			Condition: struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
			}{
				Text: "Ensolarado",
				Icon: "//cdn.weatherapi.com/icon.png",
			},
			Humidity: 60,
			WindKph:  10.0,
		},
	}, nil
}

func TestGetWeather_Success(t *testing.T) {
	addressMock := &mockAddressProvider{}
	weatherMock := &mockWeatherProvider{}
	apiKey := "fake-key"

	svc := NewWeatherService(addressMock, weatherMock, apiKey)

	result, err := svc.GetWeather("01001000")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 25.0, result.Celsius)
	assert.Equal(t, 77.0, result.Fahrenheit)
	assert.Equal(t, 298.0, result.Kelvin) // 25 + 273
}
