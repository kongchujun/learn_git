package myreflect

import (
	"learn_git/orm/myreflect/types"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateFunc(t *testing.T) {
	testCase := []struct {
		name   string
		entity any

		wantRes map[string]FuncInfo
		wantErr error
	}{
		{
			name:   "struct",
			entity: types.NewUser("tom", 18),
			wantRes: map[string]FuncInfo{
				"GetAge": {
					Name:        "GetAge",
					InputTypes:  []reflect.Type{reflect.TypeOf(types.User{})},
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
				// "ChangeName": {
				// 	Name:       "ChangeName",
				// 	InputTypes: []reflect.Type{reflect.TypeOf("")},
				// 	//OutputTypes: []reflect.Type{reflect.TypeOf(0)},
				// 	//sResult: []any{18},
				// },
			},
		},

		{
			name:   "pointer",
			entity: types.NewUserPr("tom", 18),
			wantRes: map[string]FuncInfo{
				"GetAge": {
					Name:        "GetAge",
					InputTypes:  []reflect.Type{reflect.TypeOf(&types.User{})},
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
				"ChangeName": {
					Name:        "ChangeName",
					InputTypes:  []reflect.Type{reflect.TypeOf(&types.User{}), reflect.TypeOf("")},
					OutputTypes: []reflect.Type{},
					Result:      []any{},
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFunc(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
