package storage

import (
	"outline-manager-bot/config"
	"outline-manager-bot/internal/storage/posgres"
)

type Storage struct {
	PgClient *posgres.PostgresClient
}

func NewStorage(pgConf *config.PostgresConfig) *Storage {
	pgClient := posgres.NewPostgresClient(pgConf.Host, pgConf.Port, pgConf.User, pgConf.Pass, pgConf.DbName, pgConf.SslMode)

	return &Storage{
		PgClient: pgClient,
	}
}
