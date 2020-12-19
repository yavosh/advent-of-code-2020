package valid_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/valid"
	"testing"
)

func TestValidColour(t *testing.T) {

	testCases := []struct {
		name   string
		val    string
		result bool
	}{
		{name: "valid", val: "#123456", result: true},
		{name: "valid-hex", val: "#0c0c0c", result: true},
		{name: "valid-only-hex", val: "#ffffff", result: true},
		{name: "no-hash", val: "123456", result: false},
		{name: "less-numbers", val: "#12346", result: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checker := valid.ColourChecker{}
			if !tt.result {
				assert.False(t, checker.Check(tt.val))
				return
			}
			assert.True(t, checker.Check(tt.val))
		})
	}
}
