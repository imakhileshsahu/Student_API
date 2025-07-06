// Package config provides configuration structures and utilities for loading
// application settings from environment variables or configuration files.
//
// The Config struct defines the main configuration for the application,
// including environment, HTTP server settings, and database connection details.
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
	DBUser     string     `yaml:"db_user"`
	DBPassword string     `yaml:"db_password"`
	DBHost     string     `yaml:"db_host"`
	DBPort     string     `yaml:"db_port"`
	DBName     string     `yaml:"db_name"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}

func MustLoad(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
