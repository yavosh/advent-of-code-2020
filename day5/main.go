package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getValue(data string, upper rune, lower rune) (int, int) {
	start := 0
	end := (1 << len(data)) - 1
	remain := 0

	for _, r := range data {
		remain = (end - start) % 2
		delta := ((end - start) / 2) + remain

		if r == lower {
			end = end - delta
		}

		if r == upper {
			start = start + delta
		}
	}

	return start, end
}

func main() {
	data, _ := ioutil.ReadFile("./day5/input.txt")
	lines := strings.Split(string(data), "\n")

	maxSeatID := 0
	minSeatId := 128

	seatsTaken := map[int]bool{}

	for _, line := range lines {
		start, _ := getValue(line[0:7], 'B', 'F')
		startCol, _ := getValue(line[7:], 'R', 'L')
		seatId := start*8 + startCol

		if seatId > maxSeatID {
			maxSeatID = seatId
		}

		if seatId < minSeatId {
			minSeatId = seatId
		}

		seatsTaken[seatId] = true
	}

	mySeat := 0
	for i := minSeatId; i < maxSeatID; i++ {
		if !seatsTaken[i] {
			mySeat = i
		}
	}

	fmt.Printf("maxSeatID: %d\n", maxSeatID)
	fmt.Printf("minSeatId: %d\n", minSeatId)
	fmt.Printf("mySeat: %d\n", mySeat)
}
