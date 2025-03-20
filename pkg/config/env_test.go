package config

import (
	"os"
	"testing"
)

// TestLoadConfigFromEnv tests the loadConfigFromEnv function
func TestLoadConfigFromEnv(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		defaults map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "No environment variables set",
			envVars:  map[string]string{},
			defaults: map[string]interface{}{"DB_HOST": "localhost", "DB_PORT": 3306},
			expected: map[string]interface{}{}, // No overrides, should remain empty
		},
		{
			name: "Some environment variables set",
			envVars: map[string]string{
				"DB_HOST": "remotehost",
			},
			defaults: map[string]interface{}{"DB_HOST": "localhost", "DB_PORT": 3306},
			expected: map[string]interface{}{"DB_HOST": "remotehost"}, // Only DB_HOST should be overridden
		},
		{
			name: "All environment variables set",
			envVars: map[string]string{
				"DB_HOST": "remotehost",
				"DB_PORT": "5432",
			},
			defaults: map[string]interface{}{"DB_HOST": "localhost", "DB_PORT": 3306},
			expected: map[string]interface{}{"DB_HOST": "remotehost", "DB_PORT": "5432"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Backup original environment variables
			origEnv := make(map[string]string)
			for key := range tc.envVars {
				origEnv[key], _ = os.LookupEnv(key)
				os.Setenv(key, tc.envVars[key])
			}

			// Run the function
			result := loadConfigFromEnv(tc.defaults)

			// Restore original environment variables
			for key, val := range origEnv {
				os.Setenv(key, val)
			}
			for key := range tc.envVars {
				if _, exists := origEnv[key]; !exists {
					os.Unsetenv(key)
				}
			}

			// Validate results
			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d entries, got %d", len(tc.expected), len(result))
			}
			for key, expectedValue := range tc.expected {
				if result[key] != expectedValue {
					t.Errorf("For key %q, expected %v, got %v", key, expectedValue, result[key])
				}
			}
		})
	}
}
