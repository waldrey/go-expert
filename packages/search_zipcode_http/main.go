package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type AddressResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", searchZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func searchZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipcodeParam := r.URL.Query().Get("zipcode")
	if zipcodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	zipcode, error := SearchZipCode(zipcodeParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipcode)
}

func SearchZipCode(zipcode string) (*AddressResponse, error) {
	response, error := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")
	if error != nil {
		return nil, error
	}

	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var z AddressResponse
	error = json.Unmarshal(body, &z)
	if error != nil {
		return nil, error
	}

	return &z, nil
}
