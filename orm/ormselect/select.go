package ormselect

import (
	"context"
	"database/sql"
	"reflect"
	"strings"
)

// gis
// s *Selector ormselect.Querier
// s *Selector ormselect.QueryBuilder
// s *Selector ormselect.Excecutor
type Selector[T any] struct {
	table string
	sb    *strings.Builder
	args  []any
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Selector[T]) GetMuti(ctx context.Context) ([]*T, error) {
	panic("not implemented") // TODO: Implement
}
func (s *Selector[T]) Build() (*Query, error) {
	s.sb = &strings.Builder{}
	sb := s.sb
	sb.WriteString("SELECT * FROM ")

	if s.table == "" {
		var t T
		typ := reflect.TypeOf(t)
		sb.WriteByte('`')
		sb.WriteString(typ.Name())
		sb.WriteByte('`')
	} else {
		//``this logo handled by user
		sb.WriteString(s.table)
	}
	sb.WriteString(";")
	return &Query{
		SQL:  sb.String(),
		Args: s.args,
	}, nil
}
func (s *Selector[T]) Exec(ctx context.Context) (sql.Result, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Selector[T]) From(table string) *Selector[T] {
	s.table = table
	return s
}
