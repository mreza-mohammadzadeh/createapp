package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateProjectStructure creates project directories and files based on a recursive map structure.
func GenerateProjectStructure(basePath string, structure map[string]interface{}) error {
	for name, content := range structure {
		currentPath := filepath.Join(basePath, name)

		switch value := content.(type) {
		case string: // file
			if err := createFile(basePath, name, value); err != nil {
				return fmt.Errorf("failed to create file %s: %w", name, err)
			}
		case map[string]interface{}: // folder or subfolder
			if err := os.MkdirAll(currentPath, 0755); err != nil {
				return fmt.Errorf("failed to create folder %s: %w", currentPath, err)
			}
			if err := GenerateProjectStructure(currentPath, value); err != nil {
				return err
			}
		}
	}
	return nil
}
