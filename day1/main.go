package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./day1/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	itemsCount := len(lines)

	fmt.Printf("len: %d\n", len(lines))
	for i := 0; i < itemsCount; i++ {
		for j := 0; j < itemsCount; j++ {
			for k := 0; k < itemsCount; k++ {

				a, _ := strconv.Atoi(lines[i])
				b, _ := strconv.Atoi(lines[j])
				c, _ := strconv.Atoi(lines[k])
				sum := a + b + c

				if sum == 2020 {
					fmt.Printf("Found %s %s %s a=%d\n", lines[i], lines[j], lines[k], a*b*c)

				}
			}
		}
	}

}
