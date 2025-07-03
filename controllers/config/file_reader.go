package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSimpleFile(c *gin.Context) {
	basePath := "./app"

	app := c.Param("app")
	profile := c.Param("profile")
	format := c.Query("format")

	if format == "" {
		format = "json" // Default format if not specified
	}

	folder := filepath.Join(basePath, app)

	path, err := findFileByName(folder, profile)
	if err != nil {
		fmt.Println("Errore nella ricerca del file:", err)
	}

	_, ext := fileNameExt(path)
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		c.JSON(500, gin.H{"error": "error reading file"})
		return
	}

	switch ext {
	case ".json":
		returnJsonFile(c, fileBytes)
		return
	case ".yaml", ".yml":
		returnYamlFile(c, fileBytes)
		return
	case ".properties":
		returnXmlFile(c, fileBytes)
		return
	default:
		c.JSON(400, gin.H{"error": "unsupported file type"})
		return
	}
}

func returnJsonFile(c *gin.Context, fileBytes []byte) {
	c.Data(200, "application/json", fileBytes)
}

func returnYamlFile(c *gin.Context, fileBytes []byte) {
	c.Data(200, "text/plain; charset=utf-8", fileBytes)
}

func returnXmlFile(c *gin.Context, fileBytes []byte) {
	c.Data(200, "application/xml", fileBytes)
}

/*func GetFile(c *gin.Context) {
	reqName, reqExt := fileNameRequest(c)

	basePath := "./app"

	fileBytes, err := os.ReadFile("./app/config.json")
	if err != nil {
		fmt.Println("Errore nella lettura del file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossibile leggere il file"})
		return
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileBytes, &jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File JSON non valido"})
		return
	}

	yamlBytes, err := yaml.Marshal(jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nella conversione in YAML"})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", yamlBytes)
}*/

func fileNameExt(filename string) (string, string) {
	ext := strings.ToLower(filepath.Ext(filename))
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}

// Trova il primo file in 'folder' che ha il nome base uguale a 'targetName' (senza considerare l'estensione).
func findFileByName(folder, targetName string) (string, error) {
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
