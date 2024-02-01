package lib

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func conn() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=go-crud password=1 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

var DB *sqlx.DB = conn()
