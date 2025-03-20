package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestFinalConfigMerge tests the finalConfigMerge function
func TestFinalConfigMerge(t *testing.T) {
	tests := []struct {
		name     string
		target   map[string]interface{}
		source   map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Merge non-overlapping keys",
			target:   map[string]interface{}{"A": 1},
			source:   map[string]interface{}{"B": 2},
			expected: map[string]interface{}{"A": 1, "B": 2},
		},
		{
			name:     "Overwrite existing key",
			target:   map[string]interface{}{"A": 1},
			source:   map[string]interface{}{"A": 99},
			expected: map[string]interface{}{"A": 99},
		},
		{
			name:     "Source is empty",
			target:   map[string]interface{}{"A": 1, "B": 2},
			source:   map[string]interface{}{},
			expected: map[string]interface{}{"A": 1, "B": 2}, // Should remain unchanged
		},
		{
			name:     "Target is empty",
			target:   map[string]interface{}{},
			source:   map[string]interface{}{"A": 1, "B": 2},
			expected: map[string]interface{}{"A": 1, "B": 2}, // Should copy everything
		},
		{
			name:     "Merging multiple keys with overwrite",
			target:   map[string]interface{}{"A": 1, "B": 2},
			source:   map[string]interface{}{"B": 99, "C": 3},
			expected: map[string]interface{}{"A": 1, "B": 99, "C": 3}, // B is replaced, C is added
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Execute function
			finalConfigMerge(tc.target, tc.source)

			// Compare with expected result
			if !reflect.DeepEqual(tc.target, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.target)
			}
		})
	}
}

// TestLoadConfigFromFile tests the loadConfigFromFile function
func TestLoadConfigFromFile(t *testing.T) {
	// Create temporary test directory
	tempDir := t.TempDir()

	// Valid JSON content
	validJSON := `{"key1": "value1", "key2": 42}`
	invalidJSON := `{"key1": "value1", "key2": }` // Malformed JSON
	emptyJSON := `{}`

	// Helper function to create a temp file
	createTempFile := func(name, content string) string {
		path := filepath.Join(tempDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		return path
	}

	// Create test files
	validFile := createTempFile("valid.json", validJSON)
	invalidFile := createTempFile("invalid.json", invalidJSON)
	emptyFile := createTempFile("empty.json", emptyJSON)
	nonExistentFile := filepath.Join(tempDir, "non_existent.json")

	tests := []struct {
		name        string
		filePath    string
		expectError bool
		expected    map[string]interface{}
	}{
		{
			name:        "Valid JSON file",
			filePath:    validFile,
			expectError: false,
			expected:    map[string]interface{}{"key1": "value1", "key2": float64(42)}, // JSON numbers default to float64
		},
		{
			name:        "Invalid JSON file",
			filePath:    invalidFile,
			expectError: true,
			expected:    nil,
		},
		{
			name:        "Non-existent file",
			filePath:    nonExistentFile,
			expectError: true,
			expected:    nil,
		},
		{
			name:        "Empty JSON file",
			filePath:    emptyFile,
			expectError: false,
			expected:    map[string]interface{}{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := loadConfigFromFile(tc.filePath)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				// Compare expected and actual JSON maps
				expectedJSON, _ := json.Marshal(tc.expected)
				resultJSON, _ := json.Marshal(result)

				if string(expectedJSON) != string(resultJSON) {
					t.Errorf("Expected %s, got %s", expectedJSON, resultJSON)
				}
			}
		})
	}
}

// TestResolveConfigFilePath tests the resolveConfigFilePath function
func TestResolveConfigFilePath(t *testing.T) {
	// Create temporary test directory
	tempDir := t.TempDir()

	// Create mock config files in temp directory
	configFile := filepath.Join(tempDir, "config.json")
	configDirFile := filepath.Join(tempDir, "config/config.json")

	createTempFile := func(path string) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatalf("Failed to create directory for %s: %v", path, err)
		}
		if err := os.WriteFile(path, []byte(`{"test": "data"}`), 0644); err != nil {
			t.Fatalf("Failed to create file %s: %v", path, err)
		}
	}

	tests := []struct {
		name     string
		userPath string
		setup    func(string) // Passes tempDir to setup function
		expected string
	}{
		{
			name:     "User provides a path",
			userPath: "/custom/path/config.json",
			expected: "/custom/path/config.json",
		},
		{
			name:     "Config file exists in ./config.json",
			userPath: "",
			setup: func(tempDir string) {
				createTempFile(filepath.Join(tempDir, "config.json"))
			},
			expected: configFile,
		},
		{
			name:     "Config file exists in ./config/config.json",
			userPath: "",
			setup: func(tempDir string) {
				createTempFile(filepath.Join(tempDir, "config/config.json"))
			},
			expected: configDirFile,
		},
		{
			name:     "No config file found",
			userPath: "",
			setup: func(tempDir string) {
				os.Remove(filepath.Join(tempDir, "config.json"))
				os.RemoveAll(filepath.Join(tempDir, "config")) // Delete `config/config.json`
			},
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Run setup if defined
			if tc.setup != nil {
				tc.setup(tempDir)
			}

			// Temporarily change working directory to tempDir
			originalWd, _ := os.Getwd()
			_ = os.Chdir(tempDir)
			defer os.Chdir(originalWd)

			// Run function
			result := resolveConfigFilePath(tc.userPath)

			// Check result
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
