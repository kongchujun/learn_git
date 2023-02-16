package err

import (
	"errors"
	"fmt"
)

var (
	ErrPointerOnly = errors.New("orm: 只支持一级指针")
)

func NewErrUnsupportExpression(expr any) error {
	return fmt.Errorf("orm:不支持表达式 %v", expr)
}

func NewErrUnknowfield(name string) error {
	return fmt.Errorf("orm:未知字段不支持 %s", name)
}

func NewErrUnsupportType(expr any) error {
	return fmt.Errorf("orm:不支持表达式 %v", expr)
}
