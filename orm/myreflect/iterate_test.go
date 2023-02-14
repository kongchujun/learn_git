package myreflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateArray(t *testing.T) {
	testCase := []struct {
		name    string
		entity  any
		wantRes []any
		wantErr error
	}{
		{
			name:    "array",
			entity:  [3]int{1, 2, 3},
			wantRes: []any{1, 2, 3},
		},
		{
			name:    "slice",
			entity:  []int{1, 2, 3},
			wantRes: []any{1, 2, 3},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateArrayOrSlice(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)

		})
	}
}

func TestIterateMap(t *testing.T) {
	testCase := []struct {
		name      string
		entity    any
		wantKey   []any
		wantValue []any
		wantErr   error
	}{
		{
			name: "map",
			entity: map[string]string{
				"name": "kong",
				"age":  "12",
			},
			wantKey:   []any{"name", "age"},
			wantValue: []any{"kong", "12"},
			wantErr:   nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			resKeys, res, err := IterateMap(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.wantKey, resKeys)
			assert.Equal(t, tc.wantValue, res)

		})
	}

}
