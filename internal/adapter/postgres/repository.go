package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rafaeljpc/golang-first-study/internal/domain/model"
	"github.com/rafaeljpc/golang-first-study/internal/domain/service"
)

type postgresRepository struct {
	db *sql.DB
}

const query = `
		SELECT id, name, price
		FROM products
	`

// NewPostgresRepository creates a new PostgreSQL-based repository.
func NewPostgresRepository(db *sql.DB) service.Repository {
	return &postgresRepository{db: db}
}

// ListProducts retrieves all products from the database.
func (r *postgresRepository) ListProducts() ([]model.Product, error) {
	var products []model.Product

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning rows: %w", err)
	}

	return products, nil
}
