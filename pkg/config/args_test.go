package config

import (
	"flag"
	"os"
	"testing"
)

// TestGetConfigFilePathFromArgs tests the getConfigFilePathFromArgs function
func TestGetConfigFilePathFromArgs(t *testing.T) {
	// Save original command-line arguments
	origArgs := os.Args
	defer func() { os.Args = origArgs }() // Restore original args after test

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"No flag provided", []string{"test_binary"}, ""},
		{"Flag provided with value", []string{"test_binary", "--cfg", "config.json"}, "config.json"},
		{"Flag provided with empty value", []string{"test_binary", "--cfg", ""}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset flags before parsing new ones
			flag.CommandLine = flag.NewFlagSet(tc.name, flag.ContinueOnError)
			os.Args = tc.args

			result := getConfigFilePathFromArgs()
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
