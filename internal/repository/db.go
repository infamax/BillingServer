package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connectString string) (*db, error) {
	pool, err := pgxpool.Connect(ctx, connectString)

	if err != nil {
		return nil, err
	}

	return &db{
		pool: pool,
	}, nil
}