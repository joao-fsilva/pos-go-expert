package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Db struct {
	Db *sql.DB
}

func NewDb() *Db {
	return &Db{}
}

func (d *Db) Conectar() *sql.DB {
	db, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cotacoes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        cotacao INTEGER NOT NULL
    )`)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
