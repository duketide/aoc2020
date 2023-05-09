package day10

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day10.txt
var input string

func d10p1(nums *[]int) string {
	var ones, threes int
	for i := 1; i < len(*nums); i++ {
		diff := (*nums)[i] - (*nums)[i-1]
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}
	return "Part 1 " + fmt.Sprint(ones*threes)
}

func d10p2(nums *[]int) string {
	tracker := make([]int64, len(*nums))
	tracker[0] = 1
	oob := len(*nums)
	for i := 0; i < oob; i++ {
		val := tracker[i]
		for j := i + 1; j < i+4 && j < oob; j++ {
			diff := (*nums)[j] - (*nums)[i]
			if diff > 3 {
				continue
			}
			tracker[j] += val
		}
	}
	return " Part 2 " + fmt.Sprint(tracker[len(tracker)-1])
}

func Day10() string {
	defer perf.Duration(perf.Track("Day10"))
	lines := parse.Lines(strings.TrimRight(input, "\n"))
	nums := make([]int, len(lines)+1)
	nums[0] = 0
	for i, v := range lines {
		num, e := strconv.Atoi(v)
		err.Check(e)
		nums[i+1] = num
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	nums = append(nums, nums[len(nums)-1]+3)
	return "Day 10 " + d10p1(&nums) + d10p2(&nums)
}
