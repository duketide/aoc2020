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

func moves(cups *[]int, num_moves int) {
	current := (*cups)[0]
	num_cups := len(*cups)
	curr_index := 0
	mvs := num_moves
	for mvs > 0 {
		move_indices := []int{(curr_index + 1) % num_cups, (curr_index + 2) % num_cups, (curr_index + 3) % num_cups}
		mover_set := make(map[int]bool)
		for _, v := range move_indices {
			mover_set[v] = true
		}
		target := current - 1
		target_found := false
		for !target_found {
			if target == 0 {
				target = num_cups
			}
			for i, v := range *cups {
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
		for i, v := range *cups {
			if mover_set[i] {
				continue
			}
			new_nums = append(new_nums, v)
			if v == target {
				for _, ind := range move_indices {
					new_nums = append(new_nums, (*cups)[ind])
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
		(*cups) = new_nums
		mvs--
	}
}
func moves2(cups *[]int, num_moves int) {
	num_cups := len(*cups)
	mvs := num_moves
	for mvs > 0 {
		current := (*cups)[0]
		movers := (*cups)[1:4]
		rest := (*cups)[4:]
		rest_copy := make([]int, len(rest))
		copy(rest_copy, rest)
		mover_set := make(map[int]bool)
		for _, v := range movers {
			mover_set[v] = true
		}
		target := current - 1
		for mover_set[target] || target < 1 {
			if target < 1 {
				target = num_cups
			}
			if mover_set[target] {
				target--
			}
		}
		for i, v := range rest {
			if v == target {
				*cups = append(rest_copy[:i+1], append(movers, append(rest[i+1:], current)...)...)
				mvs--
				if mvs%10000 == 0 {
				}
				break
			}
		}
	}
}

func Day23() string {
	defer perf.Duration(perf.Track("Day23"))
	var nums []int
	for _, ltr := range strings.TrimRight(input, "\n") {
		num, e := strconv.Atoi(string(ltr))
		err.Check(e)
		nums = append(nums, num)
	}
	nums2 := make([]int, 1000000)
	copy(nums2, nums)
	for i := range nums2 {
		if i > 8 {
			nums2[i] = i + 1
		}
	}
	moves2(&nums, 100)
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
	moves2(&nums2, 10000000)
	one_index_2 := -1
	for i, v := range nums2 {
		if v == 1 {
			one_index_2 = i
			break
		}
	}
	p2 := nums2[(one_index_2+1)%1000000] * nums2[(one_index_2+2)%1000000]
	return "Day 23 Part 1 " + p1 + " Part 2 " + fmt.Sprint(p2)
}
