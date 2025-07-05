package file_reader

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
	RequestFormat string // formato della richiesta (json, yaml, properties)
	FileFormat    string // formato del file (json, yaml, properties)
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

func (fr *FileReader) returnFile() {

	response := map[string]func(fr *FileReader){
		".json":       (*FileReader).returnJsonFile,
		".yaml":       (*FileReader).returnYamlFile,
		".yml":        (*FileReader).returnYamlFile,
		".properties": (*FileReader).returnXmlFile,
	}

	if handler, ok := response[fr.FileFormat]; ok {
		handler(fr)
		return
	} else {
		fr.Ctx.JSON(400, gin.H{"error": "unsupported file type"})
		return
	}
}

func (fr *FileReader) returnJsonFile() {
	j := Json{File: fr.File}
	switch fr.RequestFormat {
	case "yaml", "yml":
		// fai conversione da JSON a YAML
		if fr.File, _ = j.ToYaml(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to YAML failed"})
			return
		}
		fr.returnYamlFile()
		return
	case "properties", "xml":
		// fai conversione da JSON a XML
		if fr.File, _ = j.ToXml(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to XML failed"})
			return
		}
		fr.returnXmlFile()
		return
	}
	fr.Ctx.Data(200, "application/json", fr.File)
}

func (fr *FileReader) returnYamlFile() {
	y := Yaml{File: fr.File}
	switch fr.RequestFormat {
	case "json":
		// fai conversione da YAML a JSON
		if fr.File, _ = y.ToJson(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to JSON failed"})
			return
		}
		fr.returnJsonFile()
		return
	case "properties", "xml":
		// fai conversione da YAML a XML
		if fr.File, _ = y.ToXml(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to XML failed"})
			return
		}
		fr.returnXmlFile()
		return
	}
	fr.Ctx.Data(200, "text/plain; charset=utf-8", fr.File)
}

func (fr *FileReader) returnXmlFile() {
	x := Xml{File: fr.File}
	switch fr.RequestFormat {
	case "json":
		// fai conversione da XML a JSON
		if fr.File, _ = x.ToJson(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to JSON failed"})
			return
		}
		fr.returnJsonFile()
		return
	case "yaml", "yml":
		// fai conversione da XML a YAML
		if fr.File, _ = x.ToYaml(); fr.File == nil {
			fr.Ctx.JSON(500, gin.H{"error": "conversion to YAML failed"})
			return
		}
		fr.returnYamlFile()
		return
	}
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
