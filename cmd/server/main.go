package main

import (
	"log"

	"github.com/fmantinossi/weather-app/internal/server"
)

func main() {
	s := server.NewServer()
	s.Setup()

	log.Println("🚀 Servidor iniciado em http://localhost:8080")
	if err := s.Start(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
