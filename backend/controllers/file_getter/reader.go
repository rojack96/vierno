package file_getter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

/* FileReader */

type FileReader struct {
	File          []byte
	FileFormat    string // formato del file (json, yaml, properties)
	DefaultReturn string // formato di ritorno predefinito (json, yaml, xml)
	Ctx           *gin.Context
}

// checkout checks if a branch is specified in the request and performs a checkout if needed.
func (fr *FileReader) checkout() {
	branch := fr.Ctx.Query("branch")

	if branch != "" {
		// TODO implementare il checkout del branch
		// Questo è un placeholder per il codice che gestirà il checkout del branch
		fmt.Println("Eseguendo il checkout del branch:", branch)
		// Qui dovresti aggiungere la logica per cambiare il branch nel repository git
	}
}

func (fr *FileReader) returnOriginalFile() {

	response := map[string]func(fr *FileReader){
		".json":       (*FileReader).returnJsonFile,
		".yaml":       (*FileReader).returnYamlFile,
		".yml":        (*FileReader).returnYamlFile,
		".properties": (*FileReader).returnXmlFile,
		".xml":        (*FileReader).returnXmlFile,
	}

	if handler, ok := response[fr.FileFormat]; ok {
		handler(fr)
		return
	} else {
		fr.Ctx.JSON(400, gin.H{"error": "unsupported file type"})
		return
	}
}

func (fr *FileReader) returnConfigFile() {
	switch fr.DefaultReturn {
	case "json":
		fr.defaultByJson()
	case "yaml", "yml":
		fr.defaultByYaml()
	case "properties", "xml":
		fr.defaultByXml()
	default:
		fr.defaultByJson()
	}
}

func (fr *FileReader) defaultByJson() {
	var (
		transformer JsonTransformer
		err         error
	)

	switch fr.FileFormat {
	case ".yaml", ".yml":
		transformer = &Yaml{File: fr.File}
	case ".properties", ".xml":
		transformer = &Xml{File: fr.File}
	}

	if transformer == nil {
		fr.Ctx.JSON(500, gin.H{"error": "internal error: transformer not initialized"})
		return
	}

	if fr.File, err = transformer.ToJson(); err != nil {
		fr.Ctx.JSON(500, gin.H{"error": "conversion to json failed", "details": err.Error()})
		return
	}
	fr.returnJsonFile()
}

func (fr *FileReader) defaultByYaml() {
	var (
		transformer YamlTransformer
		err         error
	)
	switch fr.FileFormat {
	case ".json":
		transformer = &Json{File: fr.File}
	case ".properties", ".xml":
		transformer = &Xml{File: fr.File}
	}

	if transformer == nil {
		fr.Ctx.JSON(500, gin.H{"error": "internal error: transformer not initialized"})
		return
	}

	if fr.File, err = transformer.ToYaml(); err != nil {
		fr.Ctx.JSON(500, gin.H{"error": "conversion to json failed", "details": err.Error()})
		return
	}
	fr.returnYamlFile()
}

func (fr *FileReader) defaultByXml() {
	var (
		transformer XmlTransformer
		err         error
	)
	switch fr.FileFormat {
	case ".json":
		transformer = &Json{File: fr.File}
	case ".yaml", ".yml":
		transformer = &Yaml{File: fr.File}
	}
	if transformer == nil {
		fr.Ctx.JSON(500, gin.H{"error": "internal error: transformer not initialized"})
		return
	}
	if fr.File, err = transformer.ToXml(); err != nil {
		fr.Ctx.JSON(500, gin.H{"error": "conversion to xml failed", "details": err.Error()})
		return
	}
	fr.returnXmlFile()
}

func (fr *FileReader) returnJsonFile() {
	fr.Ctx.Data(200, "application/json", fr.File)
}

func (fr *FileReader) returnYamlFile() {
	fr.Ctx.Data(200, "text/plain; charset=utf-8", fr.File)
}

func (fr *FileReader) returnXmlFile() {
	fr.Ctx.Data(200, "application/xml", fr.File)
}

// fileNameExt returns the base name and extension of a file.
func fileNameExt(filename string) (string, string) {
	ext := strings.ToLower(filepath.Ext(filename))
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}

// Finds the first file in 'folder' whose base name matches 'targetName' (ignoring the extension).
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
