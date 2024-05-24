package main

import (
	"fmt"
)

func main() {
	var meuArray [3]int //criando um array
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}

}
