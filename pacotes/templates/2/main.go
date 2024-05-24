package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	//o must é um helper que faz o parse e já retorna o template pronto em um único passo
	tmp := template.Must(template.New("cursoTemplate").Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
