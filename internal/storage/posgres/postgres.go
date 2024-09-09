package posgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresClient struct {
	db *sqlx.DB
}

func NewPostgresClient(
	ctx context.Context,
	host, port, user, pass, dbName, sslMode string,
	maxOpenConns, maxIdleConns int,
) *PostgresClient {
	addr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, dbName, sslMode,
	)

	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return nil
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return &PostgresClient{db: db}
}
