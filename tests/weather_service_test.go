package tests

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/paulagates/cloud-run/internal/service"
)

func TestFindLocation(t *testing.T) {
	weatherService := service.NewWeatherService()
	t.Parallel()
	tests := []struct {
		name    string
		cep     string
		address service.Address
		wantErr error
	}{
		{
			name:    "Find São Paulo",
			cep:     "05114100",
			address: service.Address{Localidade: "São Paulo"},
		},
		{
			name:    "Find Invalid Zipcode",
			cep:     "1234567",
			address: service.Address{},
			wantErr: service.ErrInvalidZipcode,
		},
		{
			name:    "Find Zipcode Not Found",
			cep:     "00000000",
			address: service.Address{},
			wantErr: service.ErrZipcodeNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := weatherService.FindLocation(tt.cep)
			if err != tt.wantErr {
				t.Errorf("FindLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Localidade != tt.address.Localidade {
				t.Errorf("FindLocation() = %v, want %v", got.Localidade, tt.address.Localidade)
			}
		})
	}
}
func TestGetCurrentWeather(t *testing.T) {
	godotenv.Load("../.env")
	weatherService := service.NewWeatherService()
	t.Run("Find São Paulo Weather", func(t *testing.T) {
		t.Parallel()
		city := "São Paulo"
		notExpectedResponse := service.Weather{}
		got, err := weatherService.GetCurrentWeather(city)
		if err != nil {
			t.Fatal(err)
		}
		if got == notExpectedResponse {
			t.Errorf("FindLocation() = %v, want %v", got, notExpectedResponse)
		}
	})
	tests := []struct {
		name    string
		city    string
		weather service.Weather
		wantErr error
	}{
		{
			name: "Get Weather for Empty City",
			city: "",
			weather: service.Weather{
				TempC: 0.0,
				TempF: 0.0,
				TempK: 0.0,
			},
			wantErr: service.ErrWeatherAPI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := weatherService.GetCurrentWeather(tt.city)
			if err != tt.wantErr {
				t.Errorf("GetCurrentWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.weather {
				t.Errorf("GetCurrentWeather() = %v, want %v", got.TempC, tt.weather)
			}
		})
	}
}
