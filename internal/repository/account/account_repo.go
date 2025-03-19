package repository

import (
	"context"
	"time"

	"cushon/internal/router/dto"
	"github.com/jmoiron/sqlx"
)

// AccountRepository defines the interface for account-related DB operations.
type AccountRepository interface {
	GetAccountBalances(ctx context.Context, userId uint64) ([]dto.AccountBalancesResponse, error)
	CountRetailAccounts(ctx context.Context, userId uint64) (int, error)
	CreateRetailAccount(ctx context.Context, userId uint64, productId uint64, fundId uint64) (*uint64, error)
	GetTotalDepositsSince(ctx context.Context, accountId uint64, date time.Time) (float64, error)
	UserOwnsAccount(ctx context.Context, userId uint64, accountId uint64) (bool, error)
}

// AccountRepository is a concrete implementation of AccountRepository.
type accountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository creates a new instance of AccountRepository.
func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{db: db}
}
