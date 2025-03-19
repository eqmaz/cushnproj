package cfg

// DefaultConfigMap contains default values specific to your app
var defaultConfigMap = map[string]interface{}{
	"Service": map[string]interface{}{
		"Port": 8080,
		"Tls": map[string]interface{}{
			"Enabled": false,
			"Cert":    "",
			"Key":     "",
		},
		"DisableRouterGreeting":  false,
		"MaxFundsPerIsa":         1,
		"MaxIsaAnnualInvestment": 20000,
		"MaxIsaAccountsPerUser":  1,
		"IsaAllowanceResetDate":  "04-06", // MM-DD
	},
	"Console": map[string]interface{}{
		"Enabled": true,
		"Color":   true,
	},
	"Logger": map[string]interface{}{
		"Enabled": true,
		"Level":   "debug",
		"Stream":  "stdout",
	},
	"RateLimit": map[string]interface{}{
		"Enabled":     true,
		"MaxRequests": 10,
		"Timeframe":   30,
	},
	"Health": map[string]interface{}{
		"Http": map[string]interface{}{
			"Enabled": true,
			"Port":    8081,
		},
		"Grpc": map[string]interface{}{
			"Enabled": true,
			"Port":    50051,
		},
	},
}
