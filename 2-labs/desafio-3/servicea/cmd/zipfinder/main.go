package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func validateZipCode(cep string) bool {
	var cepRegex = regexp.MustCompile(`^\d{8}$`)
	return cepRegex.MatchString(cep)
}

type Request struct {
	CEP string `json:"cep"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if !validateZipCode(req.CEP) {
			http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		url := os.Getenv("SERVICEB_URL") + "weather?zipcode=" + req.CEP

		res, err := http.Get(url)
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		if err != nil {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		body, err := io.ReadAll(res.Body)

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8081", nil)
}
