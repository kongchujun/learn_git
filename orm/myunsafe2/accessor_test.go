package myunsafe2

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
	u := &User{Name: "tom", Age: 18}
	a := NewUnsafeAcessor(u)
	val, err := a.Field("Age")
	require.NoError(t, err)
	assert.Equal(t, 18, val)

	err = a.SetField("Age", 32)
	require.NoError(t, err)
}
