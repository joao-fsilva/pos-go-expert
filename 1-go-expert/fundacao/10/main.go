package main

import (
	"fmt"
)

func main() {
	total := func() int { //closure
		return sum(1, 2) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int { //funções variádicas
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
