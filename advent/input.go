package advent

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Load(f string) []string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return []string{}
	}

	if len(data) == 0 {
		return []string{}
	}

	lines := strings.Split(string(data), "\n")

	for i := range lines {
		l := strings.TrimSpace(lines[i])
		if l == "" {
			continue
		}
		lines[i] = l
	}

	return lines
}

func LoadInt(f string) []int {
	lines := Load(f)
	result := make([]int, len(lines))
	for i := range lines {
		if val, err := strconv.Atoi(lines[i]); err != nil {
			panic(err)
		} else {
			result[i] = val
		}
	}

	return result
}
