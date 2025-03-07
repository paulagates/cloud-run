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
		name string
		cep  string
		want service.Address
	}{
		{
			name: "Find São Paulo",
			cep:  "05114100",
			want: service.Address{Localidade: "São Paulo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := weatherService.FindLocation(tt.cep)
			if err != nil {
				t.Fatal(err)
			}
			if got.Localidade != tt.want.Localidade {
				t.Errorf("FindLocation() = %v, want %v", got.Localidade, tt.want)
			}
		})
	}

}

func TestGetCurrentWeather(t *testing.T) {
	godotenv.Load()
	weatherService := service.NewWeatherService()
	name := "Find São Paulo Weather"
	city := "São Paulo"
	notExpectedResponse := service.Weather{}
	t.Parallel()
	t.Run(name, func(t *testing.T) {
		got, err := weatherService.GetCurrentWeather(city)
		if err != nil {
			t.Fatal(err)
		}
		if got == notExpectedResponse {
			t.Errorf("FindLocation() = %v, want %v", got, notExpectedResponse)
		}
	})

}
