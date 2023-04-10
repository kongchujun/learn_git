package myunsafe3

import (
	"errors"
	"reflect"
	"unsafe"
)

type UnsafeAccessor struct {
	fields  map[string]FieldMeta
	address unsafe.Pointer
}

func NewUnsafeAccessor(entity any) *UnsafeAccessor {
	typ := reflect.TypeOf(entity)
	typ = typ.Elem()
	numField := typ.NumField()
	fields := make(map[string]FieldMeta, numField)
	for i := 0; i < numField; i++ {
		fd := typ.Field(i)
		fields[fd.Name] = FieldMeta{
			Offset: fd.Offset,
			typ:    fd.Type,
		}
	}
	val := reflect.ValueOf(entity)
	return &UnsafeAccessor{
		fields:  fields,
		address: val.UnsafePointer(),
	}
}

func (a *UnsafeAccessor) Field(name string) (any, error) {
	fm, ok := a.fields[name]
	if !ok {
		return nil, errors.New("非法类型")
	}
	fmAddress := unsafe.Pointer(uintptr(a.address) + fm.Offset)
	return reflect.NewAt(fm.typ, fmAddress).Elem().Interface(), nil
}

func (a *UnsafeAccessor) SetField(name string, val any) error {
	fm, ok := a.fields[name]
	if !ok {
		return errors.New("非法类型")
	}
	fmAddress := unsafe.Pointer(uintptr(a.address) + fm.Offset)
	reflect.NewAt(fm.typ, fmAddress).Elem().Set(reflect.ValueOf(val))
	return nil
}

type FieldMeta struct {
	Offset uintptr
	typ    reflect.Type
}
