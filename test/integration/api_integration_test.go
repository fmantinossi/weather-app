package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fmantinossi/weather-app/internal/adapters"
	"github.com/fmantinossi/weather-app/internal/application/service"
	"github.com/fmantinossi/weather-app/internal/handler"
	"github.com/fmantinossi/weather-app/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type weatherResponse struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

func setupTestServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	engine := gin.Default()

	apiKey := "f2c9ca7e8eec4a2b901190619252003"

	addressAdapter := adapters.NewBrasilApiAdapter()
	weatherAdapter := adapters.NewWeatherApiAdapter()
	weatherService := service.NewWeatherService(addressAdapter, weatherAdapter, apiKey)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	router.SetupRoutes(engine, weatherHandler)

	return engine
}

func TestGetWeather_Success(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/weather/01001000", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var body weatherResponse
	err := json.Unmarshal(resp.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Greater(t, body.Celsius, 0.0)
	assert.Greater(t, body.Kelvin, 0.0)
}

func TestGetWeather_InvalidCepFormat(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/weather/abc", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestGetWeather_CepNotFound(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/weather/99999999", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
