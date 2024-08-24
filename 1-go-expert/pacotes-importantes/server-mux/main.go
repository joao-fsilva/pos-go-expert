package main

import "net/http"

func main() {
	//multiplexer é mais flexível que o DefaultServeMux, porque você pode criar vários multiplexadores que servem diferentes partes da aplicação
	//maior controle das rotas criadas, porque você tem o controle total sobre o multiplexador
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem-vindo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
