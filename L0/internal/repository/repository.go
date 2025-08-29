package repository

import "github.com/jmoiron/sqlx"

// Работа с бд, CRUD

type Repository struct {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
