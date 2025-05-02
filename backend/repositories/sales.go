package repositories

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/madn-optimo/backend/models"
)

type SaleRepository struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

func (r *SaleRepository) GetAll(page, pageSize int) ([]models.Sale, error) {
	offset := (page - 1) * pageSize

	rows, err := r.db.Query(`
		SELECT id, amount, currency, date 
		FROM sales
		ORDER BY date DESC
		LIMIT ? OFFSET ?
	`, pageSize, offset)

	if err != nil {
		log.Error().Err(err).Msg("Failed to query sales")
		return nil, err
	}
	defer rows.Close()

	var sales []models.Sale
	for rows.Next() {
		var sale models.Sale
		err := rows.Scan(&sale.ID, &sale.Amount, &sale.Currency, &sale.Date)
		if err != nil {
			log.Error().Err(err).Msg("Failed to scan sale row")
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepository) GetByID(id int) (*models.Sale, error) {
	row := r.db.QueryRow(`
		SELECT id, amount, currency, date
		FROM sales
		WHERE id = ?
	`, id)

	var sale models.Sale
	err := row.Scan(&sale.ID, &sale.Amount, &sale.Currency, &sale.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Err(err).Msg("Failed to scan sale row")
		return nil, err
	}

	return &sale, nil
}
