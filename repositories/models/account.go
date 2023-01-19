package models

import "time"

type Account struct {
	Id               string    `json:"id"`
	Available        int64     `json:"available"`
	Blocked          int64     `json:"blocked"`
	Deposited        int64     `json:"deposited"`
	Withdrawn        int64     `json:"with_drawn"`
	Currency         string    `json:"currency"`
	CardName         string    `json:"card_name"`
	CardType         string    `json:"card_type"`
	CardNumber       int       `json:"card_number"`
	CardExpiryMonth  int       `json:"card_expiry_month"`
	CardExpiryYear   int       `json:"card_expiry_year"`
	CardSecurityCode int       `json:"card_security_code"`
	Statement        []string  `json:"statement"`
	CreationTime     time.Time `json:"creation_time"`
}
