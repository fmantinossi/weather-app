package domain

type AddressProvider interface {
	GetAddress(cep string) (*Address, error)
}

type WeatherProvider interface {
	GetWeather(apiKey, lat, lon string) (*WeatherResponse, error)
}

type WeatherServiceInterface interface {
	GetWeather(cep string) (*Wheater, error)
}
