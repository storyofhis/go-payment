package models

import "time"

type Payment struct {
	Id              string    `json:"id"`
	BusinessId      string    `json:"business_id"`
	OrderId         string    `json:"order_id"`
	Operation       string    `json:"operation"`
	OriginalAmount  int64     `json:"original_amount"`
	CurrentAmount   int64     `json:"current_amount"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	Currency        string    `json:"currency"`
	CardName        string    `json:"card_name"`
	CardType        string    `json:"card_type"`
	CardNumber      int64     `json:"card_number"`
	CardExpiryMonth int       `json:"card_expiry_month"`
	CardExpiryYear  int       `json:"card_expiry_year"`
	CreationTime    time.Time `json:"creation_time"`
}
