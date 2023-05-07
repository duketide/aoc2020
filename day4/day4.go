package day4

import (
	"aoc2020/parse"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day4.txt
var input string
var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasField(l []string, field string) bool {
	for _, v := range l {
		if v[:3] == field {
			return true
		}
	}
	return false
}

func hasFields(l []string) bool {
	for _, v := range fields {
		if !hasField(l, v) {
			return false
		}
	}
	return true
}

func checkYear(s string, min int, max int) bool {
	v, err := strconv.Atoi(s)
	if err != nil || len(s) != 4 || v < min || v > max {
		return false
	}
	return true
}

func test(str string) bool {
	data := str[4:]
	switch str[:3] {
	case "byr":
		return checkYear(data, 1920, 2002)
	case "iyr":
		return checkYear(data, 2010, 2020)
	case "eyr":
		return checkYear(data, 2020, 2030)
	case "hgt":
		switch len(data) {
		case 4:
			num, err := strconv.Atoi(data[:2])
			if data[2:] != "in" || err != nil || num < 59 || num > 76 {
				return false
			}
			return true
		case 5:
			num, err := strconv.Atoi(data[:3])
			if data[3:] != "cm" || err != nil || num < 150 || num > 193 {
				return false
			}
			return true

		default:
			return false
		}
	case "hcl":
		if len(data) != 7 || data[0] != '#' {
			return false
		}
		for _, val := range data[1:] {
			flag := false
			for _, i := range []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'a', 'b', 'c', 'd', 'e', 'f'} {
				if val == i {
					flag = true
					break
				}
			}
			if !flag {
				return false
			}
		}
		return true
	case "ecl":
		for _, val := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if data == val {
				return true
			}
		}
		return false
	case "pid":
		_, err := strconv.Atoi(data)
		if len(data) != 9 || err != nil {
			return false
		}
		return true
	case "cid":
		return true
	default:
		return false
	}
}

func testAll(l []string) bool {
	for _, s := range l {
		if !test(s) {
			return false
		}
	}
	return true
}

func d4p2(l [][]string) string {
	var cnt1 int
	var cnt2 int
	for _, l := range l {
		if hasFields(l) {
			cnt1++
			if testAll(l) {
				cnt2++
			}
		}
	}
	return "Part 1 " + fmt.Sprint(cnt1) + " Part 2 " + fmt.Sprint(cnt2)
}

func Day4() string {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	words := make([][]string, len(lines))
	for ind, str := range lines {
		words[ind] = parse.Words(strings.ReplaceAll(str, "\n", " "))
	}
	return "Day 4 " + d4p2(words)
}
