package main

import (
	"fmt"
	"github.com/joao-fsilva/pos-go-expert/1-go-expert/packaging/1/math"
)

func main() {
	math := math.Match{A: 1, B: 2}
	fmt.Println(math.Add())
}
