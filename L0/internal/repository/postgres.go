package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDb(config Config) (*sqlx.DB, error) {
	const op = "repository.posgres.New"

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, err
}
