package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileNameExt returns the base name and extension of a file.
func FileNameExt(filename string) (string, string) {
	ext := strings.ToLower(filepath.Ext(filename))
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}

// FindFileByName Finds the first file in 'folder' whose base name matches 'targetName' (ignoring the extension).
func FindFileByName(folder, targetName string) (string, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		base := strings.TrimSuffix(name, filepath.Ext(name))
		if base == targetName {
			return filepath.Join(folder, name), nil
		}
	}
	return "", fmt.Errorf("file '%s' non trovato in %s", targetName, folder)
}
