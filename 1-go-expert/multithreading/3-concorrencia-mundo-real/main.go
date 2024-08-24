package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

//go run -race main.go -- examina o código em busca de condições de corrida

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você é o visitante de número %d", number)))
	})

	http.ListenAndServe(":3000", nil)
}
