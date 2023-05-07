package day6

import (
	"aoc2020/parse"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day6.txt
var input string

type void = struct{}

var member void

func single_yes_counter(s []string) int {
	joined := strings.TrimRight(strings.Join(s, ""), "\n")
	seen := make(map[rune]void)
	for _, ltr := range joined {
		seen[ltr] = member
	}
	return len(seen)
}

func all_yes_counter(s []string) int {
	joined := strings.TrimRight(strings.Join(s, ""), "\n")
	seen := make(map[rune]int)
	for _, ltr := range joined {
		seen[ltr] += 1
	}
	var cnt int
	for _, v := range seen {
		if v == len(s) {
			cnt++
		}
	}
	return cnt
}

func d6p1(groups [][]string) string {
	var sum int
	for _, group := range groups {
		sum += single_yes_counter(group)
	}
	return "Part 1 " + fmt.Sprint(sum)
}

func d6p2(groups [][]string) string {
	var sum int
	for _, group := range groups {
		sum += all_yes_counter(group)
	}
	return " Part 2 " + fmt.Sprint(sum)
}

func Day6() string {
	groups := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	lines := make([][]string, len(groups))
	for i, v := range groups {
		lines[i] = parse.Lines(v)
	}
	return "Day 6 " + d6p1(lines) + d6p2(lines)
}
