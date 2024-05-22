package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois"}
	for _, v := range numeros {
		print(v)
	}

	i := 0
	for i < 10 {
		print(i)
		i++
	}

	for {
		print("Hello")
	}
}
