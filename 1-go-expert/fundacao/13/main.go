package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	a := 10
	b := 20
	soma(&a, &b)
	println(a)
}
