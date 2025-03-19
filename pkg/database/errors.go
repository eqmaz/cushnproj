package database

import (
	"strings"
)

// IsDuplicateEntryError - helper to check if error string contains all of "1062" and "Duplicate entry"
// Note this is specific to MySQL
func IsDuplicateEntryError(err error) bool {
	return err != nil &&
		strings.Contains(err.Error(), "1062") &&
		strings.Contains(err.Error(), "Duplicate entry")
}
