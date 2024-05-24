package main

type MyNumber int

type Number interface { //constraint que pode ser utilizada
	~int | float64 // o ~ considera que pode aceitar o MyNumber, pois é um int também... sem isso não funciona
}

func Soma[T Number](m map[string]T) T { //generics
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool { //comparable é uma constraint que permite que apenas valores que são comparáveis (do mesmo tipo)
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
