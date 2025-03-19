package config

import (
	"os"
)

// loadConfigFromEnv loads environment variables that match keys in defaults
func loadConfigFromEnv(defaults map[string]interface{}) map[string]interface{} {
	envConfig := make(map[string]interface{})
	for key := range defaults {
		if val, exists := os.LookupEnv(key); exists {
			finalConfigSetValue(envConfig, key, val)
		}
	}
	return envConfig
}
