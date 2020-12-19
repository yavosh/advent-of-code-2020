package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./day3/input.txt")
	lines := strings.Split(string(data), "\n")

	rows := len(lines)
	cols := len(lines[0])

	treeMap := make([][]string, rows)
	for i := 0; i < rows; i++ {
		treeMap[i] = make([]string, cols)
	}

	for row, line := range lines {
		for col := range line {
			treeMap[row][col] = string(line[col])
		}
	}

	acc := 1
	steps := []struct {
		right int
		down  int
	}{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	for _, step := range steps {
		trees := 0

		// col, row
		pos := struct {
			row int
			col int
		}{row: 0, col: 0}

		for true {
			pos.col = (pos.col + step.right) % cols
			pos.row = pos.row + step.down
			if pos.row >= rows {
				break
			}

			value := treeMap[pos.row][pos.col]
			if value == "#" {
				trees++
			}

		}

		fmt.Printf("Trees %v %d\n", step, trees)
		acc = acc * trees
	}

	fmt.Printf("Total %d\n", acc)
}
