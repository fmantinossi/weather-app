package adapters

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fmantinossi/weather-app/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress_Success(t *testing.T) {
	expected := domain.Address{
		Cep:          "01001-000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
		Service:      "brasilapi",
	}
	expected.Location.Type = "Point"
	expected.Location.Coordinates.Latitude = "-23.55052"
	expected.Location.Coordinates.Longitude = "-46.633308"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	adapter := &BrasilApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	result, err := adapter.GetAddress("01001000")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expected.City, result.City)
	assert.Equal(t, expected.Location.Coordinates.Latitude, result.Location.Coordinates.Latitude)
}

func TestGetAddress_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	adapter := &BrasilApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	addr, err := adapter.GetAddress("00000000")

	assert.ErrorIs(t, err, domain.ErrNotFound)
	assert.Nil(t, addr)
}

func TestGetAddress_UnprocessableEntity(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "invalid cep", http.StatusUnprocessableEntity)
	}))
	defer server.Close()

	adapter := &BrasilApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	addr, err := adapter.GetAddress("abc")

	assert.ErrorIs(t, err, domain.ErrUnprocessableEntity)
	assert.Nil(t, addr)
}

func TestGetAddress_UnexpectedStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer server.Close()

	adapter := &BrasilApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	addr, err := adapter.GetAddress("12345678")

	assert.Error(t, err)
	assert.Nil(t, addr)
	assert.Contains(t, err.Error(), "unexpected status code")
}

func TestGetAddress_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{ invalid json"))
	}))
	defer server.Close()

	adapter := &BrasilApiAdapter{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	addr, err := adapter.GetAddress("12345678")

	assert.Error(t, err)
	assert.Nil(t, addr)
	assert.Contains(t, err.Error(), "decode error")
}
