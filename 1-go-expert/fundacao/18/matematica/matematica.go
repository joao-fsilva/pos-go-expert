package matematica

func Soma[T int | float64](a, b T) T { //generics
	return a + b
}
