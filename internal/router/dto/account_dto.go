package dto

type AccountBalancesResponse struct {
	ProductId       string  `db:"product_id"`
	ProductTitle    string  `db:"product_title"`
	TotalInvestment float64 `db:"total_investment"`
	CurrentBalance  float64 `db:"current_balance"`
}

type AccountCreationRequest struct {
	ProductUUID string `json:"product_uuid"`
	FundUUIDs   []Fund `json:"fund_uuids"`
}

type DepositCheckRequest struct {
	AccountId uint64  `json:"account_id"`
	Amount    float64 `json:"amount"`
}
