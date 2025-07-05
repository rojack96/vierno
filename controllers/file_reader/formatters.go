package file_reader

import (
	"encoding/json"
	"fmt"

	"github.com/clbanning/mxj/v2"
	"gopkg.in/yaml.v3"
)

/* FileFormat */
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

// XML → YAML
func (x *Xml) ToYaml() ([]byte, error) {
	jsonData, err := x.ToJson()
	if err != nil {
		return nil, err
	}
	j := &Json{File: jsonData}
	return j.ToYaml()
}

// XML → JSON
func (x *Xml) ToJson() ([]byte, error) {
	mv, err := mxj.NewMapXml(x.File)
	if err != nil {
		return nil, err
	}

	// Estrarre la mappa interna
	props, ok := mv["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("struttura XML inattesa (manca 'properties')")
	}

	entries, ok := props["entry"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("struttura XML inattesa (manca 'entry')")
	}

	// Costruisci la mappa semplificata
	result := make(map[string]string)
	for _, e := range entries {
		entryMap, ok := e.(map[string]interface{})
		if !ok {
			continue
		}
		key, ok1 := entryMap["-key"].(string)
		val, ok2 := entryMap["#text"].(string)
		if ok1 && ok2 {
			result[key] = val
		}
	}

	// Convertilo in JSON
	return json.MarshalIndent(result, "", "  ")
}
