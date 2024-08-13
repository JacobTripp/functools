package functools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	testCases := map[string]struct {
		fn       func(int) int
		input    []int
		expected []int
	}{
		"times two": {
			fn:       func(n int) int { return n * 2 },
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6, 8, 10, 12},
		},
	}
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			out := Map(testCase.input, testCase.fn)
			assert.Equal(t, out, testCase.expected)
		})
	}
}
