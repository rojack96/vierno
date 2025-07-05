package file_reader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {
	// TODO add middleware to recognize if is logged or not
	const BasePath = "./app"

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	filename := c.Param("app")

	req := strings.Split(filename, ".")
	app := req[0]
	profile := req[1]
	fr.RequestFormat = req[2]

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

	fr.returnFile()
}

// Controller for endpoint /:app/:profile?format=json|yaml|properties
func GetSimpleFile(c *gin.Context) {
	// TODO add middleware to recognize if is logged or not
	const BasePath = "./app"

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	app := c.Param("app")
	profile := c.Param("profile")
	fr.RequestFormat = c.Query("format")

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

	fr.returnFile()
}
