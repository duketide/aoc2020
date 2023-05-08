package day8

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day8.txt
var input string

type inst = struct {
	inst    string
	sign    byte
	num     int
	visited bool
}

func single(i *inst, idx int, acc int, swap bool) (int, int, bool) {
	if i.visited {
		return 0, acc, true
	}
	i.visited = true
	var n int
	if i.sign == '-' {
		n = (-1) * i.num
	} else {
		n = i.num
	}

	var new_inst string
	if swap && i.inst == "nop" {
		new_inst = "jmp"
	} else if swap && i.inst == "jmp" {
		new_inst = "nop"
	} else {
		new_inst = i.inst
	}

	switch new_inst {
	case "nop":
		return idx + 1, acc, false
	case "jmp":
		return idx + n, acc, false
	case "acc":
		return idx + 1, acc + n, false
	default:
		panic("oops")
	}
}

func d8p1(i []*inst) string {
	idx, acc, flag := 0, 0, false
	for flag == false {
		idx, acc, flag = single(i[idx], idx, acc, false)
	}
	return "Part 1 " + fmt.Sprint(acc)
}

func d8p2(i []*inst) string {
	idx := 0
	idx_list := []int{}
	target := len(i)
	for ind, val := range i {
		if val.visited {
			val.visited = false
			idx_list = append(idx_list, ind)
		}
	}
	for idx < len(i) {
		idx, acc, flag := 0, 0, false
		swap_idx := idx_list[0]
		idx_list = idx_list[1:]
		for _, val := range i {
			val.visited = false
		}
		for flag == false {
			swap_bool := false
			if idx == swap_idx {
				swap_bool = true
			}
			idx, acc, flag = single(i[idx], idx, acc, swap_bool)
			if idx == target {
				return " Part 2 " + fmt.Sprint(acc)
			}
			if idx > target {
				break
			}
		}
	}
	panic("yikes")
}

func Day8() string {
	defer perf.Duration(perf.Track("Day8"))
	lines := parse.Lines(strings.TrimRight(input, "\n"))
	instructions := make([]*inst, len(lines))
	for i, v := range lines {
		num, e := strconv.Atoi(v[5:])
		err.Check(e)
		instructions[i] = &inst{
			inst:    v[0:3],
			sign:    v[4],
			num:     num,
			visited: false,
		}
	}
	return "Day 8 " + d8p1(instructions) + d8p2(instructions)
}
