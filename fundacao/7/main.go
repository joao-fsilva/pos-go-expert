package main

import (
	"fmt"
)

/**
maps são constituídos por chave e valor
*/

func main() {
	salarios := map[string]int{"Wesley": 10000, "João": 10000000}

	delete(salarios, "Wesley") //removendo chave

	salarios["Wes"] = 5000
	fmt.Println(salarios)

	//sal := make(map[string]int) //outra forma de criar
	//sal2 := map[string]int{} //outra forma de criar

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d", nome, salario)
	}

	for _, salario := range salarios { //_ (underline) é chamado de blank identifier
		fmt.Printf("O salario é %d", salario)
	}

}
