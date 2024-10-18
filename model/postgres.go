package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func StartDB() *sql.DB {
	connStr := "user=fernando dbname=controlepedidos password=admin"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Running on port 5432")
	return db

}
