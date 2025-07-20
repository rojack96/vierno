package file_getter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/helpers"
)

// Controller for endpoint /:app/:profile?origin=json|yaml|properties
func GetSimpleFile(c *gin.Context) {
	devMode := c.MustGet("devMode").(bool)
	// TODO add middleware to recognize if is logged or not

	basePath := "./app"
	if devMode {
		basePath = "../../vierno-config-server/app"
	}
	// Development path

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	app := c.Param("app")
	profile := c.Param("profile")
	isOriginal := c.Query("original") == "true"

	folder := filepath.Join(basePath, app)

	path, err := helpers.FindFileByName(folder, profile)
	if err != nil {
		fmt.Println("Errore nella ricerca del file:", err)
	}

	_, fr.FileFormat = helpers.FileNameExt(path)
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
	devMode := c.MustGet("devMode").(bool)
	// TODO add middleware to recognize if is logged or not

	basePath := "./app"
	if devMode {
		basePath = "../../vierno-config-server/app"
	}

	fr := FileReader{Ctx: c}
	fr.checkout()

	// Request parameters
	filename := c.Param("app")

	req := strings.Split(filename, ".")
	app := req[0]
	profile := req[1]
	fr.RequestFileFormat = &req[2]

	folder := filepath.Join(basePath, app)

	path, err := helpers.FindFileByName(folder, profile)
	if err != nil {
		fmt.Println("Errore nella ricerca del file:", err)
	}

	_, fr.FileFormat = helpers.FileNameExt(path)
	fr.File, err = os.ReadFile(path)
	if err != nil {
		c.JSON(500, gin.H{"error": "file not found or cannot be read"})
		return
	}

	fr.returnConfigFile()
}
