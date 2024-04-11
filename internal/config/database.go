package config

import (
	"context"
	"fmt"
	"database/sql"
	"github.com/rs/zerolog/log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (c *Config) GetDBConnString() string {
	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	log.Info().Msgf("Connection string: %s", s)
	return s
}

func (c *Config) NewDBConnection(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", c.GetDBConnString())
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to database")
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)

	err = db.PingContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Could not ping database")
		return nil, err
	}

	return db, nil
}
