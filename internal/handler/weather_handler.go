package handler

import (
	"errors"
	"net/http"

	"github.com/fmantinossi/weather-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	service domain.WeatherServiceInterface
}

type Handler interface {
	GetWeatherForecast(c *gin.Context)
}

func NewWeatherHandler(service domain.WeatherServiceInterface) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetWeatherForecast(c *gin.Context) {
	cep := c.Param("cep")

	weather, err := h.service.GetWeather(cep)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		case errors.Is(err, domain.ErrUnprocessableEntity):
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, weather)
}
