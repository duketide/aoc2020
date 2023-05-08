package day3

import (
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
)

//go:embed day3.txt
var input string

func gen(l []string, colInc int, rowInc int) int {
	cnt := 0
	wrap := len(l[0])
	for row := 0; row < len(l); row += rowInc {
		col := (row * colInc / rowInc) % wrap
		if l[row][col] == '#' {
			cnt++
		}
	}
	return cnt
}

func d3p1(l []string) string {
	return "Part 1 " + fmt.Sprint(gen(l, 3, 1))
}

func d3p2(l []string) string {
	return " Part 2 " + fmt.Sprint(gen(l, 1, 1)*gen(l, 3, 1)*gen(l, 5, 1)*gen(l, 7, 1)*gen(l, 1, 2))
}

func Day3() string {
	defer perf.Duration(perf.Track("Day3"))
	lines := parse.Lines(input)
	return "Day 3 " + d3p1(lines) + d3p2(lines)
}
