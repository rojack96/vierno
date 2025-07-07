package config

import (
	"encoding/json"
	"os"
)

type DashboardConfig struct {
	Enabled bool   `json:"enabled"`
	Title   string `json:"title"`
}

type Vierno struct {
	Folder              string `json:"folder"`
	Production          bool   `json:"production"`
	Port                string `json:"port"`
	DefaultFormatReturn string `json:"defaultFormatReturn"`
}

type Git struct {
	Repo    string `json:"repo"`
	Branch  string `json:"branch"`
	AuthKey string `json:"authKey"`
}

type ViernoConfig struct {
	Vierno    Vierno          `json:"vierno"`
	Git       Git             `json:"git"`
	Dashboard DashboardConfig `json:"dashboard"`
}

func ReadViernoConfig() (*ViernoConfig, error) {
	// filePath := "vierno.config.json"
	// Only development, so the file is in the parent directory
	filePath := "../vierno.config.json"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ViernoConfig
	if err := json.Unmarshal(fileBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
