package router

import (
	"github.com/fmantinossi/weather-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, weatherHandler *handler.WeatherHandler) {
	weather := r.Group("/weather")
	{
		weather.GET("/:cep", weatherHandler.GetWeatherForecast)
	}
}
