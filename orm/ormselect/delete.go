package ormselect

import (
	"strings"
)

type Deleter[T any] struct {
	buidler[T]
	table string
	where []Predicat
}

func (d *Deleter[T]) Build() (*Query, error) {
	d.sb = &strings.Builder{}
	sb := d.sb
	var err error
	d.model, err = parseModel(new(T))
	if err != nil {
		return nil, err
	}
	sb.WriteString("DELETE FROM ")

	if d.table == "" {

		sb.WriteByte('`')
		sb.WriteString(d.model.tableName)
		sb.WriteByte('`')
	} else {
		//``this logo handled by user
		sb.WriteString(d.table)
	}

	if len(d.where) > 0 {
		sb.WriteString(" WHERE ")
		p := d.where[0]
		for i := 1; i < len(d.where); i++ {
			p = p.And(d.where[i])
		}
		if err := d.buildExpressionBuilder(p); err != nil {
			return nil, err
		}
	}

	sb.WriteString(";")
	return &Query{
		SQL:  sb.String(),
		Args: d.args,
	}, nil
}

// From accepts model definition
func (d *Deleter[T]) From(table string) *Deleter[T] {
	d.table = table
	return d
}

// Where accepts predicates
func (d *Deleter[T]) Where(predicates ...Predicat) *Deleter[T] {
	d.where = predicates
	return d
}
