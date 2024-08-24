package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)

	//for _, cep := range os.Args[1:] {
	//	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	//	req, err := http.Get(url)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", url, err)
	//	}
	//
	//	defer req.Body.Close()
	//
	//	res, err := io.ReadAll(req.Body)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", url, err)
	//	}
	//
	//	var data ViaCEP
	//	err = json.Unmarshal(res, &data)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error unmarshalling %s: %v\n", url, err)
	//	}
	//
	//	file, err := os.Create("cidade.txt")
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error creating file cidade.txt: %v\n", err)
	//	}
	//
	//	defer file.Close()
	//
	//	_, err = file.WriteString(fmt.Sprintf("CEP: %s\n, Logradouro: %s\n, Complemento: %s\n, Bairro: %s\n, Localidade: %s\n, UF: %s\n, IBGE: %s\n, GIA: %s\n, DDD: %s\n, SIAFI: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi))
	//
	//	fmt.Println(data)
	//}
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data ViaCEP
	err = json.Unmarshal(body, &data)

	return &data, err
}
