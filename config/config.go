package config

import (
	"os"
)

// DbConfig is a structure for storage db configuration
type DbConfig struct {
	Database         string
	Username         string
	Password         string
	PostgresPassword string
}

// ServerConfig is a structure for storage server configuration
type ServerConfig struct {
	Host string
	Port string
	//Path string
}

type Config struct {
	ServerConfig ServerConfig
	DbConfig     DbConfig
}

// LoadConfiguration loads configuration from env variables
func LoadConfiguration() *Config {
	return &Config{
		ServerConfig: ServerConfig{Host: getEnv("HOST", ""), Port: getEnv("PORT", "")},
		DbConfig: DbConfig{
			Database:         getEnv("POSTGRESQL_DATABASE", ""),
			Username:         getEnv("POSTGRESQL_USERNAME", ""),
			Password:         getEnv("POSTGRESQL_PASSWORD", ""),
			PostgresPassword: getEnv("POSTGRESQL_POSTGRES_PASSWORD", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
