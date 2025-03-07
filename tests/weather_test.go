package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paulagates/cloud-run/internal/api"
	"github.com/paulagates/cloud-run/internal/service"
)

func TestGetWeather(t *testing.T) {
	tests := []struct {
		name         string
		cep          string
		mockFindLoc  service.Address
		mockWeather  service.Weather
		mockFindErr  error
		mockWeaErr   error
		expectedCode int
		expectedBody string
	}{
		{
			name: "Successful request",
			cep:  "05114100",
			mockFindLoc: service.Address{
				Localidade: "São Paulo",
			},
			mockWeather: service.Weather{
				TempC: 25,
				TempF: 77,
				TempK: 298,
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"temp_c":25,"temp_f":77,"temp_k":298}` + "\n",
		},
		{
			name:         "Invalid ZIP code",
			cep:          "00000000",
			mockFindErr:  service.ErrInvalidZipcode,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: "invalid zipcode\n",
		},
		{
			name:         "ZIP code not found",
			cep:          "99999999",
			mockFindErr:  service.ErrZipcodeNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "can not find zipcode\n",
		},
		{
			name: "Weather API error",
			cep:  "05114100",
			mockFindLoc: service.Address{
				Localidade: "São Paulo",
			},
			mockWeaErr:   service.ErrWeatherAPI,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "weather api error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weatherService := &mockWeatherService{
				mockFindLoc: tt.mockFindLoc,
				mockWeather: tt.mockWeather,
				mockFindErr: tt.mockFindErr,
				mockWeaErr:  tt.mockWeaErr,
			}
			weatherHandler := api.NewWeatherHandler(weatherService)

			req := httptest.NewRequest("GET", "/weather?cep="+tt.cep, nil)
			rr := httptest.NewRecorder()

			weatherHandler.GetWeather(rr, req)
			if rr.Code != tt.expectedCode {
				t.Errorf("unexpected status code: got %v, want %v", rr.Code, tt.expectedCode)
			}

			if tt.expectedCode == http.StatusOK {
				var got service.Weather
				if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				expected := tt.mockWeather

				if !reflect.DeepEqual(got, expected) {
					t.Errorf("unexpected response body: got %+v, want %+v", got, expected)
				}
			} else if rr.Body.String() != tt.expectedBody {
				t.Errorf("unexpected response body: got %q, want %q", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

type mockWeatherService struct {
	mockFindLoc service.Address
	mockWeather service.Weather
	mockFindErr error
	mockWeaErr  error
}

func (m *mockWeatherService) FindLocation(cep string) (service.Address, error) {
	return m.mockFindLoc, m.mockFindErr
}

func (m *mockWeatherService) GetCurrentWeather(localidade string) (service.Weather, error) {
	return m.mockWeather, m.mockWeaErr
}
