package day5

import (
	"aoc2020/ds/set"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
)

//go:embed day5.txt
var input string

type void = struct{}

var member void

func one_step(dir byte, min, max int) (int, int) {
	mid := (min + max) / 2
	if dir == 'F' || dir == 'L' {
		return min, mid
	} else {
		return mid + 1, max
	}
}

func many_steps(s string) int {
	var inner func(s string, rmin, rmax, cmin, cmax int) int
	inner = func(s string, rmin, rmax, cmin, cmax int) int {
		if len(s) == 0 {
			return rmax*8 + cmax
		}
		switch s[0] {
		case 'F', 'B':
			rmin, rmax = one_step(s[0], rmin, rmax)
			return inner(s[1:], rmin, rmax, cmin, cmax)
		case 'L', 'R':
			cmin, cmax = one_step(s[0], cmin, cmax)
			return inner(s[1:], rmin, rmax, cmin, cmax)
		default:
			panic("busted switch")
		}
	}
	return inner(s, 0, 127, 0, 7)
}

func d5p2(l []string) string {
	var max int
	seen := make(set.IntSet)
	var seat int
	for _, v := range l {
		val := many_steps(v)
		if val > max {
			max = val
		}
		seen[val] = member
	}
	for k := range seen {
		target := k + 1
		if !set.IsMember(seen, target) && set.IsMember(seen, target+1) {
			seat = target
			break
		}
	}
	return "Part 1 " + fmt.Sprint(max) + " Part 2 " + fmt.Sprint(seat)

}

func Day5() string {
	defer perf.Duration(perf.Track("Day5"))
	lines := parse.Lines(input)
	return "Day 5 " + d5p2(lines)
}
