package config

import (
	"os"

	"github.com/rojack96/vierno/helpers"
	"gopkg.in/yaml.v3"
)

type DashboardConfig struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Title   string `json:"title" yaml:"title"`
}

type Vierno struct {
	Production          bool   `json:"production" yaml:"production"`
	Port                string `json:"port" yaml:"port"`
	DefaultFormatReturn string `json:"defaultFormatReturn" yaml:"defaultFormatReturn"`
	Auth                struct {
		Enabled  bool   `json:"enabled" yaml:"enabled"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	} `json:"auth" yaml:"auth"`
}

type ViernoConfig struct {
	Vierno    Vierno          `json:"vierno" yaml:"vierno"`
	Git       helpers.Git     `json:"git" yaml:"git"`
	Dashboard DashboardConfig `json:"dashboard" yaml:"dashboard"`
}

func ReadViernoConfig(devMode bool) (*ViernoConfig, error) {

	filePath := "vierno.config.yml"
	if devMode {
		filePath = "../../vierno-config-server/vierno.config.yml"

	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ViernoConfig
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
