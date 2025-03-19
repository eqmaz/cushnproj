package repository

import (
	"context"

	"cushon/internal/models"
	"github.com/jmoiron/sqlx"
)

// FundRepository defines the interface for fund-related DB operations.
type FundRepository interface {
	GetAvailableFunds(ctx context.Context) ([]models.Fund, error)
	GetFundIdByUuid(ctx context.Context, uuid string) (uint64, error)
}

// fundRepository is a concrete implementation of FundRepository.
type fundRepository struct {
	db *sqlx.DB
}

// NewFundRepository creates a new instance of fundRepository.
func NewFundRepository(db *sqlx.DB) FundRepository {
	return &fundRepository{db: db}
}
