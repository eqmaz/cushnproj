package config

import (
	"os"
	"path/filepath"

	c "cushon/pkg/console"
	"cushon/pkg/e"
	. "cushon/pkg/optional" // Dot import is intentional and okay here
	"cushon/pkg/util"
)

// TryConfigFile validates existence and readability of cfg file
func tryConfigFile(path string) (string, bool) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}
	if _, err := os.Stat(fullPath); err == nil {
		return fullPath, true
	}
	return "", false
}

// GetConfigFilePath attempts to locate the cfg file in the current working directory and the executable directory
// First checks if a CLI argument was provided, then tries the default paths
// Config file name defaults to "cfg.json"
func GetConfigFilePath(name Optional[string], preferences Preferences) (string, error) {
	if name.IsNone() {
		name = Some(preferences.ConfigFileName)
	}

	// Get the config file path from the command line arguments, if provided
	configPath := getConfigFilePathFromArgs()
	if configPath != "" {
		fullPath, ok := tryConfigFile(configPath)
		if ok {
			c.Successf("Using cfg file from argument at: '%s'", fullPath)
			return fullPath, nil
		} else {
			c.Warnf("No cfg file at '%s'...", configPath)
		}
	}

	// Check if the executable is in the current working directory already
	// If this fails, we will not care, and just try the default paths
	var pathsToTry = preferences.ConfigFilePaths
	exeInCwd, err := util.IsExeInCwd()
	if err == nil && !exeInCwd {
		// If not (cwd is not the same as the exe dir), try the exe dir for the cfg files as well
		exeDir, err := util.ExeDir()
		if err == nil {
			pathsToTry = append(pathsToTry, exeDir)
			pathsToTry = append(pathsToTry, filepath.Join(exeDir, "cfg/"))
		}
	}

	// Try the possible cfg paths, starting with the provided path
	for _, path := range pathsToTry {
		//c.Infof("Trying the config file at: '%s'", path)
		fullPath, ok := tryConfigFile(
			filepath.Join(path, name.Unwrap()),
		)
		if ok {
			return fullPath, nil
		}
	}

	return "", e.FromCode("eNcF01").SetField("pathsTried", pathsToTry)
}
