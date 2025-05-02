package models

type Item struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Quantity  int     `json:"quantity"`
	CompanyID int     `json:"companyId"`
} 