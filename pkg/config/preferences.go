package config

// Preferences control the configuration loading behavior.
type Preferences struct {
	// Default config file name to look for (not the path)
	ConfigFileName string

	// Default paths to try the cfg file, if not provided in the arguments
	// These are relative to the current working directory!
	ConfigFilePaths []string

	// PriorityEnv: true means Env vars override file settings
	PriorityEnv bool

	// NoErrorOnMissingFile controls whether missing cfg file raises an error
	NoErrorOnMissingFile bool
}

// DefaultPreferences returns default preference settings
func DefaultPreferences() Preferences {
	return Preferences{
		ConfigFileName: "config.json",
		ConfigFilePaths: []string{
			"./",
			"./config/",
			"/etc/",
		},
		PriorityEnv:          true,
		NoErrorOnMissingFile: false,
	}
}
