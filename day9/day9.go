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

func d9p1(nums *[]int) int {
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
			return target
		}
		current++
	}
	panic("hmph")
}

func d9p2(nums *[]int) string {
	part_one := d9p1(nums)
	left, right := 0, 1
	end := len(*nums) - 1
	sum := (*nums)[left] + (*nums)[right]
	var max, min, part_two int
	for left < end && right < end {
		if sum == part_one {
			max, min = (*nums)[left], (*nums)[left]
			for _, v := range (*nums)[left : right+1] {
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
			}
			part_two = max + min
			break
		}
		if sum < part_one || left == right-1 {
			right++
			sum += (*nums)[right]
		}
		if sum > part_one {
			sum -= (*nums)[left]
			left++
		}
	}
	return "Part 1 " + fmt.Sprint(part_one) + " Part 2 " + fmt.Sprint(part_two)
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
	return "Day 9 " + d9p2(&nums)
}
