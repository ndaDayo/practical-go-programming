package main

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func main() {
	dbConn, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	db = dbConn

	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fetchAllUser(ctx)
}
