package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Sender interface {
	Set(id int, message string) error
}

type Repository struct {
	Sender
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Sender: NewSenderPgx(db)}
}
