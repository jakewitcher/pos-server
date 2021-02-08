package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Datastore struct {
	store *sql.DB
}

func NewDatastore() *Datastore {
	db, err := sql.Open("sqlite3", "./pointOfSale.db")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &Datastore{store: db}
}
