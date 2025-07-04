package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/clbanning/mxj/v2"
	"gopkg.in/yaml.v3"

	"github.com/gin-gonic/gin"
)

func GetSimpleFile(c *gin.Context) {
	const BasePath = "./app"

	fr := FileReader{Ctx: c}

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

	if fr.RequestFormat != "" {
		fr.requestWithFormat()
		return
	}

	fr.returnFile()
}

/* Functions */

type FileReader struct {
	File          []byte
	RequestFormat string // formato della richiesta (json, yaml, properties)
	FileFormat    string // formato del file (json, yaml, properties)
	Ctx           *gin.Context
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
	fr.Ctx.Data(200, "application/json", fr.File)
}

func (fr *FileReader) returnYamlFile() {
	fr.Ctx.Data(200, "text/plain; charset=utf-8", fr.File)
}

func (fr *FileReader) returnXmlFile() {
	fr.Ctx.Data(200, "application/xml", fr.File)
}

func (fr *FileReader) requestWithFormat() {

	if "."+fr.RequestFormat == fr.FileFormat {
		fr.returnFile()
		return
	}

	switch fr.FileFormat {
	case ".json":
		j := &Json{File: fr.File}
		switch fr.RequestFormat {
		case "yaml", "yml":
			// fai conversione da YAML a JSON
			fr.File, _ = j.ToYaml()
			fr.returnYamlFile()
			return
		case "properties", "xml":
			// fai conversione da Properties a JSON
			fr.File, _ = j.ToXml()
			fr.returnXmlFile()
			return
		default:
			fr.Ctx.JSON(400, gin.H{"error": "unsupported file type for JSON conversion"})
			return
		}
	case ".yaml", ".yml":
		y := &Yaml{File: fr.File}
		switch fr.RequestFormat {
		case "json":
			// fai conversione da YAML a JSON
			fr.File, _ = y.ToJson()
			fr.returnJsonFile()
			return
		case "properties":
			// fai conversione da Properties a JSON
			fr.File, _ = y.ToXml()
			fr.returnXmlFile()
			return
		default:
			fr.Ctx.JSON(400, gin.H{"error": "unsupported file type for JSON conversion"})
			return
		}
	case "properties", "xml":
		x := &Xml{File: fr.File}
		switch fr.FileFormat {
		case ".json":
			// fai conversione da Properties a JSON
			fr.File, _ = x.ToJson()
			fr.returnXmlFile()
			return
		case "yaml", "yml":
			// fai conversione da YAML a JSON
			fr.File, _ = x.ToYaml()
			fr.returnYamlFile()
			return
		default:
			fr.Ctx.JSON(400, gin.H{"error": "unsupported file type for JSON conversion"})
			return
		}
	default:
		fr.Ctx.JSON(400, gin.H{"error": "unsupported format"})
		return
	}

}

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

type Json struct {
	File []byte
}

type Yaml struct {
	File []byte
}

type Xml struct {
	File []byte
}

// JSON → YAML
func (j *Json) ToYaml() ([]byte, error) {
	var obj interface{}
	if err := json.Unmarshal(j.File, &obj); err != nil {
		return nil, err
	}
	return yaml.Marshal(obj)
}

// JSON → XML
func (j *Json) ToXml() ([]byte, error) {
	mv, err := mxj.NewMapJson(j.File)
	if err != nil {
		return nil, err
	}
	return mv.XmlIndent("", "  ")
}

// YAML → JSON
func (y *Yaml) ToJson() ([]byte, error) {
	var obj interface{}
	if err := yaml.Unmarshal(y.File, &obj); err != nil {
		return nil, err
	}
	return json.MarshalIndent(obj, "", "  ")
}

// YAML → XML
func (y *Yaml) ToXml() ([]byte, error) {
	jsonData, err := y.ToJson()
	if err != nil {
		return nil, err
	}
	j := &Json{File: jsonData}
	return j.ToXml()
}

// XML → JSON
func (x *Xml) ToJson() ([]byte, error) {
	mv, err := mxj.NewMapXml(x.File)
	if err != nil {
		return nil, err
	}
	return mv.JsonIndent("", "  ")
}

// XML → YAML
func (x *Xml) ToYaml() ([]byte, error) {
	jsonData, err := x.ToJson()
	if err != nil {
		return nil, err
	}
	j := &Json{File: jsonData}
	return j.ToYaml()
}
