package repository

import (
	"context"

	"cushon/internal/models"
)

// GetAvailableFunds fetches all available funds.
func (r *fundRepository) GetAvailableFunds(ctx context.Context) ([]models.Fund, error) {
	var result []models.Fund
	query := `
		SELECT id, uuid, title, description
		FROM fund		
	`
	err := r.db.SelectContext(ctx, &result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundIdByUuid returns a primary key id from a given UUID, from the fund table
func (r *fundRepository) GetFundIdByUuid(ctx context.Context, uuid string) (uint64, error) {
	var id uint64
	query := `
        SELECT id
        FROM fund
        WHERE uuid = ?
    `
	err := r.db.GetContext(ctx, &id, query, uuid)
	if err != nil {
		return 0, err
	}
	return id, nil
}
