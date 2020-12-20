package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/advent"
	"sort"
)

func solve(data []int, depth int) ([]int, []int) {
	has := make([]int, 0)
	nope := make([]int, 0)
	for i := depth; i < len(data); i++ {
		// start from depth
		target := data[i]
		preamble := data[i-depth : i]
		hasSum := 0
		for xi, x := range preamble {
			for yi, y := range preamble {
				if xi == yi {
					continue
				}

				if x+y == target {
					hasSum++
					has = append(has, target)
				}
			}
		}

		if hasSum == 0 {
			nope = append(nope, target)
		}
	}

	return has, nope
}

func findSum(data []int, target int) []int {

	res := make([]int, 0)
	pos := 0
	sum := 0

	for pos < len(data) {
		for i := pos; i < len(data); i++ {
			sum += data[i]
			res = append(res, data[i])
			if sum == target {
				return res
			}

			if sum > target {
				break
			}
		}

		pos++
		sum = 0
		res = make([]int, 0)
	}

	return []int{}
}

func main() {
	//xmasCipher := advent.LoadInt("./day9/input-test.txt")
	//depth := 5

	xmasCipher := advent.LoadInt("./day9/input.txt")
	depth := 25
	_, nope := solve(xmasCipher, depth)
	sumOf := nope[0]
	res := findSum(xmasCipher, sumOf)
	sort.Ints(res)
	fmt.Printf("sumof:%d res:%v\n", sumOf, res)
	fmt.Printf("weakness:%v\n", res[0]+res[len(res)-1])
}
