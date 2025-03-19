package dto

type Fund struct {
	FundUUID string `json:"fund_uuid"`
	Weight   int    `json:"weight"`
}

type FundsResponse struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
