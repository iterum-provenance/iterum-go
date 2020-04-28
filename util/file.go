package util

import "os"

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks whether path is an existing dir
func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// IsFileOrDir returns whether the passed path is an existing file or dir
func IsFileOrDir(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
