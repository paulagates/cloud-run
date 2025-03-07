package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paulagates/cloud-run/internal/service"
)

type WeatherHandler struct {
	service service.WeatherService
}

func (we *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cep := q.Get("cep")
	viacep, err := we.service.FindLocation(cep)
	if err != nil {
		if err == service.ErrZipcodeNotFound {
			http.Error(w, service.ErrZipcodeNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, service.ErrInvalidZipcode.Error(), http.StatusUnprocessableEntity)
		return
	}
	weather, err := we.service.GetCurrentWeather(viacep.Localidade)
	if err != nil {
		http.Error(w, service.ErrWeatherAPI.Error(), http.StatusInternalServerError)
		return
	}
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		log.Println(err)
		http.Error(w, service.ErrWeatherJson.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(weatherJson)
}

func NewWeatherHandler(service service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service}
}
