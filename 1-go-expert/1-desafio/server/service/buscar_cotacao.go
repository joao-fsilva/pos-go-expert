package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BuscarCotacao struct {
	Db *sql.DB
}

func NewBuscarCotacao(db *sql.DB) BuscarCotacao {
	return BuscarCotacao{
		Db: db,
	}
}

type jsonData struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Cotacao struct {
	Cotacao string `json:"cotacao"`
}

func (b BuscarCotacao) Execute() (Cotacao, error) {
	cotacao, err := b.buscarCotacao()
	if err != nil {
		return cotacao, err
	}

	err = b.salvarCotacao(cotacao)
	if err != nil {
		return cotacao, err
	}

	return cotacao, nil
}

func (b BuscarCotacao) buscarCotacao() (Cotacao, error) {
	cotacao := Cotacao{}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return cotacao, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return cotacao, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cotacao, err
	}

	var data jsonData
	if err := json.Unmarshal(body, &data); err != nil {
		return cotacao, err
	}

	cotacao.Cotacao = data.USDBRL.Bid

	return cotacao, err
}

func (b BuscarCotacao) salvarCotacao(cotacao Cotacao) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err := b.Db.ExecContext(ctx, "INSERT INTO cotacoes (cotacao) VALUES (?)", cotacao.Cotacao)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("falha ao inserir cotacao: %w (context deadline exceeded)", err)
		}
		return fmt.Errorf("falha ao inserir cotacao: %w", err)
	}

	return nil
}
