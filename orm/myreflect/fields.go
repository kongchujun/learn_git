package myreflect

import (
	"errors"
	"reflect"
)

// 其目的在于遍历struct结构的属性
func IterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, errors.New("入参不能为nil")
	}
	//1. 用type value方法提取类型和值对象
	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)
	if val.IsZero() { // （*user)(nil)
		return nil, errors.New("其值不能为nil指向")
	}
	//1.1 增加如果入参是指针的情况, 不管多少层指针
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("入参需指向结构体")
	}
	//2. 提取数目， 方便用for循环
	numFields := typ.NumField()
	res := make(map[string]any, numFields)
	for i := 0; i < numFields; i++ {
		typeField := typ.Field(i)
		typeValue := val.Field(i)
		//3. 提取值之后， 判断是否是公有变量， 然后再赋值
		if typeField.IsExported() {
			res[typeField.Name] = typeValue.Interface()
		} else {
			res[typeField.Name] = reflect.Zero(typeField.Type).Interface()
		}
	}
	return res, nil
}

// 可以设置struct
func SetField(entity any, field string, newval any) error {
	val := reflect.ValueOf(entity)
	for val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}
	fieldVal := val.FieldByName(field)
	if !fieldVal.CanSet() {
		return errors.New("不能设置")
	}
	fieldVal.Set(reflect.ValueOf(newval))
	return nil
}
