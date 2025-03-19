package service

import (
	"context"

	"cushon/internal/application/cfg"
	ar "cushon/internal/repository/account"
	fr "cushon/internal/repository/fund"
	pr "cushon/internal/repository/product"
	"cushon/internal/router/dto"
	c "cushon/pkg/console"
	"cushon/pkg/e"
)

type AccountServiceInterface interface {
	GetAccountBalances(ctx context.Context, userId uint64) ([]dto.AccountBalancesResponse, error)
	OpenRetailAccount(background context.Context, userId uint64, input dto.AccountCreationRequest) (*uint64, error)
	CheckDepositAmount(background context.Context, userId uint64, accountId uint64, amount float64, cfg cfg.Service) (bool, error)
}

type accountService struct {
	repo        ar.AccountRepository
	fundRepo    fr.FundRepository
	productRepo pr.ProductRepository
}

// NewAccountService creates a new account service.
func NewAccountService(ar ar.AccountRepository, fr fr.FundRepository, pr pr.ProductRepository) AccountServiceInterface {
	return &accountService{repo: ar, fundRepo: fr, productRepo: pr}
}

// GetAccountBalances delegates the call to the repository.
func (s *accountService) GetAccountBalances(ctx context.Context, userId uint64) ([]dto.AccountBalancesResponse, error) {
	return s.repo.GetAccountBalances(ctx, userId)
}

// OpenRetailAccount - creates a new retail account record for the user, if they're allowed to
func (s *accountService) OpenRetailAccount(ctx context.Context, userId uint64, input dto.AccountCreationRequest) (*uint64, error) {

	// Check the user doesn't have any other retail accounts
	count, err := s.repo.CountRetailAccounts(ctx, userId)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, e.FromCode("eAcEx1")
	}

	// Check if the product uuid is a real product
	givenProductUuid := input.ProductUUID
	productId, err := s.productRepo.GetProductIdByUuid(ctx, givenProductUuid)
	if err != nil {
		return nil, err
	}
	if productId == 0 {
		return nil, e.FromCode("ePuNf1")
	}

	// Check the fund uuid is a real fund
	givenFundUuid := input.FundUUIDs[0].FundUUID
	fundId, err := s.fundRepo.GetFundIdByUuid(ctx, givenFundUuid)
	if err != nil {
		return nil, err
	}
	if fundId == 0 {
		return nil, e.FromCode("eFuNf1")
	}

	// From here we are sure that our product and fund are real, we can create the account
	return s.repo.CreateRetailAccount(ctx, userId, productId, fundId)
}

// CheckDepositAmount - checks if the user can deposit to the given account by the intended amount
func (s *accountService) CheckDepositAmount(ctx context.Context, userId uint64, accountId uint64, amount float64, cfg cfg.Service) (bool, error) {
	// Check if the account belongs to the user
	userOwnsAccount, err := s.repo.UserOwnsAccount(ctx, userId, accountId)
	if err != nil {
		return false, err
	}
	if !userOwnsAccount {
		// In real life we should never need to do this check
		// because a) we'd have a proper auth service and
		// b) by design it would be impossible to get into to this situation
		// This is just a demonstration of understanding the concept
		return false, e.FromCode("eAcEx2")
	}

	// Get this account's remaining deposit allowance
	// Discover how much has been deposited since the last 6th April
	m, d, err := parseMonthDay(cfg.IsaAllowanceResetDate)
	if err != nil {
		return false, err
	}

	taxYearStartDate := GetLastTaxYearStartDate(m, d)
	totalDeposits, err := s.repo.GetTotalDepositsSince(ctx, accountId, taxYearStartDate)
	if err != nil {
		return false, err
	}

	c.Infof("Total deposits since %s: %f", taxYearStartDate, totalDeposits)

	// Get the deposit allowance for this account
	headroom := float64(cfg.MaxIsaAnnualInvestment) - (totalDeposits + amount)
	if headroom < 0 {
		c.Warnf("ISA deposit headroom is negative: %f", headroom)
		return false, e.FromCode("eIafId")
	}

	return true, nil
}
