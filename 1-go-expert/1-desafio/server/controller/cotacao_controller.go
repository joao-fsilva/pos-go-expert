package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/server/service"
)

type CotacaoController struct {
	BuscarCotacaoService service.BuscarCotacao
}

func NewCotacaoController(BuscarCotacaoService service.BuscarCotacao) CotacaoController {
	return CotacaoController{
		BuscarCotacaoService: BuscarCotacaoService,
	}
}

func (c CotacaoController) BuscarCotacao(w http.ResponseWriter, r *http.Request) {
	cotacao, err := c.BuscarCotacaoService.Execute()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erro ao processar a requisição"})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cotacao)
}
