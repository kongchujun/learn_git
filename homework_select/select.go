package orm

import (
	"learn_git/homework_select/model"
	"strings"
)

// Selector 用于构造 SELECT 语句
type Selector[T any] struct {
	sb      strings.Builder
	alias   []string
	args    []any
	table   string
	where   []Predicate
	having  []Predicate
	model   *model.Model
	db      *DB
	columns []Selectable
	groupBy []Column
	orderBy []OrderBy
	offset  int
	limit   int
}

func (s *Selector[T]) Select(cols ...Selectable) *Selector[T] {
	s.columns = cols
	return s
}

// From 指定表名，如果是空字符串，那么将会使用默认表名
func (s *Selector[T]) From(tbl string) *Selector[T] {
	s.table = tbl
	return s
}

func (s *Selector[T]) Build() (*Query, error) {
	panic("implement me")
}

// Where 用于构造 WHERE 查询条件。如果 ps 长度为 0，那么不会构造 WHERE 部分
func (s *Selector[T]) Where(ps ...Predicate) *Selector[T] {
	s.where = ps
	return s
}

// GroupBy 设置 group by 子句
func (s *Selector[T]) GroupBy(cols ...Column) *Selector[T] {
	panic("implement me")
}

func (s *Selector[T]) Having(ps ...Predicate) *Selector[T] {
	panic("implement me")
}

func (s *Selector[T]) Offset(offset int) *Selector[T] {
	panic("implement me")
}

func (s *Selector[T]) Limit(limit int) *Selector[T] {
	panic("implement me")
}

func (s *Selector[T]) OrderBy(orderBys ...OrderBy) *Selector[T] {
	panic("implement me")
}

func NewSelector[T any](db *DB) *Selector[T] {
	return &Selector[T]{
		db: db,
	}
}

type Selectable interface {
	selectable()
}

type OrderBy struct {
}

func Asc(col string) OrderBy {
	panic("implement me")
}

func Desc(col string) OrderBy {
	panic("implement me")
}
