package config

import (
    "os"
)

type Config struct {
    Port        string
    DatabaseURL string
}

func LoadConfig() *Config {
    return &Config{
        Port:        getEnv("PORT", ":8080"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost/dbname?sslmode=disable"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}