package http

import (
	"errors"
	"net/http"

	"github.com/fmantinossi/weather-app/internal/application/service"
	"github.com/fmantinossi/weather-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	weatherService *service.WheaterService
}

func NewWeatherHandler(weatherService *service.WheaterService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

func (h *WeatherHandler) GetWeatherForecast(c *gin.Context) {
	weather, err := h.weatherService.GetWeather(c.Param("cep"))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"Message": err.Error()})
			return
		case errors.Is(err, domain.ErrUnprocessableEntity):
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, weather)
}
