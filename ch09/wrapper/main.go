package main

import (
	"context"
	"database/sql"
	"fmt"
)

type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}

	return tx.Commit()
}

func (s *Service) UpdateUser(ctx context.Context, userID string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(
			ctx,
			"UPDATE users SET user_name = 'nda' WHERE user_id = $1;", userID); err != nil {
			return err
		}
		return nil
	}

	return s.tx.Transaction(ctx, updateFunc)
}
