package db

import (
	"database/sql"
	"time"

	"github.com/madn-optimo/backend/models"
)

func groupSalesByMonth(db *sql.DB) ([]models.SaleGroupedByMonth, error) {
	rows, err := db.Query(`
        SELECT TO_CHAR(date, 'YYYY-MM') AS month, SUM(amount) AS total_amount, currency
        FROM sales
        GROUP BY month, currency
        ORDER BY month
    `)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	byMonth := make(map[string][]models.Sale)

	for rows.Next() {
		var monthStr string
		var sale models.Sale
		err = rows.Scan(&monthStr, &sale.Amount, &sale.Currency)
		if err != nil {
			return nil, err
		}

		// Parse the month string to get the month name
		t, err := time.Parse("2006-01", monthStr)
		if err != nil {
			return nil, err
		}
		monthName := t.Month().String()

		sale.Date = t
		byMonth[monthName] = append(byMonth[monthName], sale)
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
