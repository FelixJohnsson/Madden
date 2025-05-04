package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func InitDB() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "madden.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info().Msg("Database connection established")

	if err = createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	if err := createCompanies(db); err != nil {
		return err
	}

	if err := createItems(db); err != nil {
		return err
	}

	if err := createSales(db); err != nil {
		return err
	}

	if err := createPurchaseOrders(db); err != nil {
		return err
	}

	log.Info().Msg("Database tables created")
	return nil
} 

func createCompanies(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS companies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		address TEXT NOT NULL
	)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	INSERT INTO companies (name, address) VALUES ('Our Legacy', 'Odengatan 53'), ('Nudie Jeans', 'Nybrogatan 12')
	`)
	if err != nil {
		return err
	}

	return nil
}

func createItems(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		currency TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		company_id INTEGER,
		FOREIGN KEY (company_id) REFERENCES companies (id)
	)`)
	if err != nil {
		return err
	}

	return nil
}



func createSales(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		currency TEXT NOT NULL,
		date TIMESTAMP NOT NULL
	)`)
	if err != nil {
		return err
	}

	err = fillSales(db)
	if err != nil {
		return err
	}

	return nil
}

func createPurchaseOrders(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS purchase_orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id INTEGER NOT NULL,
		amount REAL NOT NULL,
		currency TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		status TEXT NOT NULL,
		company_id INTEGER,
		FOREIGN KEY (item_id) REFERENCES items (id),
		FOREIGN KEY (company_id) REFERENCES companies (id)
	)`)
	if err != nil {
		return err
	}

	return nil
}

