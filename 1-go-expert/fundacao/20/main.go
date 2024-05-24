package main

func main() {
	a := 1
	b := 2
	c := 2

	if a > b || c > a {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("x")
	}
}
