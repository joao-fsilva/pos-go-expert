package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"-"` // "-" significa que não vai ser serializado
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{1, 100}
	res, err := json.Marshal(conta) //aqui eu guardo o valor
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) //aqui eu mando direto para o output
	if err != nil {
		panic(err)
	}

	jsonPuro := `{"n":1,"s":200}`

	var contaX Conta
	err = json.Unmarshal([]byte(jsonPuro), &contaX)
	if err != nil {
		panic(err)
	}

	println(contaX.Numero)
	println(contaX.Saldo)
}
