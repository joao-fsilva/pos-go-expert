package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3))
}

func sum(numeros ...int) int { //funções variádicas
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
