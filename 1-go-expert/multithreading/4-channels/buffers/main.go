package main

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "world"

	println(<-ch)
	println(<-ch)
}
