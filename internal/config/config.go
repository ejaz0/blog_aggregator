package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	path := filepath.Join(home, configFileName)

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUsername = name
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, configFileName)

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 06444)
}
