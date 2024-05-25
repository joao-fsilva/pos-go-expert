package main

import (
	"net/http"
	"server/server/controller"
	"server/server/db"
	"server/server/service"
)

func main() {
	dbConn := (db.NewDb()).Conectar()

	buscarCotacaoService := service.NewBuscarCotacao(dbConn)
	cotacaoController := controller.NewCotacaoController(buscarCotacaoService)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/cotacao", cotacaoController.BuscarCotacao)
	http.ListenAndServe(":8080", nil)
}
