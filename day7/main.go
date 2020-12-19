package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	emptyBag  = "no other"
	separator = "contain"
)

type bagsMapType map[string][]bagsType

type bagsType struct {
	bag      string
	bagCount int
}

func loadBags(data string) bagsMapType {
	lines := strings.Split(data, "\n")

	bagger := map[string][]bagsType{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.ReplaceAll(line, ".", "")

		parts := strings.Split(line, separator)
		bagsInput := strings.TrimSpace(parts[1][:len(parts[1])-1])
		bagsInputItems := strings.Split(bagsInput, ", ")

		if len(bagsInputItems) == 1 && bagsInputItems[0] == emptyBag {
			bagger[parts[0]] = nil
			continue
		}

		bags := make([]bagsType, 0)
		for _, bagInput := range bagsInputItems {
			bagInput := strings.TrimSpace(bagInput)
			bagArg := strings.Index(bagInput, " ") // first space
			bagCount, _ := strconv.Atoi(bagInput[0:bagArg])
			bagType := strings.TrimSpace(bagInput[bagArg+1:])
			bag := bagsType{bagCount: bagCount, bag: bagType}
			bags = append(bags, bag)
		}

		bagger[strings.TrimSpace(parts[0])] = bags
	}

	return bagger
}

func find(needs string, bags bagsMapType, acc map[string]int) int {
	for key, val := range bags {
		for _, bag := range val {
			if bag.bag == needs {
				acc[key] += 1
				find(key, bags, acc)
			}
		}
	}

	return len(acc)
}

func count(start string, bags bagsMapType) int {
	acc := 0

	containedBags := bags[start]
	for _, bag := range containedBags {
		acc += bag.bagCount
		acc += count(bag.bag, bags) * bag.bagCount
	}

	return acc
}

func main() {
	data, _ := ioutil.ReadFile("./day7/input-test2.txt")
	bags := loadBags(string(data))
	needs := "shiny gold"

	acc := map[string]int{}
	res1 := find(needs, bags, acc)
	fmt.Printf("Found %d for %s\n", res1, needs)

	res2 := count(needs, bags)
	fmt.Printf("Found %d for %s\n", res2, needs)
}
