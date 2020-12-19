package main

import (
	"fmt"
	"github.com/yavosh/advent-of-code-2020/valid"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./day4/input.txt")
	lines := strings.Split(string(data), "\n")

	var passports []map[string]string
	passport := map[string]string{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}

		for _, keyval := range strings.Split(line, " ") {
			index := strings.IndexAny(keyval, ":")
			key := keyval[0:index]
			val := keyval[index+1:]
			passport[key] = val
		}
	}

	passports = append(passports, passport)

	wantedKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	sort.Strings(wantedKeys)

	validFieldsPassports := 0
	validPassports := 0

	for _, p := range passports {
		keys := map[string]bool{}
		keysOnly := []string{}
		for key := range p {
			keys[key] = true
			keysOnly = append(keysOnly, key)
		}

		sort.Strings(keysOnly)

		allValid := true

		for _, key := range wantedKeys {
			allValid = allValid && keys[key]
		}

		if allValid {
			validFieldsPassports++
		}

		birthYearChecker := valid.MinMaxChecker{Min: 1920, Max: 2002}
		if !birthYearChecker.Check(p["byr"]) {
			allValid = false
		}

		issueYearChecker := valid.MinMaxChecker{Min: 2010, Max: 2020}
		if !issueYearChecker.Check(p["iyr"]) {
			allValid = false
		}

		expireYearChecker := valid.MinMaxChecker{Min: 2020, Max: 2030}
		if !expireYearChecker.Check(p["eyr"]) {
			allValid = false
		}

		colourChecker := valid.ColourChecker{}
		if !colourChecker.Check(p["hcl"]) {
			allValid = false
		}

		eyeChecker := valid.NewOneOfChecker("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
		if !eyeChecker.Check(p["ecl"]) {
			allValid = false
		}

		//passport false          158898193 9
		if found, _ := regexp.MatchString("^\\d{9}$", p["pid"]); p["pid"] != "" && !found {
			allValid = false
		}

		heightChecker := valid.HeightChecker{MinCm: 150, MaxCm: 193, MinInch: 59, MaxInch: 76}
		if !heightChecker.Check(p["hgt"]) {
			allValid = false
		}

		if allValid {
			validPassports++
		}

		fmt.Printf("passport %t \t\t%v %d\n", allValid, p["pid"], len(p["pid"]))
	}

	fmt.Printf("Passports total=%d validFields=%d\n", len(passports), validFieldsPassports)
	fmt.Printf("Passports total=%d valid=%d\n", len(passports), validPassports)

	// 159 too high
}
