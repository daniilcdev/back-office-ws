package config

import (
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DbAdapter    string
	DbConnection string
}

type AppConfig struct {
	dbConfig
}

func LoadEnv() *AppConfig {
	godotenv.Load("./env/private/database.env")

	cfg := AppConfig{}

	cfg.DbAdapter = os.Getenv("DB_DRIVER")
	cfg.DbConnection = os.Getenv("DB_CONN")

	return &cfg
}
