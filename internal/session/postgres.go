package session

import (
	"context"
	"github.com/gotd/td/telegram"
)

type PostgresSession struct {
}

func (p PostgresSession) LoadSession(ctx context.Context) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresSession) StoreSession(ctx context.Context, data []byte) error {
	//TODO implement me
	panic("implement me")
}

var _ telegram.SessionStorage = PostgresSession{}
