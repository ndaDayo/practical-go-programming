package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	db *sql.DB
}

func main() {
	dbConn, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s := &Service{db: dbConn}
	ctx := context.Background()

	if err := s.UpdateUser(ctx, "user001"); err != nil {
		log.Fatalf("Update User failed: %v", err)
	}

}

func (s *Service) UpdateUser(ctx context.Context, userID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(
		ctx,
		"UPDATE users SET user_name = 'nda' WHERE user_id = $1;", userID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
