package ormselect

import (
	"testing"

	errs "learn_git/orm/ormselect/internal/err"

	"github.com/stretchr/testify/assert"
)

// type TestModel struct {
// 	Id        int64
// 	FirstName string
// 	Age       int8
// 	LastName  *sql.NullString
// }

func Test_parseModel(t *testing.T) {

	testCase := []struct {
		name      string
		entity    any
		wantModel *model
		wantErr   error
	}{
		{
			name:    "test model",
			entity:  TestModel{},
			wantErr: errs.ErrPointerOnly,
		},
		{
			name:   "stuct pointer",
			entity: &TestModel{},
			wantModel: &model{
				tableName: "test_model",
				fields: map[string]*field{
					"Id": {
						colName: "id",
					},
					"FirstName": {
						colName: "first_name",
					},
					"LastName": {
						colName: "last_name",
					},
					"Age": {
						colName: "age",
					},
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			m, err := parseModel(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.wantModel, m)
		})
	}
}
