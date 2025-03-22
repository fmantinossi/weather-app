package adapters

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fmantinossi/weather-app/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetWeather_Success(t *testing.T) {
	mockResponse := domain.WeatherResponse{
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
			WindKph:  10,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	adapter := &WeatherApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	result, err := adapter.GetWeather("dummy_key", "1.23", "4.56")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 25.0, result.Current.TempC)
	assert.Equal(t, "Ensolarado", result.Current.Condition.Text)
}

func TestGetWeather_StatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "error", http.StatusInternalServerError)
	}))
	defer server.Close()

	adapter := &WeatherApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	result, err := adapter.GetWeather("key", "lat", "lon")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "status de erro")
}

func TestGetWeather_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	adapter := &WeatherApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	result, err := adapter.GetWeather("key", "lat", "lon")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "decodificar JSON")
}
