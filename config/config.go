package config

import (
	"os"
	"strconv"
)

var instance *Config

type Config struct {
	Server serverConfig
	Db     dbConfig
}

func Get() *Config {
	if instance != nil {
		return instance
	}

	return &Config{
		Db:     getDbConfig(),
		Server: getServerConfig(),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
