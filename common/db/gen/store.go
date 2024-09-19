package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// Store provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func (store *SQLStore) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return store.db.QueryRow(ctx, query, args)
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
