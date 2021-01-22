package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/advent"
)

func group(data []string) []map[string]int {
	var groups []map[string]int
	var group map[string]int

	group = make(map[string]int)
	for _, line := range data {
		if line == "" {
			groups = append(groups, group)
			group = make(map[string]int)
			continue
		}

		for _, answer := range line {
			group[string(answer)] += 1
		}

		group["_count"] += 1
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups
}

func main() {

	lines := advent.Load("./day6/input.txt")
	groups := group(lines)

	acc := 0
	accAll := 0
	for _, g := range groups {
		acc += len(g)
		for answer := range g {
			if answer != "_count" {
				if g[answer] == g["_count"] { // -1 to ignore the _count
					accAll++
				}
			}
		}
	}

	fmt.Printf("Answers: %d\n", acc)
	fmt.Printf("Answers all: %d\n", accAll)
}
