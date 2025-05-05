package models

import (
	"time"
)

type PurchaseOrder struct {
	ID int
	ItemID int
	Amount float64
	Currency string
	CreatedAt time.Time
	Status string
	CompanyID int
}
