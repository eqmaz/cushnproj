package api

import (
	"cushon/internal/application/cfg"
	repoAccount "cushon/internal/repository/account"
	repoFund "cushon/internal/repository/fund"
	repoProduct "cushon/internal/repository/product"

	svcAccount "cushon/internal/service/account"
	svcFund "cushon/internal/service/fund"
	svcProduct "cushon/internal/service/product"
	"github.com/jmoiron/sqlx"
)

type RestServer struct {
	Config         cfg.Service
	AccountService svcAccount.AccountServiceInterface
	FundService    svcFund.FundServiceInterface
	ProductService svcProduct.ProductServiceInterface
}

func NewRestServer(db *sqlx.DB, Config cfg.Service) *RestServer {
	if db == nil {
		// This should never happen, but just in case
		panic("Database connection is nil! Exiting...")
	}

	// Set up services and repositories
	fundRepo := repoFund.NewFundRepository(db)
	FundService := svcFund.NewFundService(fundRepo)

	productRepo := repoProduct.NewProductRepository(db)
	ProductService := svcProduct.NewProductService(productRepo)

	accountRepo := repoAccount.NewAccountRepository(db)
	AccountService := svcAccount.NewAccountService(accountRepo, fundRepo, productRepo)

	return &RestServer{
		Config,
		AccountService,
		FundService,
		ProductService,
	}
}
