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

func moves(cups *[]int, num_moves int, starter int) {
	num_cups := len(*cups) - 1
	mvs := num_moves
	current := starter
	for mvs > 0 {
		mover1 := (*cups)[current]
		mover2 := (*cups)[mover1]
		mover3 := (*cups)[mover2]
		next := (*cups)[mover3]
		dest := current - 1
		mover_set := make(map[int]bool)
		for _, val := range []int{mover1, mover2, mover3} {
			mover_set[val] = true
		}
		for mover_set[dest] || dest < 1 {
			if dest < 1 {
				dest = num_cups
			}
			if mover_set[dest] {
				dest--
			}
		}
		(*cups)[mover3] = (*cups)[dest]
		(*cups)[dest] = mover1
		(*cups)[current] = next
		current = next
		mvs--
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
	starter := nums[0]
	//create a pseudo-linked-list with []int
	numsll := make([]int, len(nums)+1)
	for i, num := range nums {
		numsll[num] = nums[(i+1)%len(nums)]
	}
	nums2 := make([]int, 1000001)
	copy(nums2, numsll)
	nums2[nums[len(nums)-1]] = 10
	for i := 10; i < 1000000; i++ {
		nums2[i] = i + 1
	}
	nums2[1000000] = nums[0]
	moves(&numsll, 100, starter)
	c := numsll[1]
	p1 := fmt.Sprint(c)
	for numsll[c] != 1 {
		p1 += fmt.Sprint(numsll[c])
		c = numsll[c]
	}
	moves(&nums2, 10000000, starter)
	first := nums2[1]
	second := nums2[first]
	p2 := first * second
	return "Day 23 Part 1 " + p1 + " Part 2 " + fmt.Sprint(p2)
}
