package cfg

import (
	"cushon/pkg/config"
)

// Config is the structured configuration specific to your app
// See types_ files for the struct definitions
type Config struct {
	Service   Service   `json:"service"`
	Database  Database  `json:"database"`
	Console   Console   `json:"console"`
	Logger    Logger    `json:"logger"`
	RateLimit RateLimit `json:"rate_limit"`
	Health    Health    `json:"health"`
}

// LoadConfigs - loads the configuration into a strongly typed Config struct on a best-effort basis
func (c *Config) LoadConfigs(path string) error {
	// Just use the default preferences as per /pkg/config/preferences.go
	// These are not app configs - these are the preferences for the config package behaviour
	defaultPref := config.DefaultPreferences()

	// Populate Config with values, as per defaults, config file and environment variables
	err := config.LoadConfig(path, defaultPref, defaultConfigMap, c)
	if err != nil {
		return err
	}
	return nil
}
