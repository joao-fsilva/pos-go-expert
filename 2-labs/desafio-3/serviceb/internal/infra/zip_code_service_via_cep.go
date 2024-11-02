package infra

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"weather-zip-code/internal/entity"
)

var ErrRequestZipCode = errors.New("an error occurred while processing your request of zipcode")

type ViaCEP struct {
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

type ZipCodeServiceViaCep struct {
	httpTracer *HttpTracer
}

func NewZipCodeServiceViaCep(tracer *HttpTracer) *ZipCodeServiceViaCep {
	return &ZipCodeServiceViaCep{
		httpTracer: tracer,
	}
}

func (z *ZipCodeServiceViaCep) GetCityByZipCode(ctx context.Context, zipCode *entity.ZipCode) (string, error) {
	log.Printf("Info: checking the city of the zipcode %s", zipCode.ZipCode)

	res, err := z.httpTracer.Get(ctx, "http://viacep.com.br/ws/"+zipCode.ZipCode+"/json/", "chamada viacep")
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: unexpected status code %d from ViaCEP API response", res.StatusCode)
		return "", ErrRequestZipCode
	}

	if err != nil {
		log.Printf("Error: %v", err)
		return "", ErrRequestZipCode
	}

	body, err := io.ReadAll(res.Body)

	var data ViaCEP
	err = json.Unmarshal(body, &data)

	log.Printf("Info: the city of the zipcode %s is %s", zipCode.ZipCode, data.Localidade)

	return data.Localidade, err
}
