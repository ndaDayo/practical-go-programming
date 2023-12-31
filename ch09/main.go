package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
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

	fetchUser(ctx, "user001")
}

type User struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func fetchAllUser(ctx context.Context) {

	rows, err := db.QueryContext(
		ctx,
		`SELECT user_id, user_name FROM users`,
	)

	if err != nil {
		log.Fatalf("query all users: %v", err)
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var (
			userID, userName string
		)

		if err := rows.Scan(&userID, &userName); err != nil {
			log.Fatalf("scan the user: %v", err)
		}

		users = append(users, &User{
			UserID:   userID,
			UserName: userName,
		})
	}

	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan close: %v", err)
	}

	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatalf("json marshaling failed: %v", err)
	}

	fmt.Println(string(jsonData))
}

func fetchUser(ctx context.Context, userID string) {
	var (
		userName string
	)

	row := db.QueryRowContext(
		ctx,
		`SELECT user_id, user_name FROM users WHERE user_id = $1;`,
		userID,
	)

	err := row.Scan(&userID, &userName)

	if err != nil {
		log.Fatalf("query row(user_id=%s): %v", userID, err)
	}

	u := User{
		UserID:   userID,
		UserName: userName,
	}
	jsonData, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Fatalf("json marshaling failed: %v", err)
	}

	fmt.Println(string(jsonData))
}
