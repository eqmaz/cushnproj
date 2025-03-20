package config

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"cushon/internal/application/apperr"
	"cushon/pkg/e"
	"cushon/pkg/optional"
	"cushon/pkg/util"
)

// TestTryConfigFile tests the tryConfigFile function
func TestTryConfigFile(t *testing.T) {
	// Create a temporary test file
	tempFile, err := os.CreateTemp("", "config_test_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()
	defer os.Remove(tempFilePath) // Clean up after test

	tests := []struct {
		name         string
		inputPath    string
		expectExists bool
	}{
		{"Valid config file", tempFilePath, true},
		{"Non-existent file", "non_existent_file.json", false},
		{"Empty path", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resultPath, exists := tryConfigFile(tc.inputPath)
			expectedPath, _ := filepath.Abs(tc.inputPath)

			if exists != tc.expectExists {
				t.Errorf("Expected exists=%v, got %v", tc.expectExists, exists)
			}
			if exists && resultPath != expectedPath {
				t.Errorf("Expected path %q, got %q", expectedPath, resultPath)
			}
		})
	}
}

// Function wrappers for mocking
var (
	tryConfigFileMock = tryConfigFile
	isExeInCwdMock    = util.IsExeInCwd
	exeDirMock        = util.ExeDir
)

// TestGetConfigFilePath tests the GetConfigFilePath function
func TestGetConfigFilePath(t *testing.T) {
	e.SetCatalogue(apperr.GetErrorMap())

	// Create a temporary config file
	tempFile, err := os.CreateTemp("", "cfg_test_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()
	defer os.Remove(tempFilePath) // Clean up

	mockExeDir := filepath.Dir(tempFilePath)

	tests := []struct {
		name          string
		argConfigPath string
		preferences   Preferences
		mockExeInCwd  bool
		mockExeDirErr error
		expectError   bool
		expectedPath  string
	}{
		{
			name:          "Config file provided as argument",
			argConfigPath: tempFilePath,
			preferences:   DefaultPreferences(),
			expectError:   false,
			expectedPath:  tempFilePath,
		},
		{
			name:          "Config not found in any path",
			argConfigPath: "",
			preferences:   DefaultPreferences(),
			expectError:   true,
			expectedPath:  "",
		},
	}

	// Save original function implementations
	originalTryConfigFile := tryConfigFileMock
	originalIsExeInCwd := isExeInCwdMock
	originalExeDir := exeDirMock

	defer func() {
		tryConfigFileMock = originalTryConfigFile
		isExeInCwdMock = originalIsExeInCwd
		exeDirMock = originalExeDir
	}()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset the flag package to prevent redefinition apperr
			flag.CommandLine = flag.NewFlagSet(tc.name, flag.ContinueOnError)

			// Mock tryConfigFile function
			tryConfigFileMock = func(path string) (string, bool) {
				if path == tempFilePath || path == filepath.Join(mockExeDir, "cfg.json") {
					return path, true
				}
				return "", false
			}

			// Mock IsExeInCwd function
			isExeInCwdMock = func() (bool, error) {
				return tc.mockExeInCwd, nil
			}

			// Mock ExeDir function
			exeDirMock = func() (string, error) {
				if tc.mockExeDirErr != nil {
					return "", tc.mockExeDirErr
				}
				return mockExeDir, nil
			}

			// Set command-line argument mock
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()
			if tc.argConfigPath != "" {
				os.Args = []string{"test_binary", "--cfg", tc.argConfigPath}
			} else {
				os.Args = []string{"test_binary"}
			}

			// Run the function using `optional.Some`
			result, err := GetConfigFilePath(optional.Some("cfg.json"), tc.preferences)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				} else {
					expectedErr := e.FromCode("eNcF01")
					if err.Error() != expectedErr.Error() {
						t.Errorf("Expected error %q, got %q", expectedErr.Error(), err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if result != tc.expectedPath {
					t.Errorf("Expected path %q, got %q", tc.expectedPath, result)
				}
			}
		})
	}
}
