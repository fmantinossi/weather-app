package main

import (
	"github.com/fmantinossi/weather-app/internal/adapters"
	"github.com/fmantinossi/weather-app/internal/application/service"
	"github.com/fmantinossi/weather-app/internal/http"
)

func main() {
	brasilApiAdapter := adapters.NewBrasilApiAdapter()
	weatherApiAdapter := adapters.NewWeatherApiAdapter()
	service := service.NewWeatherService(brasilApiAdapter, weatherApiAdapter)
	handler := http.NewWeatherHandler(service)
	server := http.NewServer(handler)
	server.Start()
}
