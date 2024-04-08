package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)


func (c *Config) GetDBConnString() string {
	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	log.Info().Msgf("Connection string: %s", s)
	return s
}


func (c *Config) NewDBConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, c.GetDBConnString())
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to database")
		return nil, err
	}
	err = conn.Ping(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Could not ping database")
		return nil, err
	}

	// defer conn.Close(ctx)
	return conn, nil
}
