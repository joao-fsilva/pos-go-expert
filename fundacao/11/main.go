package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Ativo      bool
}

type Cliente struct { //criando uma struct
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Desativar()
}

func (c Cliente) Desativar() { //metodo
	c.Ativo = false
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Ativo = false
	wesley.Logradouro = "test"

	fmt.Println(wesley)
}
