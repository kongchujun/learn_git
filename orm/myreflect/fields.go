package myreflect

import (
	"reflect"
)

func IterateFields(entity any) (map[string]any, error) {
	//1. 用type value方法提取类型和值对象
	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)
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
