package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	OrderNotFound    = errors.New("order not found")
	DeliveryNotFound = errors.New("delivery not found")
	ItemsNotFound    = errors.New("items not found")
	PaymentNotFound  = errors.New("payment not found")
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDb(config Config, log *slog.Logger) *sqlx.DB {

	for {
		db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

		if err != nil {
			log.Error("connection to db failed | retrying in 3 sec")
			time.Sleep(3 * time.Second)
			continue
		}
		return db
	}
}
