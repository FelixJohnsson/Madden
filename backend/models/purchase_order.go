package models

import (
	"time"
)

type PurchaseOrder struct {
	ID        int       `json:"id"`
	Item      Item      `json:"item"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`
	CompanyID int       `json:"companyId"`
} 