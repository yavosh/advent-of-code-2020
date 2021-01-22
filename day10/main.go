package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/advent"
	"sort"
)

func count(adapters []int) int {
	builtinRating := adapters[len(adapters)-1] + 3

	differences := map[int]int{}
	adapterRating := 0

	for _, adapter := range adapters {

		diff := adapter - adapterRating
		//fmt.Printf("adapter: %d adapterRating:%d diff:%d\n", adapter, adapterRating, diff)

		if diff > 3 {
			panic("error too much diff")
		}

		adapterRating = adapter
		differences[diff]++
	}

	// last diff
	differences[builtinRating-adapterRating]++

	fmt.Printf("differences: %v\n", differences)
	return differences[1] * differences[3]
}

func distinct3(adapters []int) int {

	builtinRating := adapters[len(adapters)-1] + 3

	// add 0 add builtin an adapter
	//adapters = append([]int{0}, adapters...)
	adapters = append(adapters, builtinRating)

	pending := make([][]int, 0)
	pending = append(pending, []int{adapters[0]})
	done := make([][]int, 0)

	for len(pending) > 0 {
		current := pending[0]
		currentLast := current[len(current)-1]
		pos := sort.SearchInts(adapters, currentLast)

		rest := adapters[pos+1:]
		pending = pending[1:]

		if currentLast == builtinRating {
			done = append(done, current)
			continue
		}

		//fmt.Printf("curr: %v pending:%v rest:%v\n", current, pending, rest)
		for _, v := range rest {
			diff := v - currentLast
			if diff <= 3 {
				//fmt.Printf("diff: %d %d %d %v\n", diff, v, currentLast, current)
				active := append(current, v)
				pending = append(pending, active)
			}
		}
	}

	for _, d := range done {
		fmt.Printf("> %v\n", d)
	}

	return len(done)
}

func distinct(adapters []int) int {
	builtinRating := adapters[len(adapters)-1] + 3
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, builtinRating)

	adapterMap := make(map[int]int)
	for _, v := range adapters {
		adapterMap[v] = 0
	}

	adapterMap[0] = 1
	for _, current := range adapters {
		if _, ok := adapterMap[current+1]; ok {
			adapterMap[current+1] += adapterMap[current]
		}

		if _, ok := adapterMap[current+2]; ok {
			adapterMap[current+2] += adapterMap[current]
		}

		if _, ok := adapterMap[current+3]; ok {
			adapterMap[current+3] += adapterMap[current]
		}
	}

	fmt.Printf("res: %v\n", adapterMap)
	return adapterMap[builtinRating]
}

func main() {
	data := advent.LoadInt("./day10/input-test.txt")
	sort.Ints(data)
	fmt.Printf("data:%v\n", data)
	answer := count(data)
	fmt.Printf("answer: %v\n", answer)
	distinctWays := distinct(data)
	fmt.Printf("distinctWays: %v\n", distinctWays)
}
