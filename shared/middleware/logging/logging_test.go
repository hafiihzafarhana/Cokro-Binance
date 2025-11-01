package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging_ReturnsHandler(t *testing.T) {
	handler := Logging()
	assert.NotNil(t, handler)
}
