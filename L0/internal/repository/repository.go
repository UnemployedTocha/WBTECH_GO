package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// Работа с бд, CRUD

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	db, err := NewDb(Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres: %w: ", err.Error())
	}

	return &Repository{db: db}, nil
}

func (r *Repository) SaveOrder() (orderUid string, err error) {
	return
}

func (r *Repository) GetOrderById(order_uid string) {

}
