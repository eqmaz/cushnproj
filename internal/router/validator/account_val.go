package validator

import (
	"errors"

	"cushon/internal/router/dto"
	"github.com/google/uuid"
)

func ValidateAccountCreation(input dto.AccountCreationRequest) error {
	if input.ProductUUID == "" {
		return errors.New("product_uuid is required")
	}
	if _, err := uuid.Parse(input.ProductUUID); err != nil {
		return errors.New("product_uuid is not a valid UUID")
	}
	if len(input.FundUUIDs) != 1 {
		return errors.New("exactly one fund entry is required")
	}
	fund := input.FundUUIDs[0]
	if fund.FundUUID == "" {
		return errors.New("fund_uuid is required")
	}
	if _, err := uuid.Parse(fund.FundUUID); err != nil {
		return errors.New("fund_uuid is not a valid UUID")
	}
	if fund.Weight != 100 {
		return errors.New("weight must be exactly 100")
	}
	return nil
}

func ValidateDepositCheck(input dto.DepositCheckRequest) error {
	if input.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if input.AccountId <= 0 {
		// In real life account ID would not be an integer obviously
		// We'd have account numbers, IBANs, Sort codes, UUIDs, etc.
		return errors.New("account_id is required and must be greater than 0")
	}
	return nil
}
