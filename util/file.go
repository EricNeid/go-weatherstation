package util

import "os"

// IsFilePresent returns true if a file with the given path exists.
// If the path points to a directory, it also returns false.
func IsFilePresent(filepath string) bool {
	f, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}
