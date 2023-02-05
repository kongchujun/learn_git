package ormselect

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectBuildFun(t *testing.T) {
	testCase := []struct {
		name      string
		builder   QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name:    "select with name",
			builder: &Selector[TestModel]{},
			wantQuery: &Query{
				SQL:  "SELECT * FROM `TestModel`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "select with from",
			builder: (&Selector[TestModel]{}).From("`kong`"),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `kong`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "select no name",
			builder: (&Selector[TestModel]{}).From(""),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `TestModel`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "select db from",
			builder: (&Selector[TestModel]{}).From("`metrics`.`kpitable`"),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `metrics`.`kpitable`;",
				Args: nil,
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.builder.Build()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
