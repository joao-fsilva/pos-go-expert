package main

import (
	"net/http"
	"strings"

	//"text/template"
	"html/template" //mais seguro para evitar ataques, possui capacidade de adicionar helpers
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t, _ = t.ParseFiles(templates...)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 15},
			{"Java", 10},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
