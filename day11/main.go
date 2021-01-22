package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/advent"
	"strings"
)

const (
	empty    = "L"
	occupied = "#"
	space    = "."
)

type coordinate struct {
	r int // row
	c int // col
}

var (
	UL = coordinate{-1, -1}
	UR = coordinate{-1, 1}
	DR = coordinate{1, 1}
	DL = coordinate{1, -1}
	L  = coordinate{0, 1}
	R  = coordinate{0, -1}
	U  = coordinate{-1, 0}
	D  = coordinate{1, 0}

	allDirections = []coordinate{
		UL, UR, DL, DR,
		L, R, U, D,
	}
)

func checkSeatDirection(origin coordinate, dir coordinate, w, h int, m [][]string) string {

	pos := coordinate{}
	pos.r = origin.r + dir.r
	pos.c += origin.c + dir.c

	for {
		// out of bounds
		if pos.c < 0 || pos.c > w-1 {
			return empty
		}

		// out of bounds
		if pos.r < 0 || pos.r > h-1 {
			return empty
		}

		//fmt.Printf("pos %v\n", pos)
		//fmt.Printf("dir %v\n", dir)

		// step
		if m[pos.r][pos.c] != space {
			// if not space return first visible
			return m[pos.r][pos.c]
		}

		pos.r += dir.r
		pos.c += dir.c
	}
}

func visOccupiedCount(p coordinate, w, h int, m [][]string) int {
	occ := 0

	for _, dir := range allDirections {
		if checkSeatDirection(p, dir, w, h, m) == occupied {
			occ++
		}
	}

	return occ
}

func adjOccupiedCount(r, c int, w, h int, m [][]string) int {
	occ := 0

	// UL
	if r > 0 && c > 0 {
		if m[r-1][c-1] == "#" {
			occ++
		}
	}

	// UR
	if r > 0 && c < w-1 {
		if m[r-1][c+1] == "#" {
			occ++
		}
	}

	// DR
	if r < h-1 && c < w-1 {
		if m[r+1][c+1] == "#" {
			occ++
		}
	}

	// DL
	if r < h-1 && c > 0 {
		if m[r+1][c-1] == "#" {
			occ++
		}
	}

	// L
	if c > 0 {
		if m[r][c-1] == "#" {
			occ++
		}
	}

	// R
	if c < w-1 {
		if m[r][c+1] == "#" {
			occ++
		}
	}

	// U
	if r > 0 {
		if m[r-1][c] == "#" {
			occ++
		}
	}

	// D
	if r < h-1 {
		if m[r+1][c] == "#" {
			occ++
		}
	}

	return occ
}

func dimension(plan [][]string) (int, int) {
	width := len(plan[0])
	height := len(plan)
	return width, height
}

func printSeats(plan [][]string) {
	width, height := dimension(plan)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			//a := adjOccupiedCount(r, c, width, height, plan)
			//a := visOccupiedCount(coordinate{r, c}, width, height, plan)
			//fmt.Printf("(%d,%d)%s[%d] ", r, c, plan[r][c], a)
			fmt.Printf("%s", plan[r][c])
		}

		fmt.Println()
	}

	fmt.Println()
}

func planStep(plan [][]string) [][]string {
	width, height := dimension(plan)

	newPlan := make([][]string, height)
	for r := 0; r < height; r++ {
		newPlan[r] = make([]string, width)

		for c := 0; c < width; c++ {
			a := adjOccupiedCount(r, c, width, height, plan)

			if plan[r][c] == empty && a == 0 {
				newPlan[r][c] = occupied
				continue
			}

			if plan[r][c] == occupied && a >= 4 {
				newPlan[r][c] = empty
				continue
			}

			newPlan[r][c] = plan[r][c]
		}
	}

	return newPlan
}

func planStepVis(plan [][]string) [][]string {
	width, height := dimension(plan)

	newPlan := make([][]string, height)
	for r := 0; r < height; r++ {
		newPlan[r] = make([]string, width)

		for c := 0; c < width; c++ {
			a := visOccupiedCount(coordinate{r, c}, width, height, plan)

			if plan[r][c] == empty && a == 0 {
				newPlan[r][c] = occupied
				continue
			}

			if plan[r][c] == occupied && a >= 5 {
				newPlan[r][c] = empty
				continue
			}

			newPlan[r][c] = plan[r][c]
		}
	}

	return newPlan
}

func source(data [][]string) string {
	b := strings.Builder{}
	for _, row := range data {
		b.WriteString(strings.Join(row, ""))
	}
	return b.String()
}

func main() {
	input := advent.Load("./day11/input.txt")
	data := make([][]string, 0)

	for _, v := range input {
		data = append(data, strings.Split(v, ""))
	}

	resultOccupied := 0
	printSeats(data)
	lastPlanSource := source(data)
	steps := 0
	for {

		//data = planStep(data)
		data = planStepVis(data)

		printSeats(data)

		s := source(data)

		if s == lastPlanSource {
			width, height := dimension(data)

			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					if data[r][c] == occupied {
						resultOccupied++
					}
				}
			}

			break
		}

		lastPlanSource = s
		steps++
	}

	fmt.Println("Steps", steps)
	fmt.Println("Occupied", resultOccupied)
}
