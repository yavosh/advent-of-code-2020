package valid_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/valid"
	"testing"
)

func TestMixMax(t *testing.T) {

	testCases := []struct {
		name   string
		val    string
		min    int
		max    int
		result bool
	}{
		{name: "valid", val: "", min: 1980, max: 2000, result: false},
		{name: "valid", val: "1982", min: 1980, max: 2000, result: true},
		{name: "less-valid", val: "1972", min: 1980, max: 2000, result: false},
		{name: "less-more", val: "2021", min: 1980, max: 2020, result: false},
		{name: "exact-min", val: "1980", min: 1980, max: 2000, result: true},
		{name: "exact-max", val: "2000", min: 1980, max: 2000, result: true},
		{name: "non-numeric", val: "a2000", min: 1980, max: 2020, result: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checker := valid.MinMaxChecker{Min: tt.min, Max: tt.max}
			if !tt.result {
				assert.False(t, checker.Check(tt.val))
				return
			}
			assert.True(t, checker.Check(tt.val))
		})
	}
}
