package api

import (
	"context"

	"cushon/internal/router/converter"
	"cushon/internal/router/dto"
	"cushon/internal/router/validator"
	"github.com/gofiber/fiber/v2"
)

func (s *RestServer) GetUserProductAvailable(ctx *fiber.Ctx) error {
	// Placeholder userID - it would be coming from the JWT or similar auth
	// get "userId" header from ctx
	userId, err := getUserId(ctx)
	if err != nil {
		getCtxLogger(ctx).Errorf("Error parsing userId: %v", err)
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{"error": "Invalid userId"})
	}

	// Convert Fiber context to a standard context.
	products, err := s.ProductService.GetAvailableProducts(context.Background(), userId)
	//products, err := s.ProductService.GetAvailableProducts(ctx, userID)
	if err != nil {
		// Get the logger from the Fiber context
		log := getCtxLogger(ctx)
		log.Errorf("Error fetching available products: %v", err)
		return replyError(ctx, fiber.StatusInternalServerError, fiber.Map{
			"error": "Unable to fetch products",
		})
	}

	result := converter.MapProductsToResponse(products)

	return replyResult(ctx, result)
}

func (s *RestServer) GetIsaFundList(ctx *fiber.Ctx) error {
	// Placeholder userID - it would be coming from the JWT or similar auth
	// userID := int64(1)

	// Convert Fiber context to a standard context.
	funds, err := s.FundService.GetAvailableFunds(context.Background())
	if err != nil {
		// Get the logger from the Fiber context
		log := getCtxLogger(ctx)
		log.Errorf("Error fetching ISA funds: %v", err)
		return replyError(ctx, fiber.StatusInternalServerError, fiber.Map{
			"error": "Unable to fetch ISA funds",
		})
	}

	result := converter.MapFundsToResponse(funds)

	return replyResult(ctx, result)
}

func (s *RestServer) GetUserAccountBalances(ctx *fiber.Ctx) error {
	// Placeholder userID - it would be coming from the JWT or similar auth
	// get "userId" header from ctx
	userId, err := getUserId(ctx)
	if err != nil {
		getCtxLogger(ctx).Errorf("Error parsing userId: %v", err)
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{"error": "Invalid userId"})
	}

	// Convert Fiber context to a standard context.
	balances, err := s.AccountService.GetAccountBalances(context.Background(), userId)
	if err != nil {
		// Get the logger from the Fiber context
		log := getCtxLogger(ctx)
		log.Errorf("Error fetching account balances: %v", err)
		return replyError(ctx, fiber.StatusInternalServerError, fiber.Map{
			"error": "Unable to fetch account balances " + err.Error(),
		})
	}

	//result := converter.MapBalancesToResponse(balances)

	return replyResult(ctx, balances)
}

func (s *RestServer) OpenRetailAccount(ctx *fiber.Ctx) error {
	// Placeholder userID - it would be coming from the JWT or similar auth
	// get "userId" header from ctx
	userId, err := getUserId(ctx)
	if err != nil {
		getCtxLogger(ctx).Errorf("Error parsing userId: %v", err)
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{"error": "Invalid userId"})
	}

	// Read the body of the request into a map
	var input dto.AccountCreationRequest
	if err := ctx.BodyParser(&input); err != nil {
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{
			"error": "Invalid request body",
		})
	}

	// We will need to validate the input, so we use a validator for that
	err = validator.ValidateAccountCreation(input)
	if err != nil {
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	newAccountId, err := s.AccountService.OpenRetailAccount(context.Background(), userId, input)
	if err != nil {
		// Get the logger from the Fiber context
		log := getCtxLogger(ctx)
		log.Errorf("Error opening retail account: %v", err)
		return replyError(ctx, fiber.StatusInternalServerError, err)
	}

	return replyResult(ctx, fiber.Map{"id": newAccountId})
}

func (s *RestServer) UserCheckDepositAmount(ctx *fiber.Ctx) error {
	// Placeholder userID - it would be coming from the JWT or similar auth
	// get "userId" header from ctx
	userId, err := getUserId(ctx)
	if err != nil {
		getCtxLogger(ctx).Errorf("Error parsing userId: %v", err)
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{"error": "Invalid userId"})
	}

	// Read the body of the request into a map
	var input dto.DepositCheckRequest
	if err := ctx.BodyParser(&input); err != nil {
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{
			"error": "Invalid request body",
		})
	}

	// We will need to validate the input, so we use a validator for that
	err = validator.ValidateDepositCheck(input)
	if err != nil {
		return replyError(ctx, fiber.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert Fiber context to a standard context.
	result, err := s.AccountService.CheckDepositAmount(context.Background(), userId, input.AccountId, input.Amount, s.Config)
	if err != nil {
		// Get the logger from the Fiber context
		log := getCtxLogger(ctx)
		log.Errorf("Error checking deposit amount: %v", err)
		return replyError(ctx, fiber.StatusInternalServerError, err)
	}

	return replyResult(ctx, result)
}
