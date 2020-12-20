package algo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2020/algo"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	data := []int{1, 2, 3, 4}
	found, acc := algo.SubsetSum(data, 9)
	assert.True(t, found)
	assert.Equal(t, []int{4, 3, 2}, acc)
}
