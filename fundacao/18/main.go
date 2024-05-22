package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joao-fsilva/pos-go-expert/matematica" //utilizando pacote
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", soma)
	//baixar um pacote
	// go get github.com/google/uuid

	fmt.Println(uuid.New())
}
