package ormselect

type op string

const (
	opEq  op = "="
	opNot op = "NOT"
	opAnd op = "AND"
	opOr  op = "OR"
)

func (p op) String() string {
	return string(p)
}

type Expression interface {
	expr()
}
type Predicat struct {
	left  Expression
	op    op
	right Expression
}

// C("id").Eq(12)
func (c Column) Eq(arg any) Predicat {
	return Predicat{
		left:  c,
		op:    opEq,
		right: value{val: arg},
	}
}

// NOT(C("id").Eq(12))
func Not(p Predicat) Predicat {
	return Predicat{
		op:    opNot,
		right: p,
	}
}

// C("id").Eq(12).And(C("name").Eq("tom"))
func (left Predicat) And(right Predicat) Predicat {
	return Predicat{
		left:  left,
		op:    opAnd,
		right: right,
	}
}

// C("id").Eq(12).And(C("name").Eq("tom"))
func (left Predicat) Or(right Predicat) Predicat {
	return Predicat{
		left:  left,
		op:    opOr,
		right: right,
	}
}

func (p Predicat) expr() {}

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}

type value struct {
	val any
}

func (p value) expr()  {}
func (p Column) expr() {}
