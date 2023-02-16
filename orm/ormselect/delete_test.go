package ormselect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testCase := []struct {
		name      string
		builder   QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name:    "select with name",
			builder: &Deleter[TestModel]{},
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "select with from",
			builder: (&Deleter[TestModel]{}).From("`kong`"),
			wantQuery: &Query{
				SQL:  "DELETE FROM `kong`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "delete no name",
			builder: (&Deleter[TestModel]{}).From(""),
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "delete db from",
			builder: (&Deleter[TestModel]{}).From("`metrics`.`kpitable`"),
			wantQuery: &Query{
				SQL:  "DELETE FROM `metrics`.`kpitable`;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "delete where eq",
			builder: (&Deleter[TestModel]{}).Where(C("Id").Eq(2)),
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model` WHERE `id` = ?;",
				Args: []any{2},
			},
			wantErr: nil,
		},
		{
			name:    "delete where not",
			builder: (&Deleter[TestModel]{}).Where(Not(C("Id").Eq(2))),
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model` WHERE  NOT (`id` = ?);",
				Args: []any{2},
			},
			wantErr: nil,
		},
		{
			name:    "delete where and",
			builder: (&Deleter[TestModel]{}).Where(C("Id").Eq(2).And(C("FirstName").Eq("tom"))),
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model` WHERE (`id` = ?) AND (`first_name` = ?);",
				Args: []any{2, "tom"},
			},
			wantErr: nil,
		},
		{
			name:    "delete where or",
			builder: (&Deleter[TestModel]{}).Where(C("Id").Eq(2).Or(C("FirstName").Eq("tom"))),
			wantQuery: &Query{
				SQL:  "DELETE FROM `test_model` WHERE (`id` = ?) OR (`first_name` = ?);",
				Args: []any{2, "tom"},
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
