package config

import (
	"reflect"
	"testing"
)

// TestSetNestedValue tests the SetNestedValue function
func TestSetNestedValue(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]interface{}
		path     string
		value    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Set top-level key",
			initial:  map[string]interface{}{},
			path:     "port",
			value:    8080,
			expected: map[string]interface{}{"port": 8080},
		},
		{
			name:    "Set nested key",
			initial: map[string]interface{}{},
			path:    "http_server.port",
			value:   8080,
			expected: map[string]interface{}{
				"http_server": map[string]interface{}{
					"port": 8080,
				},
			},
		},
		{
			name: "Overwrite existing nested key",
			initial: map[string]interface{}{
				"http_server": map[string]interface{}{
					"port": 3000,
				},
			},
			path:  "http_server.port",
			value: 8080,
			expected: map[string]interface{}{
				"http_server": map[string]interface{}{
					"port": 8080,
				},
			},
		},
		{
			name:    "Set deeply nested key",
			initial: map[string]interface{}{},
			path:    "server.http.settings.timeout",
			value:   30,
			expected: map[string]interface{}{
				"server": map[string]interface{}{
					"http": map[string]interface{}{
						"settings": map[string]interface{}{
							"timeout": 30,
						},
					},
				},
			},
		},
		{
			name: "Handle existing non-map value",
			initial: map[string]interface{}{
				"server": "invalid",
			},
			path:  "server.http.port",
			value: 8080,
			expected: map[string]interface{}{
				"server": "invalid", // Should not change, as it prevents map overwrite
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Copy initial map to prevent mutation issues across test cases
			cfg := make(map[string]interface{})
			for k, v := range tc.initial {
				cfg[k] = v
			}

			// Execute function
			defer func() {
				if r := recover(); r != nil && tc.name == "Handle existing non-map value" {
					// Expected panic, ignore
				} else if r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()
			SetNestedValue(cfg, tc.path, tc.value)

			// Compare with expected result
			if !reflect.DeepEqual(cfg, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, cfg)
			}
		})
	}
}

// TestGetNestedValue tests the GetNestedValue function
func TestGetNestedValue(t *testing.T) {
	configData := map[string]interface{}{
		"port": 8080,
		"http_server": map[string]interface{}{
			"port": 3000,
			"settings": map[string]interface{}{
				"timeout": 30,
			},
		},
		"server": "invalid",
	}

	tests := []struct {
		name     string
		config   map[string]interface{}
		path     string
		expected interface{}
	}{
		{
			name:     "Retrieve top-level value",
			config:   configData,
			path:     "port",
			expected: 8080,
		},
		{
			name:     "Retrieve nested value",
			config:   configData,
			path:     "http_server.port",
			expected: 3000,
		},
		{
			name:     "Retrieve deeply nested value",
			config:   configData,
			path:     "http_server.settings.timeout",
			expected: 30,
		},
		{
			name:     "Key does not exist",
			config:   configData,
			path:     "http_server.nonexistent",
			expected: nil,
		},
		{
			name:     "Partial path exists but is not a map",
			config:   configData,
			path:     "server.http.port",
			expected: nil, // "server" is a string, not a map
		},
		{
			name:     "Empty path",
			config:   configData,
			path:     "",
			expected: configData,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GetNestedValue(tc.config, tc.path)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For path %q, expected %+v, got %+v", tc.path, tc.expected, result)
			}
		})
	}
}
