package main

import "fmt"

func main() {
	teste()
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}

func teste() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Quarta linha")
}
