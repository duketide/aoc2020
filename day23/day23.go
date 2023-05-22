package day23

import (
	"aoc2020/err"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day23.txt
var input string

func Day23() string {
	defer perf.Duration(perf.Track("Day23"))
	var nums []int
	for _, ltr := range strings.TrimRight(input, "\n") {
		num, e := strconv.Atoi(string(ltr))
		err.Check(e)
		nums = append(nums, num)
	}
	current := nums[0]
	curr_index := 0
	moves := 100
	for moves > 0 {
		move_indices := []int{(curr_index + 1) % 9, (curr_index + 2) % 9, (curr_index + 3) % 9}
		mover_set := make(map[int]bool)
		for _, v := range move_indices {
			mover_set[v] = true
		}
		target := current - 1
		target_found := false
		for !target_found {
			if target == 0 {
				target = 9
			}
			for i, v := range nums {
				if v == target {
					if !mover_set[i] {
						target_found = true
						break
					} else {
						target--
						break
					}
				}
			}
		}
		var new_nums []int
		for i, v := range nums {
			if mover_set[i] {
				continue
			}
			new_nums = append(new_nums, v)
			if v == target {
				for _, ind := range move_indices {
					new_nums = append(new_nums, nums[ind])
				}
			}
		}
		for i, v := range new_nums {
			if v == current {
				curr_index = (i + 1) % 9
				current = new_nums[curr_index]
				break
			}
		}
		nums = new_nums
		moves--
	}
	one_index := -1
	for i, v := range nums {
		if v == 1 {
			one_index = i
			break
		}
	}
	p1 := ""
	for i := 1; i < 9; i++ {
		p1 += fmt.Sprint(nums[(i+one_index)%9])
	}
	return "Day 23 Part 1 " + p1
}
