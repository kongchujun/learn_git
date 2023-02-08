package myreflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFields(t *testing.T) {
	//in order to test func:IterateFields
	type User struct {
		Name string
		age  int
	}
	testCase := []struct {
		name    string
		entity  any
		wantErr error
		wantRes map[string]any
	}{
		{
			name: "user",
			entity: User{
				Name: "tom",
				age:  18,
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
