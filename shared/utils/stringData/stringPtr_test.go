package stringdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringPtr_Success(t *testing.T) {
	value := "hello"
	result := StringPtr(value)
	assert.NotNil(t, result)
	assert.Equal(t, value, *result)
}