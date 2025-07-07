package file_getter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Controller for endpoint /:app/:profile?origin=json|yaml|properties
func GetSimpleFile(c *gin.Context) {
	// TODO add middleware to recognize if is logged or not
	// const BasePath = "./app"
	// Development path
	const BasePath = "../../vierno-config-server/app"

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	app := c.Param("app")
	profile := c.Param("profile")
	isOriginal := c.Query("original") == "true"

	folder := filepath.Join(BasePath, app)

	path, err := findFileByName(folder, profile)
	if err != nil {
		fmt.Println("Errore nella ricerca del file:", err)
	}

	_, fr.FileFormat = fileNameExt(path)
	fr.File, err = os.ReadFile(path)
	if err != nil {
		c.JSON(500, gin.H{"error": "file not found or cannot be read"})
		return
	}

	if isOriginal {
		fr.returnOriginalFile()
		return
	}

	fr.returnConfigFile()
}

func GetFile(c *gin.Context) {
	// TODO add middleware to recognize if is logged or not
	const BasePath = "./app"

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	filename := c.Param("app")
	isOriginal := c.Query("original") == "true"

	req := strings.Split(filename, ".")
	app := req[0]
	profile := req[1]

	folder := filepath.Join(BasePath, app)

	path, err := findFileByName(folder, profile)
	if err != nil {
		fmt.Println("Errore nella ricerca del file:", err)
	}

	_, fr.FileFormat = fileNameExt(path)
	fr.File, err = os.ReadFile(path)
	if err != nil {
		c.JSON(500, gin.H{"error": "file not found or cannot be read"})
		return
	}

	if isOriginal {
		fr.returnOriginalFile()
		return
	}

	fr.returnConfigFile()
}
