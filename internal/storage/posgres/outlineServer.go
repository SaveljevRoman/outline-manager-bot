package posgres

import "time"

type OutlineServer struct {
	Id         int64     `db:"id"`
	ChatId     int64     `db:"chat_id"`
	ApiUrl     string    `db:"api_url"`
	CertSha256 string    `db:"cert_sha_256"`
	CreatedAt  time.Time `db:"created_at"`
}
