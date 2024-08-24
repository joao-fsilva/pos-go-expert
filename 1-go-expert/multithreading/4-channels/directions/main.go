package main

func recebe(nome string, hello chan string) {
	hello <- "Olá " + nome
}

func ler(hello chan string) {
	println(<-hello)
}

func main() {
	hello := make(chan string)
	go recebe("Mundo", hello)
	ler(hello)
}
