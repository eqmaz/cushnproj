package config

import (
	"flag"
)

// GetConfigFileFromArgs -retrieves cfg file path from --cfg flag
func getConfigFilePathFromArgs() string {
	configPath := flag.String("cfg", "", "Path to configuration file")
	flag.Parse()

	if *configPath != "" {
		return *configPath
	}
	return ""
}
