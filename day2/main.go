package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./day2/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	fmt.Printf("len: %d\n", len(lines))

	valid := 0
	valid2 := 0

	for _, line := range lines {

		parts := strings.Split(line, ":")
		pasword := strings.TrimSpace(parts[1])
		cond := strings.Split(parts[0], " ")
		conditionRune := cond[1]
		condition := cond[0]
		from := condition[:strings.Index(condition, "-")]
		to := condition[strings.Index(condition, "-")+1:]

		fmt.Printf("from %s to %s letter %s pass %s\n", from, to, conditionRune, pasword)

		fromCount, _ := strconv.Atoi(from)
		toCount, _ := strconv.Atoi(to)

		conditionFound := 0
		for i := range pasword {
			if pasword[i] == conditionRune[0] {
				conditionFound++
			}
		}

		if conditionFound >= fromCount && conditionFound <= toCount {
			valid++
		}

		if len(pasword) < toCount {
			if len(pasword) > fromCount && pasword[fromCount-1] == conditionRune[0] {
				valid2++
			}
		} else {
			if pasword[fromCount-1] == conditionRune[0] && pasword[toCount-1] != conditionRune[0] {
				valid2++
			}

			if pasword[fromCount-1] != conditionRune[0] && pasword[toCount-1] == conditionRune[0] {
				valid2++
			}

		}

	}

	fmt.Printf("Valid: %d\n", valid)
	fmt.Printf("Valid2: %d\n", valid2)
}
