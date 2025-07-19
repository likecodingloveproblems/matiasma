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
	queries *models.Queries
}

func (p PostgresSession) LoadSession(ctx context.Context) ([]byte, error) {
	var sessionData []byte
	session, err := p.queries.GetLatestSession(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return session.SessionData, nil
}

func (p PostgresSession) StoreSession(ctx context.Context, data []byte) error {
	return p.queries.CreateSession(ctx, models.CreateSessionParams{
		SessionData: data,
		CreatedAt:   time.Now().UTC(),
	})
}

var _ telegram.SessionStorage = PostgresSession{}
