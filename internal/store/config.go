package store

import (
	"os"
)

type Config struct {
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBName: getEnv("DB_NAME", "politics.sqlite"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
