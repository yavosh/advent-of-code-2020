package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/advent"
	"strconv"
)

type position struct {
	x int
	y int
}

type ship struct {
	heading int // degrees
	pos     position
}

type instruction struct {
	action string
	param  int
}

func (s *ship) Move(i instruction) {

	switch i.action {
	case "N":
		s.pos.y -= i.param
	case "S":
		s.pos.y += i.param
	case "E":
		s.pos.x += i.param
	case "W":
		s.pos.x -= i.param
	case "L":
		s.heading -= i.param
		if s.heading < 0 {
			s.heading += 360
		}
	case "R":
		s.heading += i.param
		if s.heading >= 360 {
			s.heading -= 360
		}
	case "F":
		// move according to heading
		switch s.heading {
		case 0: // north
			s.pos.y -= i.param
		case 90: // east
			s.pos.x += i.param
		case 180: // south
			s.pos.y += i.param
		case 270: // west
			s.pos.x -= i.param
		}

	}

	if i.action == "L" || i.action == "R" {
		fmt.Printf("move %v => %v\n", i, s)
	}
}

func NewInstruction(val string) instruction {
	param, _ := strconv.Atoi(val[1:])
	return instruction{
		action: val[0:1],
		param:  param,
	}
}

func LoadInstructions(data []string) []instruction {
	result := make([]instruction, len(data))
	for i, val := range data {
		result[i] = NewInstruction(val)
	}

	return result
}

func absInt(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}

func main() {
	instructions := LoadInstructions(advent.Load("./day12/input.txt"))

	ferry := ship{
		heading: 90, // east
		pos:     position{0, 0},
	}

	fmt.Println(instructions)
	fmt.Println(ferry)

	for _, inst := range instructions {
		ferry.Move(inst)
	}

	fmt.Printf("distance %d\n", absInt(ferry.pos.x+ferry.pos.y))
}
