package ormselect

import (
	"learn_git/orm/ormselect/internal/err"
	"strings"
)

type buidler[T any] struct {
	model *model
	sb    *strings.Builder
	args  []any
}

func (s *buidler[T]) buildExpressionBuilder(expr Expression) error {
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
		fd, ok := s.model.fields[exp.name]
		if !ok {
			return err.NewErrUnknowfield(exp.name)
		}

		s.sb.WriteByte('`')
		s.sb.WriteString(fd.colName)
		s.sb.WriteByte('`')
	case value:
		s.sb.WriteString("?")
		s.args = append(s.args, exp.val)
	default:
		return err.NewErrUnsupportType(expr)
	}
	return nil
}
