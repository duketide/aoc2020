package day18

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed day18.txt
var input string

type level struct {
	operator    int
	accumulator int
}

type level_2 struct {
	operator      int
	accumulator   int
	add_accum     int
	held_operator int
}

var op_map = map[int]func(i, j int) int{
	0: func(i, j int) int { return j },
	1: func(i, j int) int { return i + j },
	2: func(i, j int) int { return i * j },
}

func eval_level(l *level, n *int) {
	(*l).accumulator = op_map[(*l).operator]((*l).accumulator, *n)
	(*l).operator = 0
}

func eval(s *string) int {
	tracker := make(map[int]*level)
	l := 0
	tracker[0] = &level{0, 0}
	for _, char := range *s {
		switch char {
		case '(':
			l++
			tracker[l] = &level{0, 0}
			continue

		case ')':
			bring_up := tracker[l].accumulator
			l--
			eval_level(tracker[l], &bring_up)
			continue
		case '+':
			tracker[l].operator = 1
			continue
		case '*':
			tracker[l].operator = 2
			continue
		//looks like no zeroes
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, e := strconv.Atoi(string(char))
			err.Check(e)
			eval_level(tracker[l], &num)
		default:
			panic("eval panic")

		}
	}
	return tracker[0].accumulator
}

func eval_level_2(l *level_2, n *int) {
	switch (*l).operator {
	case 1:
		(*l).add_accum += (*n)
		if (*l).held_operator == 0 {
			(*l).held_operator = 1
		}
		break
	case 2:
		if (*l).add_accum > 0 {
			(*l).accumulator = op_map[l.held_operator]((*l).accumulator, (*l).add_accum)
		}
		(*l).held_operator = 2
		(*l).add_accum = (*n)
	case 0:
		(*l).accumulator = *n
	}
	(*l).operator = 0
}

func eval_2(s *string) int {
	tracker := make(map[int]*level_2)
	l := 0
	tracker[0] = &level_2{0, 0, 0, 0}
	for _, char := range *s {
		switch char {
		case '(':
			l++
			tracker[l] = &level_2{0, 0, 0, 0}
			continue

		case ')':
			bring_up := tracker[l].accumulator
			if tracker[l].add_accum > 0 {
				bring_up = op_map[tracker[l].held_operator](bring_up, tracker[l].add_accum)
			}
			l--
			eval_level_2(tracker[l], &bring_up)
			continue
		case '+':
			tracker[l].operator = 1
			continue
		case '*':
			tracker[l].operator = 2
			continue
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, e := strconv.Atoi(string(char))
			err.Check(e)
			eval_level_2(tracker[l], &num)
		default:
			panic("eval_2 panic")
		}

	}
	result := tracker[0].accumulator
	if tracker[0].held_operator > 0 {
		result = op_map[tracker[0].held_operator](result, tracker[0].add_accum)
	}
	return result
}

func Day18() string {
	defer perf.Duration(perf.Track("Day18"))
	lines := parse.Lines(input)
	for i, line := range lines {
		var new_line string
		for _, char := range line {
			if char != ' ' {
				new_line += string(char)
			}
		}
		lines[i] = new_line
	}
	var p1 int
	for _, line := range lines {
		p1 += eval(&line)
	}
	var p2 int
	for _, line := range lines {
		p2 += eval_2(&line)
	}
	return "Day 18 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
