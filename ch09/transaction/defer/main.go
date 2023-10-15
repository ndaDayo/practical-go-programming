package main

import (
	"context"
	"database/sql"
)

type Service struct {
	db *sql.DB
}

func (s *Service) UpdateUser(ctx context.Context, userID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.ExecContext(
		ctx,
		"UPDATE users SET user_name = 'nda' WHERE user_id = $1;", userID); err != nil {
		return err
	}

	return tx.Commit()
}
