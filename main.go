package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		name string `orm:"123345"`
	}

	typ := reflect.TypeOf(S{})
	numberField := typ.NumField()
	for i := 0; i < numberField; i++ {
		fb := typ.Field(i)
		if value, ok := fb.Tag.Lookup("orm"); ok {
			fmt.Println(value)
		}
	}
}
