package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func (c *Config) GetDBConnString() string {
	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable pool_min_conns=5 pool_max_conns=50", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	log.Info().Msgf("Connection string: %s", s)
	return s
}

func (c *Config) NewDBConnection(ctx context.Context) (*pgxpool.Pool, error) {

	dbPool, err := pgxpool.New(context.Background(), c.GetDBConnString())

	if err != nil {
		log.Error().Err(err).Msg("Could not connect to database")
		return nil, err
	}
	err = dbPool.Ping(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Could not ping database")
		return nil, err
	}

	return dbPool, nil
}
