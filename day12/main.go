package main

import (
	"github.com/yavosh/advent-of-code-2020/advent"
	"sort"
)

func main() {
	data := advent.LoadInt("./day12/input-test.txt")
	sort.Ints(data)
}
