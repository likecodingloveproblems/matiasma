package session

import (
	"context"
	"github.com/gotd/td/telegram"
)

import (
	"database/sql"
	"time"
	
	_ "github.com/lib/pq"
)

type PostgresSession struct {
	DB *sql.DB
}

func (p PostgresSession) LoadSession(ctx context.Context) ([]byte, error) {
	var sessionData []byte
	err := p.DB.QueryRowContext(ctx,
		`SELECT session_data FROM telegram_sessions 
		ORDER BY created_at DESC 
		LIMIT 1`).Scan(&sessionData)
		
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return sessionData, err
}

func (p PostgresSession) StoreSession(ctx context.Context, data []byte) error {
	_, err := p.DB.ExecContext(ctx,
		`INSERT INTO telegram_sessions (session_data, created_at)
		VALUES ($1, $2)`,
		data, time.Now().UTC())
	return err
}

var _ telegram.SessionStorage = PostgresSession{}
