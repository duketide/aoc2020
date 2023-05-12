package day15

import (
	"aoc2020/err"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day15.txt
var input string

func Day15() string {
	defer perf.Duration(perf.Track("Day15"))
	strs := strings.Split(strings.TrimRight(input, "\n"), ",")
	list := make([]int, len(strs))
	for i, v := range strs {
		num, e := strconv.Atoi(v)
		err.Check(e)
		list[i] = num
	}
	tracker := make(map[int]int)
	for i, num := range list {
		if i != len(list)-1 {
			tracker[num] = i
		}
	}
	idx := len(list) - 1
	last := list[len(list)-1]
	var p1, p2 int
	for idx < 30000000 {
		val, member := tracker[last]
		tracker[last] = idx
		if idx == 2019 {
			p1 = last
		}
		if idx == 29999999 {
			p2 = last
		}
		last = 0
		if member {
			last = idx - val
		}
		idx++
	}
	return "Day 15 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
