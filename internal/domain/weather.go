package domain

type Wheater struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

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
