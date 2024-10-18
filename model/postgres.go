package model

import (
	"database/sql"
	"log"
)

func StartDB() *sql.DB {
	connStr := "user=fernando dbname=controlepedidos sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Running on port 5432")
	return db

}
