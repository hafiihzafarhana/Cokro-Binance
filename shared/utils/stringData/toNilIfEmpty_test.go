package stringdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNilIfEmpty_Success(t *testing.T) {
	str := func(s string) *string { return &s } // make pointer from literal string
	var inputNil *string
	assert.Nil(t, ToNilIfEmpty(inputNil))

	empty := ""
	assert.Nil(t, ToNilIfEmpty(&empty))

	value := "hello"
	result := ToNilIfEmpty(str(value))
	assert.NotNil(t, result)
	assert.Equal(t, value, *result)
}