package main

import (
	"fmt" //importando pacote
)

type ID int

var (
	b bool = true //declarando e atribuindo
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de e é %T", f)
}
