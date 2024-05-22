/*
* Quando uma variável é declarada, é inferido um valor.
 */

package main

// escopo global
var (
	b bool = true //declarando e atribuindo
	c int
	d string
	e float64
)

func main() {
	//escopo local
	//se não for utilizada, o go não é compilado
	var a string
	f := "oi" //feita a inferência para string. := deve ser feito apenas na primeira vez

	println(a, f)
}
