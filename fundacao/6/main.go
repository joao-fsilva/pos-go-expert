package main

import (
	"fmt"
)

/**
um slice é constituído por arrays, não possui um tamanho infinito.
Quando a capacidade máxima é antigida, o go cria array por debaixo dos panos com dobro do tamanho do atual
o : é um ponto de corte, onde criamos o slice de um slice
*/

func main() {
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
