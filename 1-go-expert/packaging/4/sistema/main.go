package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joao-fsilva/pos-go-expert/1-go-expert/packaging/3/math"
)

func main() {
	math := math.Match{A: 1, B: 2}
	fmt.Println(math.Add(), uuid.New().String())
}
