package repository

import (
	"context"
	"time"

	"cushon/internal/router/dto"
)

// GetAccountBalances - Get the balances of all accounts for a user
func (a *accountRepository) GetAccountBalances(ctx context.Context, userId uint64) ([]dto.AccountBalancesResponse, error) {
	query := `
        SELECT a.product_id, p.title AS product_title, a.total_investment, a.current_balance
        FROM user_account a
        JOIN product p ON a.product_id = p.id
        WHERE a.user_id = ?
    `
	var result []dto.AccountBalancesResponse
	err := a.db.SelectContext(ctx, &result, query, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CountRetailAccounts - count the number of "direct" (retail) accounts for a user
func (a *accountRepository) CountRetailAccounts(ctx context.Context, userId uint64) (int, error) {
	query := `
        SELECT COUNT(*)
        FROM user_account ua
        INNER JOIN product p ON ua.product_id = p.id
        WHERE ua.user_id = ? AND p.type = 'direct'
    `
	var count int
	err := a.db.GetContext(ctx, &count, query, userId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CreateRetailAccount - Create a new retail account safely
func (a *accountRepository) CreateRetailAccount(ctx context.Context, userId uint64, productId uint64, fundId uint64) (*uint64, error) {
	tx, err := a.db.BeginTxx(ctx, nil) // Start a transaction
	if err != nil {
		return nil, err
	}

	// Insert into user_account
	query := `
        INSERT INTO user_account (user_id, product_id, total_investment, current_balance)
        VALUES (?, ?, 0, 0)
    `
	qResult, err := tx.ExecContext(ctx, query, userId, productId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get the new account ID that was just inserted
	newAccountId, err := qResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Insert into user_account_funds
	query = `
        INSERT INTO user_account_funds (fund_id, user_account_id, weight_pc)
        VALUES (?, ?, 100) # TODO - Weight is always 100% for this example
    `
	_, err = tx.ExecContext(ctx, query, fundId, newAccountId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	result := uint64(newAccountId)

	return &result, nil
}

// GetTotalDepositsSince - Get the total deposits made to an account since a certain date
// This includes a sum of positive amounts only
func (a *accountRepository) GetTotalDepositsSince(ctx context.Context, accountId uint64, date time.Time) (float64, error) {
	query := `
        SELECT COALESCE(SUM(amount), 0) # Ensure we always return a value
        FROM user_account_transaction
        WHERE account_id = ? AND created >= ? AND amount > 0
    `
	var total float64
	err := a.db.GetContext(ctx, &total, query, accountId, date)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// UserOwnsAccount - returns true if the user owns the account. Error is only if there is a DB issue
func (a *accountRepository) UserOwnsAccount(ctx context.Context, userId uint64, accountId uint64) (bool, error) {
	query := `
        SELECT COUNT(*)
        FROM user_account
        WHERE user_id = ? AND id = ?
    `
	var count int
	err := a.db.GetContext(ctx, &count, query, userId, accountId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
