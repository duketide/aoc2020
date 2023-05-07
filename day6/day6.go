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

func all_yes_counter(s []string) (int, int) {
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
	return len(seen), cnt
}

func d6p2(groups [][]string) string {
	var sum1, sum2 int
	for _, group := range groups {
		inc1, inc2 := all_yes_counter(group)
		sum1 += inc1
		sum2 += inc2
	}
	return "Part 1 " + fmt.Sprint(sum1) + " Part 2 " + fmt.Sprint(sum2)
}

func Day6() string {
	groups := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	lines := make([][]string, len(groups))
	for i, v := range groups {
		lines[i] = parse.Lines(v)
	}
	return "Day 6 " + d6p2(lines)
}
