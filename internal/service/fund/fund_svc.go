package service

import (
	"context"

	"cushon/internal/models"
	fr "cushon/internal/repository/fund"
)

type FundServiceInterface interface {
	GetAvailableFunds(ctx context.Context) ([]models.Fund, error)
}

type fundService struct {
	repo fr.FundRepository
}

// NewFundService creates a new fundService.
func NewFundService(repo fr.FundRepository) FundServiceInterface {
	return &fundService{repo: repo}
}

// GetAvailableFunds delegates the call to the repository.
func (s *fundService) GetAvailableFunds(ctx context.Context) ([]models.Fund, error) {
	return s.repo.GetAvailableFunds(ctx)
}
