package day1

import (
	s "aoc2020/ds/set"
	"aoc2020/parse"
	_ "embed"
	"fmt"
)

//go:embed day1.txt
var input string

func d1p1(set s.IntSet) string {
	for val := range set {
		target := 2020 - val
		if s.IsMember(set, target) {
			return "Part 1 " + fmt.Sprint(val*target)
		}
	}
	return "Part 1 miss"
}

func d1p2(set s.IntSet) string {
	copiedSet := s.CopySet(set)
	for k := range set {
		delete(copiedSet, k)
		for k2 := range copiedSet {
			target := 2020 - k - k2
			if s.IsMember(set, target) {
				return " Part 2 " + fmt.Sprint(k*k2*target)
			}
		}
	}
	return " Part 2 miss"
}

func Day1() string {
	set := s.StrArrToIntSet(parse.Lines(input))
	return "Day 1 " + d1p1(set) + d1p2(set)
}
