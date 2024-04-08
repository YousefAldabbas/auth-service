package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"

	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	DB_ENGINE   string
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
}

func (c Config) New() (*Config, error) {
	wd, err := os.Getwd()

	if err != nil {
		log.Error().Err(err).Msg("Unable to read main dir")
		return nil, errors.New("unable to read main dir")
	}
	godotenv.Load(filepath.Join(wd, ".env"))

	env := os.Getenv("ENV")

	godotenv.Load(filepath.Join(wd, "profiles", env+".env"))
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return &Config{
		Environment: env,
		DB_ENGINE:   os.Getenv("DB_ENGINE"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      port,
		DBUser:      os.Getenv("DB_USERNAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}, nil
}

