package db

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/madn-optimo/backend/models"
)

func returnRandomDate() string {
	year := 2025
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func fillSales(db *sql.DB) error {
	for i := 0; i < 100; i++ {
	_, err := db.Exec(`
	INSERT INTO sales (amount, currency, date) VALUES (?, ?, ?)
	`, rand.Float64() * 1000, "USD", returnRandomDate())
		if err != nil {
			return err
		}
	}

	return nil
}

func groupSalesByMonth(db *sql.DB) ([]models.SaleGroupedByMonth, error) {
	rows, err := db.Query(`
        SELECT DATE_TRUNC('month', sale_date) AS month, SUM(amount) AS total_sales
        FROM sales
        GROUP BY month
        ORDER BY month
    `)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	byMonth := make(map[string][]models.Sale)

	for rows.Next() {
		var sale models.Sale
		err = rows.Scan(&sale.Date, &sale.Amount, &sale.Currency)
		if err != nil {
			return nil, err
		}

		byMonth[sale.Date.Month().String()] = append(byMonth[sale.Date.Month().String()], sale)
	}

	results := make([]models.SaleGroupedByMonth, 0, len(byMonth))
	for month, sales := range byMonth {
		results = append(results, models.SaleGroupedByMonth{
			Month: month,
			Sales: sales,
		})
	}

	return results, nil
}
