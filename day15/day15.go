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
	for len(list) < 30000000 {
		idx := len(list) - 1
		last := list[idx]
		val, member := tracker[last]
		next := 0
		if member {
			next = idx - val
		}
		tracker[last] = idx
		list = append(list, next)
	}
	return "Day 15 Part 1 " + fmt.Sprint(list[2019]) + " Part 2 " + fmt.Sprint(list[29999999])
}
