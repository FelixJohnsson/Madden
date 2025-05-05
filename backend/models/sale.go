package models

import (
	"time"
)

type Sale struct {
	ID       int       `json:"id"`
	ItemID   int       `json:"itemId"`
	ItemName string    `json:"itemName"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Date     time.Time `json:"date"`
} 
