package myreflect

import (
	"errors"
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
		{
			name: "user",
			entity: &User{
				Name: "tom",
				age:  18,
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},
		{
			name:    "entity is nil",
			entity:  nil,
			wantErr: errors.New("入参不能为nil"),
		},
		{
			name: "multi pointer",
			entity: func() **User {
				res := &User{
					Name: "tom",
					age:  18,
				}
				return &res
			}(),
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},
		{
			name:    "base type",
			entity:  23,
			wantErr: errors.New("入参需指向结构体"),
		},
		{
			name:    "no nil typevalue",
			entity:  (*User)(nil),
			wantErr: errors.New("其值不能为nil指向"),
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

func TestSetField(t *testing.T) {
	type User struct {
		Name string
		age  int
	}
	testCase := []struct {
		name   string
		entity any
		field  string
		newval any

		wantEntity any

		wantErr error
	}{
		{
			name: "pointer",
			entity: &User{
				Name: "tom",
			},
			field:  "Name",
			newval: "jeck",

			wantEntity: &User{
				Name: "jeck",
			},
			wantErr: nil,
		},

		{
			name: "struct",
			entity: User{
				Name: "tom",
			},
			field:   "Name",
			newval:  "jeck",
			wantErr: errors.New("不能设置"),
		},
		{
			name: "private ",
			entity: User{
				age: 18,
			},
			field:   "age",
			newval:  19,
			wantErr: errors.New("不能设置"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newval)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.entity, tc.wantEntity)

		})
	}
}
