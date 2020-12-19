package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/vm"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	returnSuccess = 0
)

func load(lines []string) []vm.Instruction {
	code := make([]vm.Instruction, 0)
	for _, line := range lines {
		code = append(code, parse(line))
	}
	return code
}

func parse(line string) vm.Instruction {
	argv := strings.Split(line, " ")
	arg, _ := strconv.Atoi(argv[1])
	return vm.Instruction{
		OP:  argv[0],
		ARG: arg,
	}
}

func print(code []vm.Instruction) {
	fmt.Println()
	for _, op := range code {
		fmt.Printf("  %s %d\n", op.OP, op.ARG)
	}
	fmt.Println()
}

func run(code []vm.Instruction) (int, int) {
	register := 0
	addr := 0
	visited := map[int]int{}

	for addr != len(code) {

		if addr < 0 || addr > len(code) {
			return -2, register // illegal instruction
		}

		inst := code[addr]
		visited[addr]++

		if visited[addr] > 1 {
			fmt.Printf("Found loop at addr:%d acc:%d\n", addr, register)
			return -1, register // loop
		}

		switch inst.OP {
		case "nop":
			addr++
		case "acc":
			register += inst.ARG
			addr++
		case "jmp":
			addr += inst.ARG
		}

	}

	return 0, register
}

func main() {

	data, _ := ioutil.ReadFile("./day8/input.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Code:%s\n", lines)

	modifier := 0
	code := load(lines)

	for {

		retCode, register := run(code)
		fmt.Printf("return retcode:%d register:%d\n", retCode, register)

		if retCode == returnSuccess {
			fmt.Printf("### return retcode:%d register:%d\n", retCode, register)
			break
		}

		code = load(lines)
		fmt.Printf("Code mutate with modifier:%d\n", modifier)

		if modifier >= len(code)-1 {
			fmt.Printf("*** error could not find fix")
			break
		}

		for i := modifier; i < len(code); i++ {
			if code[i].OP == "nop" {
				code[i].OP = "jmp"
				modifier = i + 1
				break
			}

			if code[i].OP == "jmp" {
				code[i].OP = "nop"
				modifier = i + 1
				break
			}
		}

		fmt.Printf("Code mutate with modifier:%d\n", modifier)
	}
}
