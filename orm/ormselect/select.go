package ormselect

import (
	"context"
	"database/sql"
	"fmt"
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
	where []Predicat
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

func (s *Selector[T]) buildExpressionBuilder(expr Expression) error {
	switch exp := expr.(type) {
	case nil:
	case Predicat:
		_, ok := exp.left.(Predicat)
		if ok {
			s.sb.WriteByte('(')
		}
		if err := s.buildExpressionBuilder(exp.left); err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}
		s.sb.WriteByte(' ')
		s.sb.WriteString(exp.op.String())
		s.sb.WriteByte(' ')

		_, ok = exp.right.(Predicat)
		if ok {
			s.sb.WriteByte('(')
		}
		if err := s.buildExpressionBuilder(exp.right); err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}
	case Column:
		s.sb.WriteByte('`')
		s.sb.WriteString(exp.name)
		s.sb.WriteByte('`')
	case value:
		s.sb.WriteString("?")
		s.args = append(s.args, exp.val)
	default:
		return fmt.Errorf("orm: 不支持模块 %v", expr)
	}
	return nil
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
