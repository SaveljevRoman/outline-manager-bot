package storage

import (
	"context"
	"outline-manager-bot/config"
	"outline-manager-bot/internal/storage/posgres"
)

type Storage struct {
	PgClient *posgres.PostgresClient
}

func NewStorage(ctx context.Context, pgConf *config.PostgresConfig) *Storage {
	pgClient := posgres.NewPostgresClient(
		ctx,
		pgConf.Host, pgConf.Port, pgConf.User, pgConf.Pass, pgConf.DbName, pgConf.SslMode,
		pgConf.MaxOpenConns, pgConf.MaxIdleConns,
	)

	return &Storage{
		PgClient: pgClient,
	}
}
