package ormselect

import (
	"context"
	"database/sql"

	// errs "learn_git/orm/ormselect/internal/err"
	"strings"
)

// gis
// s *Selector ormselect.Querier
// s *Selector ormselect.QueryBuilder
// s *Selector ormselect.Excecutor
type Selector[T any] struct {
	table string
	where []Predicat
	buidler[T]
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
	var err error
	s.model, err = parseModel(new(T))
	if err != nil {
		return nil, err
	}
	sb.WriteString("SELECT * FROM ")

	if s.table == "" {

		sb.WriteByte('`')
		sb.WriteString(s.model.tableName)
		sb.WriteByte('`')
	} else {
		//``this logo handled by user
		sb.WriteString(s.table)
	}

	if len(s.where) > 0 {
		sb.WriteString(" WHERE ")
		p := s.where[0]
		for i := 1; i < len(s.where); i++ {
			p = p.And(s.where[i])
		}
		if err := s.buildExpressionBuilder(p); err != nil {
			return nil, err
		}
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

func (s *Selector[T]) Where(ps ...Predicat) *Selector[T] {
	s.where = ps
	return s
}
