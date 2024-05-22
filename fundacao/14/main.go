package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := NewConta()
	conta.simular(200)
	println(conta.saldo)
}
