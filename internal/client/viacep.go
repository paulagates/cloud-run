package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func FindLocation(cep int) (Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%d/json/", cep)
	res, err := http.Get(url)
	if err != nil {
		return Address{}, err
	}
	defer res.Body.Close()
	var address Address
	err = json.NewDecoder(res.Body).Decode(&address)
	if err != nil {
		return Address{}, err
	}
	return address, nil
}
