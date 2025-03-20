package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func finalConfigMerge(target, source map[string]interface{}) {
	for key, val := range source {
		target[key] = val
	}
}

// Helper functions for nested map values
// TODO
func finalConfigSetValue(cfg map[string]interface{}, key string, val interface{}) {
	cfg[key] = val // simplified; could use dot-notation support if needed
}

// LoadConfig loads the configuration following this order of precedence:
// Defaults -> Config File -> Environment Variables (priority configurable in preferences).
func LoadConfig(configFilePath string, prefs Preferences, defaults map[string]interface{}, configStruct interface{}) error {
	// Start with defaults
	finalConfig := defaults

	// Load from file if provided
	if configFilePath := resolveConfigFilePath(configFilePath); configFilePath != "" {
		fileConfig, err := loadConfigFromFile(configFilePath)
		if err != nil {
			if !prefs.NoErrorOnMissingFile {
				return fmt.Errorf("error loading cfg file: %w", err)
			}
			fmt.Printf("Warning: cfg file not loaded: %v\n", err)
		} else {
			finalConfigMerge(finalConfig, fileConfig)
		}
	}

	// Load environment variables and merge
	envConfig := loadConfigFromEnv(defaults)
	if prefs.PriorityEnv {
		finalConfigMerge(finalConfig, envConfig)
	} else {
		temp := finalConfig
		finalConfig = envConfig
		finalConfigMerge(finalConfig, temp)
	}

	// Convert to user-defined struct
	configBytes, err := json.Marshal(finalConfig)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(configBytes, configStruct); err != nil {
		return err
	}

	return nil
}

// resolveConfigFilePath resolves the final cfg file path
func resolveConfigFilePath(userPath string) string {
	if userPath != "" {
		return userPath
	}
	// default search locations
	possiblePaths := []string{
		"./config/config.json",
		"./config.json",
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			absPath, _ := filepath.Abs(path) // Convert to absolute path
			return absPath
		}
	}
	return ""
}

// loadConfigFromFile reads a JSON cfg file into a map
func loadConfigFromFile(filePath string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read cfg file: %v", err)
	}
	var cfg map[string]interface{}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse cfg file: %v", err)
	}
	return cfg, nil
}
