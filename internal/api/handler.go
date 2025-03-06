package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/paulagates/cloud-run/internal/client"
	"github.com/paulagates/cloud-run/internal/service"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cepParam := q.Get("cep")
	if len(cepParam) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	cep, err := strconv.Atoi(cepParam)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	viacep, err := client.FindLocation(cep)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	if viacep == (client.Address{}) {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	weather, err := client.GetCurrentWeather(viacep.Localidade)
	if err != nil {
		http.Error(w, "weather api error", http.StatusInternalServerError)
		return
	}
	weather.TempK = service.ConvertCelsiusToKelvin(weather.TempC)
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		http.Error(w, "error converting weather to json", http.StatusInternalServerError)
		return
	}
	w.Write(weatherJson)
}
