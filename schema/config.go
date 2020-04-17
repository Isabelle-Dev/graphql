package schema

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is the the JSON config structure
type Config struct {
	Database PostgresConfig `json:"database"`
}

// PostgresConfig represents metadata required to start and maintain postgres
// database connection
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// loadConfig will load the postgres configuration file
func loadConfig() (*Config, error) {
	f, err := os.Open(".config")
	if err != nil {
		return nil, fmt.Errorf("loadConfig(): could not find .config file")
	}
	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("loadConfig(): could not decode json .config file")
	}
	return &cfg, nil
}
