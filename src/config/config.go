package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	Server Server `json:"server"`
}

type Server struct {
	Address      string        `json:"address"`
	WriteTimeout time.Duration `json:"write_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
}

func Load() (*Configuration, error) {
	data, err := os.ReadFile("settings.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read settings.json: %w", err)
	}

	var config Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse settings.json: %w", err)
	}

	return &config, nil
}
