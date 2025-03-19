package cfg

// Service - Configurations for business logic within this service
type Service struct {
	// Port - The port on which the REST service will listen
	Port uint64 `json:"port"`

	// Tls - TLS configuration for the service
	Tls Tls `json:"tls"`

	// DisableRouterGreeting - When true, the "Fiber" greeting won't appear on the CLI at app start
	DisableRouterGreeting bool `json:"disable_router_greeting"`

	// MaxFundsPerIsa - Maximum number of funds allowed per account for the ISA product
	MaxFundsPerIsa uint16 `json:"max_funds_per_isa"`

	// MaxIsaAnnualInvestment - Maximum annual investment allowed for the ISA product
	MaxIsaAnnualInvestment uint32 `json:"max_isa_annual_investment"`

	// MaxIsaAccountsPerUser - Maximum number of ISA accounts allowed per user
	MaxIsaAccountsPerUser uint8 `json:"max_isa_accounts_per_user"`

	// IsaAllowanceResetDate - The date on which the ISA allowance resets in the UK
	IsaAllowanceResetDate string `json:"isa_allowance_reset_date"`
}

type Tls struct {
	Enabled bool   `json:"enabled"` // Whether TLS is enabled
	Cert    string `json:"cert"`    // Cert file path
	Key     string `json:"key"`     // Key file path
}
