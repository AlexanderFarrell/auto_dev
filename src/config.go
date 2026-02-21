package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ConfigItem struct {
	Type      string   `json:"type"`
	Path      string   `json:"path"`
	Output    *string  `json:"output,omitempty"`
	Script    *string  `json:"script,omitempty"`
	OrderFile *string  `json:"order_file,omitempty"`
	Inputs    []string `json:"inputs,omitempty"`
	Compose   *string  `json:"compose,omitempty"`
	Stack     *string  `json:"stack,omitempty"`
	Service   *string  `json:"service,omitempty"`
}

type ConfigFile struct {
	Targets map[string]ConfigItem `json:"targets"`
	Groups  map[string][]string   `json:"groups,omitempty"`
}

func LoadConfig() (ConfigFile, error) {
	path := "auto_dev.json"

	if base := os.Getenv("AUTO_DEV_CWD"); base != "" {
		if data, err := os.ReadFile(filepath.Join(base, path)); err == nil {
			return parseConfig(data)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		exe, exeErr := os.Executable()
		if exeErr != nil {
			return ConfigFile{}, err
		}
		exeDir := filepath.Dir(exe)
		altPath := filepath.Join(exeDir, path)
		data, err = os.ReadFile(altPath)
		if err != nil {
			return ConfigFile{}, err
		}
	}

	return parseConfig(data)
}

func parseConfig(data []byte) (ConfigFile, error) {
	var cfg ConfigFile
	if err := json.Unmarshal(data, &cfg); err != nil {
		return ConfigFile{}, err
	}

	if cfg.Targets == nil {
		var legacy map[string]ConfigItem
		if err := json.Unmarshal(data, &legacy); err != nil {
			return ConfigFile{}, err
		}
		cfg.Targets = legacy
	}

	return cfg, nil
}
