package version

import (
	c "cushon/pkg/console"
)

// These values would get overwritten at build time.
var (
	Version   = "0.0.0.0" // format: major.minor.patch.build
	BuildTime = "unknown"
	Commit    = "unknown"
)

// GetVersion returns the version of the application
func GetVersion() string {
	return Version
}

func PrintVersion() {
	c.Infof("App version: %s; Built at: %s [commit: %s]", Version, BuildTime, Commit)
	//fmt.Println("cushon " + Version + BuildTime + " " + Commit)
}
