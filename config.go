package main

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Config struct {
	Editor     string `json:"editor"`
	StorageDir string `json:"storage_dir"`
}

func (c *Config) process() {
	if c.Editor == "" {
		c.Editor = "vi"
	}

	if c.StorageDir == "" {
		c.StorageDir = "~/notes"
	}

	c.StorageDir = expandHome(c.StorageDir)
}

var defaultConfig *Config

func DefaultConfig() *Config {
	if defaultConfig == nil {
		defaultConfig = &Config{}
		defaultConfig.process()
	}

	return defaultConfig
}

func getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir = path.Join(configDir, appName)

	return configDir, nil
}

func expandHome(p string) string {
	if len(p) > 0 && p[0] == '~' {
		home, _ := os.UserHomeDir()
		return path.Join(home, p[1:])
	}
	return p
}

func LoadConfig() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	configFilePath := path.Join(configDir, "config.jsonc")

	raw, err := os.ReadFile(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return DefaultConfig(), nil
		}

		return nil, err
	}

	config := &Config{}

	err = json.Unmarshal(raw, config)
	if err != nil {
		return nil, err
	}

	config.process()

	return config, nil
}
