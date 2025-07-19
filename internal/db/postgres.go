package db

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"os"
)

func New(ctx context.Context, logger *zap.Logger) *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}
	return conn
}
