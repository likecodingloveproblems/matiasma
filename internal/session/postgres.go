package session

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/likecodingloveproblems/matiasma/internal/db"
	"github.com/likecodingloveproblems/matiasma/internal/models"
	"go.uber.org/zap"
)

import (
	"database/sql"
	"time"
)

type PostgresSessionStorage struct {
	PhoneNumber string
	Queries     *models.Queries
	Logger      *zap.Logger
	conn        *sql.DB
}

func New(ctx context.Context, phoneNumber string, logger *zap.Logger) *PostgresSessionStorage {
	conn := db.New(ctx, logger)
	return &PostgresSessionStorage{
		PhoneNumber: phoneNumber,
		Queries:     models.New(conn),
		Logger:      logger,
		conn:        conn,
	}
}

func (p PostgresSessionStorage) Close() error {
	return p.conn.Close()
}

func (p PostgresSessionStorage) LoadSession(ctx context.Context) ([]byte, error) {
	p.Logger.Info("Load Session ", zap.String("phone_number", p.PhoneNumber))
	session, err := p.Queries.GetUserLatestSession(ctx, p.PhoneNumber)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return session.SessionData, nil
}

func (p PostgresSessionStorage) StoreSession(ctx context.Context, data []byte) error {
	p.Logger.Info("Store Session ", zap.String("phone_number", p.PhoneNumber))
	return p.Queries.UpsertSession(ctx, models.UpsertSessionParams{
		SessionData: data,
		PhoneNumber: p.PhoneNumber,
		CreatedAt:   time.Now().UTC(),
	})
}

var _ telegram.SessionStorage = PostgresSessionStorage{}
