package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Willians"

	println(minhaVar.(string)) //type assersation

	res, ok := minhaVar.(int)
	fmt.Printf("o valor de res é %v e o resultado de ok é %v", res, ok)
}
