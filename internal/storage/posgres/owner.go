package posgres

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Owner struct {
	Id             int64     `db:"id"`
	Name           string    `db:"name"`
	ChatId         int64     `db:"chat_id"`
	IsOutlineAdmin bool      `db:"is_outline_admin"`
	CreatedAt      time.Time `db:"created_at"`
}

func (pc *PostgresClient) InsertNewOwner(ctx context.Context, chatId int64, name string) (*Owner, error) {
	tx, err := pc.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction insert new owner: %w", err)
	}

	var owner Owner
	err = tx.QueryRowContext(ctx, `INSERT INTO owner (chat_id, name)
        VALUES (:name, :chatId)
        ON CONFLICT (chat_id) DO UPDATE SET name = EXCLUDED.name
        RETURNING id, name, chat_id, is_outline_admin, created_at`, chatId, name).Scan(&owner)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, fmt.Errorf("failed to rollbacke insert new owner: %w", err)
		}
		return nil, fmt.Errorf("failed to insert new owner: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &owner, nil
}
