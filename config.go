package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Editor     string `json:"editor"`
	StorageDir string `json:"storage_dir"`
}

func (c *Config) process() {
	if c.Editor == "" {
		if runtime.GOOS == "windows" {
			c.Editor = "notepad"
		} else {
			c.Editor = "vi"
		}
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

	configDir = filepath.Join(configDir, appName)

	return configDir, nil
}

func expandHome(p string) string {
	if len(p) > 0 && p[0] == '~' {
		home, _ := os.UserHomeDir()
		suffix := p[1:]
		if len(suffix) > 0 && (suffix[0] == '/' || suffix[0] == '\\') {
			suffix = suffix[1:]
		}
		return filepath.Join(home, suffix)
	}
	return p
}

func LoadConfig() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	configFilePath := filepath.Join(configDir, "config.jsonc")

	raw, err := os.ReadFile(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return DefaultConfig(), nil
		}

		return nil, fmt.Errorf("failed to read config file (%s): %w", configFilePath, err)
	}

	config := &Config{}

	err = json.Unmarshal(raw, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	config.process()

	return config, nil
}
