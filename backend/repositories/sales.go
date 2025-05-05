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

// GetAll returns all sales, ordered by date descending
func (r *SaleRepository) GetAll() ([]models.Sale, error) {
    rows, err := r.db.Query(`
        SELECT id, amount, item_id, item_name, currency, date
        FROM sales
        ORDER BY date DESC
    `)
    if err != nil {
        log.Error().Err(err).Msg("Failed to query sales")
        return nil, err
    }
    defer rows.Close()

    var sales []models.Sale
    for rows.Next() {
        var s models.Sale
        if err := rows.Scan(&s.ID, &s.Amount, &s.ItemID, &s.ItemName, &s.Currency, &s.Date); err != nil {
            log.Error().Err(err).Msg("Failed to scan sale row")
            return nil, err
        }
        sales = append(sales, s)
    }
    if err := rows.Err(); err != nil {
        log.Error().Err(err).Msg("Error iterating sale rows")
        return nil, err
    }
    return sales, nil
}

// GetByID returns a single sale by its ID, or nil,nil if not found
func (r *SaleRepository) GetByID(id int) (*models.Sale, error) {
    row := r.db.QueryRow(`
        SELECT id, amount, currency, date, item_name
        FROM sales
        WHERE id = $1
    `, id)

    var s models.Sale
    if err := row.Scan(&s.ID, &s.Amount, &s.Currency, &s.Date); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        log.Error().Err(err).Msg("Failed to scan sale row")
        return nil, err
    }
    return &s, nil
}
