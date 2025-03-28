package core

import (
	"os"
	"path/filepath"
)

// IsAppNameDuplicate checks if the given app name already exists in the current directory.
func IsAppNameDuplicate(appName string) bool {
	rootPath := filepath.Join(".", appName)
	_, err := os.Stat(rootPath)
	return !os.IsNotExist(err)
}
