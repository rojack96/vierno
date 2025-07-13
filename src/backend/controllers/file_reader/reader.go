package file_reader

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFolders(c *gin.Context) {
	root := "../../vierno-config-server/app"
	entries, err := os.ReadDir(root)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore lettura directory"})
		return
	}

	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{"folders": folders})
}

func GetFolder(c *gin.Context) {
	folder := c.Param("folder")

	// Previeni directory traversal
	if strings.Contains(folder, "..") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Percorso non valido"})
		return
	}

	basePath := "../../vierno-config-server/app"
	targetPath := filepath.Join(basePath, folder)

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore lettura directory"})
		return
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
