package repository

import (
	"context"

	"cushon/internal/models"
)

// ProductRepository defines the interface for product-related DB operations.
type ProductRepository interface {
	// GetAvailableProducts returns a list of products available for the given user.
	GetAvailableProducts(ctx context.Context, userID uint64) ([]models.Product, error)
	GetProductIdByUuid(ctx context.Context, uuid string) (uint64, error)
}
