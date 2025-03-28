package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// createFile writes content into the specified file.
func createFile(folderPath, fileName, content string) error {
	filePath := filepath.Join(folderPath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to %s: %w", filePath, err)
	}

	return nil
}
