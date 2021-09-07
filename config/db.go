package config

import "fmt"

type dbConfig struct {
	Driver   string
	Dialect  string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

func (dc *dbConfig) GetDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		dc.Host,
		dc.Port,
		dc.User,
		dc.DbName,
		dc.SslMode,
		dc.Password,
	)
}

func getDbConfig() dbConfig {
	return dbConfig{
		Driver:   getEnv("DB_DRIVER", "pgx"),
		Dialect:  getEnv("DB_DIALECT", "postgres"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DbName:   getEnv("DB_NAME", "location"),
		SslMode:  getEnv("DB_SSL_MODE", "disable"),
	}
}
