package file_reader

import (
	"net/http"
	"os"

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
