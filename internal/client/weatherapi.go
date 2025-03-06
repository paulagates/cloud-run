package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CurrentWeather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type responseWeather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current CurrentWeather `json:"current"`
}

func GetCurrentWeather(city string) (CurrentWeather, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", os.Getenv("WEATHER_API_KEY"), city)
	res, err := http.Get(url)
	if err != nil {
		return CurrentWeather{}, err
	}
	defer res.Body.Close()
	var responseWeather responseWeather
	err = json.NewDecoder(res.Body).Decode(&responseWeather)
	if err != nil {
		return CurrentWeather{}, err
	}
	return responseWeather.Current, nil
}
