package algo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/algo"
	"testing"
)

func TestPowerSet(t *testing.T) {
	data := []int{1, 2, 3}
	res := algo.PowerSet(data)
	assert.Len(t, res, 8)
	expected := [][]int{
		{3, 2, 1}, {3, 2}, {3, 1}, {3},
		{2, 1}, {2},
		{1},
		{},
	}
	assert.Equal(t, expected, res)
}
