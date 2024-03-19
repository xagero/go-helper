package helper

import "os"

// IsFileExists check if file exists
func IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
