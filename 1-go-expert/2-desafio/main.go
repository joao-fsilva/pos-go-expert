package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

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

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	ch := make(chan map[string]interface{})

	cep := "01310930"

	go BuscaCepViaCEP(cep, ch)
	go BuscaCepBrasilApiCep(cep, ch)

	for {
		select {
		case result := <-ch:
			if result["error"] != nil {
				fmt.Println(result["error"])
				return
			}

			fmt.Println(result["api"])
			fmt.Println(result["d	ata"])
			return

		case <-time.After(time.Second * 1):
			fmt.Println("Erro de timeout")
			return
		}
	}

}

func BuscaCepViaCEP(cep string, ch chan map[string]interface{}) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	var data ViaCEP
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	ch <- map[string]interface{}{"data": data, "api": "viacep"}
}

func BuscaCepBrasilApiCep(cep string, ch chan map[string]interface{}) {
	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	var data BrasilApiCep
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- map[string]interface{}{"error": err}
	}

	ch <- map[string]interface{}{"data": data, "api": "brasilapi"}
}
