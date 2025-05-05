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
        SELECT id, item_id, amount, currency, created_at, status, company_id
        FROM purchase_orders
        ORDER BY created_at DESC
    `)
    if err != nil {
        log.Error().Err(err).Msg("Failed to query purchase_orders")
        return nil, err
    }
    defer rows.Close()

    var orders []models.PurchaseOrder
    for rows.Next() {
        var o models.PurchaseOrder
        if err := rows.Scan(
            &o.ID,
            &o.ItemID,
            &o.Amount,
            &o.Currency,
            &o.CreatedAt,
            &o.Status,
            &o.CompanyID,
        ); err != nil {
            log.Error().Err(err).Msg("Failed to scan purchase_order row")
            return nil, err
        }
        orders = append(orders, o)
    }
    if err := rows.Err(); err != nil {
        log.Error().Err(err).Msg("Error iterating purchase_order rows")
        return nil, err
    }
    return orders, nil
}

func (r *PurchaseOrderRepository) GetByID(id int) (*models.PurchaseOrder, error) {
    row := r.db.QueryRow(`
        SELECT id, item_id, amount, currency, created_at, status, company_id
        FROM purchase_orders
        WHERE id = $1
    `, id)

    var o models.PurchaseOrder
    if err := row.Scan(
        &o.ID,
        &o.ItemID,
        &o.Amount,
        &o.Currency,
        &o.CreatedAt,
        &o.Status,
        &o.CompanyID,
    ); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        log.Error().Err(err).Msg("Failed to scan purchase_order row")
        return nil, err
    }
    return &o, nil
}

func (r *PurchaseOrderRepository) Create(order models.PurchaseOrder) (int, error) {
    var id int
    err := r.db.QueryRow(`
        INSERT INTO purchase_orders
          (item_id, amount, currency, created_at, status, company_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `,
        order.ItemID,
        order.Amount,
        order.Currency,
        order.CreatedAt,
        order.Status,
        order.CompanyID,
    ).Scan(&id)
    if err != nil {
        log.Error().Err(err).Msg("Failed to insert purchase_order")
        return 0, err
    }
    return id, nil
}

func (r *PurchaseOrderRepository) Delete(id int) error {
    _, err := r.db.Exec(`
        DELETE FROM purchase_orders
        WHERE id = $1
    `, id)
    if err != nil {
        log.Error().Err(err).Msg("Failed to delete purchase_order")
    }
    return err
}
