package service

import (
	"context"

	"cushon/internal/models"
	pr "cushon/internal/repository/product"
)

// ProductServiceInterface defines business operations related to products.
type ProductServiceInterface interface {
	GetAvailableProducts(ctx context.Context, userID uint64) ([]models.Product, error)
}

// productService is a concrete implementation of ProductServiceInterface.
type productService struct {
	repo pr.ProductRepository
}

// NewProductService creates a new ProductService.
func NewProductService(repo pr.ProductRepository) ProductServiceInterface {
	return &productService{repo: repo}
}

// GetAvailableProducts delegates the call to the repository.
func (s *productService) GetAvailableProducts(ctx context.Context, userID uint64) ([]models.Product, error) {
	return s.repo.GetAvailableProducts(ctx, userID)
}
