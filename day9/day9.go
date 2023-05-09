package day9

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day9.txt
var input string

func d9p1(nums *[]int) string {
	current := 25
	for current < len(*nums) {
		target := (*nums)[current]
		flag := false
		for i := current - 25; i < current; i++ {
			if flag == true {
				break
			}
			for j := i + 1; j < current; j++ {
				if (*nums)[i]+(*nums)[j] == target {
					flag = true
					break
				}
			}
		}
		if !flag {
			return "Part 1 " + fmt.Sprint(target)
		}
		current++
	}
	panic("hmph")
}

func Day9() string {
	defer perf.Duration(perf.Track("Day9"))
	lines := parse.Lines(strings.TrimRight(input, "\n"))
	nums := make([]int, len(lines))
	for i, v := range lines {
		num, e := strconv.Atoi(v)
		err.Check(e)
		nums[i] = num
	}
	return "Day 9 " + d9p1(&nums)
}
