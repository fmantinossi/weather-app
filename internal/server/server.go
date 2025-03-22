package server

import (
	"github.com/fmantinossi/weather-app/internal/adapters"
	"github.com/fmantinossi/weather-app/internal/application/service"
	"github.com/fmantinossi/weather-app/internal/handler"
	"github.com/fmantinossi/weather-app/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) Setup() {
	apiKey := "f2c9ca7e8eec4a2b901190619252003"

	addressAdapter := adapters.NewBrasilApiAdapter()
	weatherAdapter := adapters.NewWeatherApiAdapter()

	weatherService := service.NewWeatherService(addressAdapter, weatherAdapter, apiKey)

	weatherHandler := handler.NewWeatherHandler(weatherService)

	router.SetupRoutes(s.engine, weatherHandler)
}

func (s *Server) Start(addr string) error {
	return s.engine.Run(addr)
}
