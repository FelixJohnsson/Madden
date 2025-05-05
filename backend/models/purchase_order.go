package models

import (
	"time"
)

type PurchaseOrder struct {
	ID        int       `json:"id"`
	ItemID    int       `json:"itemId"`
	Amount    int       `json:"amount"`
	ItemName  string    `json:"itemName"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`
	CompanyID int       `json:"companyId"`
}
