package myunsafe3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnsafeAccessor_Field(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	u := &User{Name: "tom", Age: 29}
	a := NewUnsafeAccessor(u)
	val, err := a.Field("Age")
	require.NoError(t, err)
	assert.Equal(t, 29, val)

	err = a.SetField("Age", 33)
	require.NoError(t, err)

}
