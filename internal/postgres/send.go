package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SenderPgx struct {
	db *pgxpool.Pool
}

func NewSenderPgx(db *pgxpool.Pool) *SenderPgx {
	return &SenderPgx{db: db}
}

func (s *SenderPgx) Set(id int, message string) error {
	query := `
	INSERT INTO
	    message(id, time, message)
	VALUES($1, NOW(), $2) `

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), query, id, message)
	if err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}
