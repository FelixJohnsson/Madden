package repositories

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/madn-optimo/backend/models"
)

type PurchaseOrderRepository struct {
	db *sql.DB
}

func NewPurchaseOrderRepository(db *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

func (r *PurchaseOrderRepository) GetAll() ([]models.PurchaseOrder, error) {
	rows, err := r.db.Query(`
		SELECT po.id, po.amount, po.currency, po.created_at, po.status, po.company_id,
		       i.id, i.name, i.price, i.currency, i.quantity, i.company_id
		FROM purchase_orders po
		JOIN items i ON po.item_id = i.id
		ORDER BY po.created_at DESC
	`)

	if err != nil {
		log.Error().Err(err).Msg("Failed to query purchase orders")
		return nil, err
	}
	defer rows.Close()

	var orders []models.PurchaseOrder
	for rows.Next() {
		var order models.PurchaseOrder
		var item models.Item
		err := rows.Scan(
			&order.ID, &order.Amount, &order.Currency, &order.CreatedAt, &order.Status, &order.CompanyID,
			&item.ID, &item.Name, &item.Price, &item.Currency, &item.Quantity, &item.CompanyID,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to scan purchase order row")
			return nil, err
		}
		order.Item = item
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *PurchaseOrderRepository) GetByID(id int) (*models.PurchaseOrder, error) {
	row := r.db.QueryRow(`
		SELECT po.id, po.amount, po.currency, po.created_at, po.status, po.company_id,
		       i.id, i.name, i.price, i.currency, i.quantity, i.company_id
		FROM purchase_orders po
		JOIN items i ON po.item_id = i.id
		WHERE po.id = ?
	`, id)

	var order models.PurchaseOrder
	var item models.Item
	err := row.Scan(
		&order.ID, &order.Amount, &order.Currency, &order.CreatedAt, &order.Status, &order.CompanyID,
		&item.ID, &item.Name, &item.Price, &item.Currency, &item.Quantity, &item.CompanyID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Err(err).Msg("Failed to scan purchase order row")
		return nil, err
	}
	order.Item = item

	return &order, nil
}

func (r *PurchaseOrderRepository) Create(order models.PurchaseOrder) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("Failed to begin transaction")
		return 0, err
	}
	defer tx.Rollback()

	// First, insert or get the item
	var itemID int
	if order.Item.ID > 0 {
		itemID = order.Item.ID
	} else {
		result, err := tx.Exec(`
			INSERT INTO items (name, price, currency, quantity, company_id)
			VALUES (?, ?, ?, ?, ?)
		`, order.Item.Name, order.Item.Price, order.Item.Currency, order.Item.Quantity, order.Item.CompanyID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to insert item")
			return 0, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Error().Err(err).Msg("Failed to get last insert ID for item")
			return 0, err
		}
		itemID = int(id)
	}

	// Insert the purchase order
	result, err := tx.Exec(`
		INSERT INTO purchase_orders (item_id, amount, currency, created_at, status, company_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`, itemID, order.Amount, order.Currency, order.CreatedAt, order.Status, order.CompanyID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert purchase order")
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get last insert ID for purchase order")
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction")
		return 0, err
	}

	return int(id), nil
}

func (r *PurchaseOrderRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM purchase_orders WHERE id = ?", id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete purchase order")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Int("id", id).Msg("No purchase order found to delete")
	}

	return nil
} 