package sqlite

import (
	"database/sql"
	"log"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./pos.db")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db
}
