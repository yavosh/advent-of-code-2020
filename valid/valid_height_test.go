package valid_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/valid"
	"testing"
)

func TestValidHeight(t *testing.T) {

	testCases := []struct {
		name   string
		val    string
		result bool
		mincm  int
		maxcm  int
		minin  int
		maxin  int
	}{
		{name: "valid-cm", val: "160cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: true},
		{name: "valid-cm", val: "150cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: true},
		{name: "valid-cm", val: "193cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: true},
		{name: "invalid-cm", val: "149cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: false},
		{name: "invalid-cm", val: "194cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: false},
		{name: "valid-in", val: "160cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: true},
		{name: "invalid-cm", val: "130cm", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: false},
		{name: "invalid-in", val: "70in", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: false},
		{name: "invalid-in", val: "80ain", mincm: 150, maxcm: 193, minin: 76, maxin: 159, result: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checker := valid.HeightChecker{
				MinCm:   tt.mincm,
				MaxCm:   tt.maxcm,
				MinInch: tt.minin,
				MaxInch: tt.maxin,
			}
			if !tt.result {
				assert.False(t, checker.Check(tt.val))
				return
			}
			assert.True(t, checker.Check(tt.val))
		})
	}
}
