package ormselect

import (
	"context"
	"database/sql"
)

type Querier[T any] interface {
	Get(ctx context.Context) (*T, error)
	GetMuti(ctx context.Context) ([]*T, error)
}

type Excecutor interface {
	Exec(ctx context.Context) (sql.Result, error)
}

type QueryBuilder interface {
	Build() (*Query, error)
}

type Query struct {
	SQL  string
	Args []any
}
