package service

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	ErrInvalidZipcode  = errors.New("invalid zipcode")
	ErrZipcodeNotFound = errors.New("can not find zipcode")
	ErrWeatherAPI      = errors.New("weather api error")
	ErrWeatherJson     = errors.New("error converting weather to json")
)

type WeatherService interface {
	FindLocation(cep string) (Address, error)
	GetCurrentWeather(city string) (Weather, error)
}

type weatherService struct {
}

func (c *weatherService) FindLocation(cepStr string) (Address, error) {
	if len(cepStr) != 8 {
		return Address{}, ErrInvalidZipcode
	}
	_, err := strconv.Atoi(cepStr)
	if err != nil {
		return Address{}, ErrInvalidZipcode
	}
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cepStr)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	res, err := client.Get(url)
	if err != nil {
		return Address{}, err
	}
	defer res.Body.Close()
	var address Address
	err = json.NewDecoder(res.Body).Decode(&address)
	if err != nil {
		return Address{}, err
	}

	if address.Cep == "" {
		return Address{}, ErrZipcodeNotFound
	}
	return address, nil
}

func (c *weatherService) GetCurrentWeather(city string) (Weather, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", os.Getenv("WEATHER_API_KEY"), url.QueryEscape(city))
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	res, err := client.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer res.Body.Close()
	var responseWeather ResponseWeather
	err = json.NewDecoder(res.Body).Decode(&responseWeather)
	if err != nil {
		return Weather{}, err
	}
	responseWeather.Current.TempK = responseWeather.Current.TempC + 273
	return responseWeather.Current, nil
}

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Weather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type ResponseWeather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current Weather `json:"current"`
}

func NewWeatherService() WeatherService {
	return &weatherService{}
}
