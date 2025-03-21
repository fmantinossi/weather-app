package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine         *gin.Engine
	weatherHandler *WeatherHandler
}

func (s *Server) Start() {
	addr := ":8080"
	log.Printf("Server is up in http://localhost:%s", addr)
	if err := s.engine.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func (s *Server) setupRoutes() {
	s.engine.GET("/cep/:cep", s.weatherHandler.GetWeatherForecast)
}

func NewServer(weatherHandler *WeatherHandler) *Server {
	engine := gin.Default()
	server := &Server{
		engine:         engine,
		weatherHandler: weatherHandler,
	}
	server.setupRoutes()
	return server
}
