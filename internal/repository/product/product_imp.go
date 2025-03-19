package repository

import (
	"context"

	"cushon/internal/models"
	"github.com/jmoiron/sqlx"
)

// productRepository is a concrete implementation of ProductRepository.
type productRepository struct {
	db *sqlx.DB
}

// NewProductRepository creates a new instance of productRepository.
func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

// GetAvailableProducts fetches products available for a user.
// (For this example, it returns all products of type 'direct'. Adjust the query as needed.)
func (r *productRepository) GetAvailableProducts(ctx context.Context, userID uint64) ([]models.Product, error) {
	var products []models.Product
	query := `
		SELECT id, uuid, title, type, description
		FROM product
		WHERE type = 'direct'
	`
	err := r.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductIdByUuid(ctx context.Context, uuid string) (uint64, error) {
	var id uint64
	query := `
        SELECT id
        FROM product
        WHERE uuid = ?
    `
	err := r.db.GetContext(ctx, &id, query, uuid)
	if err != nil {
		return 0, err
	}
	return id, nil
}
