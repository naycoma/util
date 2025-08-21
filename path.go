package util

import (
	"os"
	"path/filepath"
)

// FindUp searches for a file by walking up parent directories from the current directory.
// If the file is found, it returns the absolute path and true.
// If the file is not found up to the root directory, it returns an empty string and false.
func FindUp(file string) (path string, ok bool) {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	for {
		path = filepath.Join(currentDir, file)
		ok = IsExist(path)
		if ok {
			return
		}
		parent := filepath.Dir(currentDir)
		// root
		if currentDir == parent {
			return
		}
		currentDir = parent
	}
}

// IsExist checks if a file or directory exists.
func IsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// ProjectRoot returns the absolute path of the project root directory by finding the go.mod file.
// It returns an empty string if go.mod is not found.
func ProjectRoot() string {
	mod, _ := FindUp("go.mod")
	return filepath.Dir(mod)
}
