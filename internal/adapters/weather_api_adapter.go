package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		Humidity int     `json:"humidity"`
		WindKph  float64 `json:"wind_kph"`
	} `json:"current"`
}

func NewWeatherApiAdapter() *WeatherResponse {
	return &WeatherResponse{}
}

func (w *WeatherResponse) GetWeather(apiKey, lat, lon string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s", apiKey, lat, lon)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("WeatherApi returns error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("WeatherApi returns error: status %d", resp.StatusCode)
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("Error decoding WeatherApi JSON: %v", err)
	}

	return &weatherData, nil
}
