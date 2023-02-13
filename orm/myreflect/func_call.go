package myreflect

import "reflect"

type FuncInfo struct {
	Name        string
	InputTypes  []reflect.Type
	OutputTypes []reflect.Type
	Result      []any
}

func IterateFunc(entity any) (map[string]FuncInfo, error) {
	typ := reflect.TypeOf(entity)
	numMethod := typ.NumMethod()
	res := make(map[string]FuncInfo, numMethod)
	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		fn := method.Func // 从这里可以获取方法的反射， 从而利用其类型遍历
		//入参类型遍历
		numIn := fn.Type().NumIn()
		input := make([]reflect.Type, 0, numIn)
		inputValues := make([]reflect.Value, 0, numIn)

		inputValues = append(inputValues, reflect.ValueOf(entity))
		input = append(input, reflect.TypeOf(entity))
		for j := 1; j < numIn; j++ {
			fnInType := fn.Type().In(j)

			input = append(input, fnInType)
			inputValues = append(inputValues, reflect.Zero(fnInType))
		}
		//输出类型遍历
		numOut := fn.Type().NumOut()
		output := make([]reflect.Type, 0, numOut)
		// outputValues :=make([]reflect.Type, 0, numOut)
		for j := 0; j < numOut; j++ {
			fnOutType := fn.Type().Out(j)
			output = append(output, fnOutType)
			// outputValues = append(outputValues, reflect.Zero(fnOutType))
		}
		//调用方法
		resValues := fn.Call(inputValues)
		result := make([]any, 0, len(resValues))
		for _, v := range resValues {
			result = append(result, v.Interface())
		}
		res[method.Name] = FuncInfo{
			Name:        method.Name,
			InputTypes:  input,
			OutputTypes: output,
			Result:      result,
		}
	}
	return res, nil
}
