package service

import (
	"fmt"

	"github.com/fmantinossi/weather-app/internal/adapters"
	"github.com/fmantinossi/weather-app/internal/domain"
)

type WheaterService struct {
	address *adapters.Address
	weather *adapters.WeatherResponse
}

func NewWeatherService(brasilApiAdapter *adapters.Address, weatherApiAdapter *adapters.WeatherResponse) *WheaterService {
	return &WheaterService{
		address: brasilApiAdapter,
		weather: weatherApiAdapter,
	}
}

func (s WheaterService) GetWeather(cep string) (*domain.Wheater, error) {
	adr, err := s.address.GetAddress(cep)
	if err != nil {
		return nil, fmt.Errorf("GetAddress returns error: %v", err)
	}

	wthr, err := s.weather.GetWeather("f2c9ca7e8eec4a2b901190619252003", adr.Location.Coordinates.Latitude, adr.Location.Coordinates.Longitude)
	if err != nil {
		return nil, fmt.Errorf("GetWheather returns error: %v", err)
	}

	return &domain.Wheater{
		Celsius:    wthr.Current.TempC,
		Fahrenheit: wthr.Current.TempF,
		Kelvin:     wthr.Current.TempC + 273,
	}, nil
}
