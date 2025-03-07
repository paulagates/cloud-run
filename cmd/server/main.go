package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/paulagates/cloud-run/internal/api"
	"github.com/paulagates/cloud-run/internal/service"
)

func main() {
	godotenv.Load()
	weatherService := service.NewWeatherService()
	api := api.NewWeatherHandler(weatherService)
	http.HandleFunc("GET /", api.GetWeather)
	http.HandleFunc("GET /weather", api.GetWeather)
	http.ListenAndServe(":8080", nil)
}
