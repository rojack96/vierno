package config

import (
	"encoding/json"
	"os"
)

type DashboardConfig struct {
	Enabled bool   `json:"enabled"`
	Title   string `json:"title"`
}

type ViernoConfig struct {
	Port                string          `json:"port"`
	Production          bool            `json:"production"`
	DefaultConfigReturn string          `json:"defaultConfigReturn"`
	Dashboard           DashboardConfig `json:"dashboard"`
}

func ReadViernoConfig() (*ViernoConfig, error) {
	fileBytes, err := os.ReadFile("vierno.config.json")
	if err != nil {
		return nil, err
	}

	var config ViernoConfig
	if err := json.Unmarshal(fileBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
