package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-recs-backend/internal/repository/recs"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (recs.Repository, error) {
	return &Repository{db: db}, nil
}
