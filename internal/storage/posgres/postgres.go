package posgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresClient struct {
	db *sqlx.DB
}

func NewPostgresClient(host, port, user, pass, dbName, sslMode string) *PostgresClient {
	addr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, dbName, sslMode,
	)

	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return nil
	}
	defer db.Close()

	return &PostgresClient{db: db}
}
