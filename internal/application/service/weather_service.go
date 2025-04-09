package service

import (
	"fmt"

	"github.com/fmantinossi/weather-app/internal/domain"
)

type WeatherService struct {
	addressProvider domain.AddressProvider
	weatherProvider domain.WeatherProvider
	apiKey          string
}

func NewWeatherService(
	addressProvider domain.AddressProvider,
	weatherProvider domain.WeatherProvider,
	apiKey string,
) *WeatherService {
	return &WeatherService{
		addressProvider: addressProvider,
		weatherProvider: weatherProvider,
		apiKey:          apiKey,
	}
}

func (s *WeatherService) GetWeather(cep string) (*domain.Wheater, error) {
	address, err := s.addressProvider.GetAddress(cep)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	weather, err := s.weatherProvider.GetWeather(s.apiKey, address.Location.Coordinates.Latitude, address.Location.Coordinates.Longitude)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &domain.Wheater{
		Celsius:    weather.Current.TempC,
		Fahrenheit: weather.Current.TempF,
		Kelvin:     weather.Current.TempC + 273,
	}, nil
}
