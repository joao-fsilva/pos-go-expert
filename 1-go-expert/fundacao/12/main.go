package main

func main() {
	a := 10
	var ponteiro *int = &a // o ponteiro é o endereço na memória
	*ponteiro = 20

	b := &a
	println(*b) //desreferenciando
}
