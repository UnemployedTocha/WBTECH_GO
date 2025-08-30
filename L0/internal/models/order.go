package models

import (
	"time"
)

type Order struct {
	OrderUId          string    `json:"order_uid"`
	OrderDelivery     Delivery  `json:"delivery"`
	OrderPayment      Payment   `json:"payment"`
	OrderItems        []Item    `json:"items"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}
