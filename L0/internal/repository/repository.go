package repository

import (
	"demo_service/internal/models"
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

// Работа с бд, CRUD

type Repository struct {
	db *sqlx.DB
}

func NewRepository(log *slog.Logger) (*Repository, error) {
	db := NewDb(Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_INTERNAL_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}, log)

	return &Repository{db: db}, nil
}

func (r *Repository) SaveOrder(order models.Order) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("beginning transaction error: %w", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
                customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err = tx.Exec(query, order.OrderUId, order.TrackNumber, order.Entry, order.Locale,
		order.InternalSignature, order.CustomerId, order.DeliveryService,
		order.ShardKey, order.SmId, order.DateCreated, order.OofShard)

	if err != nil {
		return fmt.Errorf("order insert error: %w", err)
	}

	query = `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = tx.Exec(query, order.OrderUId, order.OrderDelivery.Name, order.OrderDelivery.Phone,
		order.OrderDelivery.Zip, order.OrderDelivery.City, order.OrderDelivery.Address,
		order.OrderDelivery.Region, order.OrderDelivery.Email)

	if err != nil {
		return fmt.Errorf("delivery insert error: %w", err)
	}

	query = `INSERT INTO payment (transaction, request_id, currency, provider, amount,
                payment_dt, bank, delivery_cost, goods_total, custom_fee)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = tx.Exec(query, order.OrderUId, order.OrderPayment.RequestId,
		order.OrderPayment.Currency, order.OrderPayment.Provider, order.OrderPayment.Amount,
		order.OrderPayment.PaymentDt, order.OrderPayment.Bank, order.OrderPayment.DeliveryCost,
		order.OrderPayment.GoodsTotal, order.OrderPayment.CustomFee)

	if err != nil {
		return fmt.Errorf("payment insert error: %w", err)
	}

	query = `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size,
                total_price, nm_id, brand, status)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	for _, item := range order.OrderItems {
		_, err = tx.Exec(query, order.OrderUId, item.ChrtId, item.TrackNumber, item.Price,
			item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId,
			item.Brand, item.Status)
		if err != nil {
			return fmt.Errorf("item insert error: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("transaction committing error: %w", err)
	}

	return nil
}

func (r *Repository) GetOrderById(order_uid string) (models.Order, error) {
	// TODO: тоже сделать в транзакции??
	var order models.Order
	query := `SELECT * FROM orders WHERE order_uid = $1`
	err := r.db.Get(&order, query, order_uid)

	if err != nil {
		return order, OrderNotFound
	}

	var delivery models.Delivery
	query = `SELECT * FROM delivery WHERE order_uid = $1`
	err = r.db.Get(&delivery, query, order_uid)

	if err != nil {
		return order, DeliveryNotFound
	}
	order.OrderDelivery = delivery

	var payment models.Payment
	query = `SELECT * FROM payment WHERE order_uid = $1`
	err = r.db.Get(&payment, query, order_uid)

	if err != nil {
		return order, PaymentNotFound
	}
	order.OrderPayment = payment

	var items []models.Item
	query = `SELECT * FROM items WHERE order_uid = $1`
	err = r.db.Select(&items, query, order_uid)

	if err != nil {
		return order, ItemsNotFound
	}
	order.OrderItems = items

	return order, nil
}
