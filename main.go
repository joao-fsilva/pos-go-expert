package main

import "fmt"

func main() {
	event := []string{"teste", "teste2", "teste3"}

	//event = append(event[:1], event[1:]...)

	fmt.Println(event[:1])
}
