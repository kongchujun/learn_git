package myreflect

import "reflect"

func IterateArrayOrSlice(entity any) ([]any, error) {
	val := reflect.ValueOf(entity)
	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		res = append(res, ele.Interface())
	}
	return res, nil
}

func IterateMap(entity any) ([]any, []any, error) {
	val := reflect.ValueOf(entity)
	keys := val.MapKeys()
	res := make([]any, 0, val.Len())
	resKeys := make([]any, 0, val.Len())
	for _, key := range keys {
		v := val.MapIndex(key)
		res = append(res, v.Interface())
		resKeys = append(resKeys, key.Interface())
	}
	return resKeys, res, nil
}
