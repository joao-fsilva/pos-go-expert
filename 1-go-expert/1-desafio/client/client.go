package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type cotacao struct {
	Cotacao string `json:"cotacao"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var cotacao cotacao

	if err := json.Unmarshal(body, &cotacao); err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("Dólar: %s\n", cotacao.Cotacao)

	f.Write([]byte(output))
	fmt.Print(output)
}
