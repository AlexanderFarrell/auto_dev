
package main

import (
	"encoding/json"
	"os"
)

type ConfigItem struct {
	Type string `json:"type"`
	Path string `json:"path"`
	Output *string `json:"output"`
}

type Config = map[string]ConfigItem

func LoadConfig() (Config, error) {
	path := "auto_dev.json"

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

