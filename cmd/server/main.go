package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/paulagates/cloud-run/internal/api"
)

func main() {
	godotenv.Load()
	http.HandleFunc("GET /weather", api.GetWeather)
	http.ListenAndServe(":8080", nil)
}
