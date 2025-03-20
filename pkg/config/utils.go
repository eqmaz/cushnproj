package config

import (
	"fmt"
	"strings"
)

// SetNestedValue sets a nested map value given a dot-notation key ("http_server.port")
func SetNestedValue(cfg map[string]interface{}, path string, val interface{}) {
	keys := strings.Split(path, ".")
	lastKeyIndex := len(keys) - 1
	for i, key := range keys {
		if i == lastKeyIndex {
			cfg[key] = val
			return
		}

		// ensure the nested map exists
		if _, ok := cfg[key]; !ok {
			cfg[key] = make(map[string]interface{})
		}
		cfg = cfg[key].(map[string]interface{})
	}
}

// GetNestedValue retrieves a nested value using dot notation
func GetNestedValue(cfg map[string]interface{}, path string) interface{} {

	// If the path is empty, return the entire config
	// could just return nil?
	if path == "" {
		fmt.Println("🔍 DEBUG: Empty path detected, returning full config")
		return cfg
	}

	keys := strings.Split(path, ".")
	var current interface{} = cfg
	for _, key := range keys {
		if m, ok := current.(map[string]interface{}); ok {
			current, ok = m[key]
			if !ok {
				return nil
			}
		} else {
			return nil
		}
	}
	return current
}
