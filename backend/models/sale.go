package models

import (
	"time"
)

type Sale struct {
	ID       int       `json:"id"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Date     time.Time `json:"date"`
} 

type SaleGroupedByMonth struct {
	Month string `json:"month"`
	Sales []Sale `json:"sales"`
}
