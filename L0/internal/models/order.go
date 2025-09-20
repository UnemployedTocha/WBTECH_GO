package models

import (
	"time"
)

type Order struct {
	OrderUId          string    `json:"order_uid" db:"order_uid"`
	OrderDelivery     Delivery  `json:"delivery" db:"-"`
	OrderPayment      Payment   `json:"payment" db:"-"`
	OrderItems        []Item    `json:"items" db:"-"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shardkey"`
	SmId              int64     `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}
