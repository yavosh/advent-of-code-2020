package valid_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/valid"
	"testing"
)

func TestOneOf(t *testing.T) {

	testCases := []struct {
		name     string
		expected []string
		val      string
		result   bool
	}{
		{name: "valid", expected: []string{"one", "two"}, val: "one", result: true},
		{name: "valid", expected: []string{"one", "two"}, val: "One", result: false},
		{name: "valid", expected: []string{"one", "two"}, val: "three", result: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checker := valid.NewOneOfChecker(tt.expected...)
			if !tt.result {
				assert.False(t, checker.Check(tt.val))
				return
			}
			assert.True(t, checker.Check(tt.val))
		})
	}

}
