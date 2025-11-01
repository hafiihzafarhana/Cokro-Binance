package arraydata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSplitAndTrim_success(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{
			name:     "normal case",
			input:    "a, b ,c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "no spaces",
			input:    "apple,banana,orange",
			sep:      ",",
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "empty string",
			input:    "",
			sep:      ",",
			expected: []string{""},
		},
		{
			name:     "different separator",
			input:    "one| two |three",
			sep:      "|",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "single element",
			input:    "  hello  ",
			sep:      ",",
			expected: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitAndTrim(tt.input, tt.sep)
			assert.Equal(t, tt.expected, result)
		})
	}
}