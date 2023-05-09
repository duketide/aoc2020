package day11

import (
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day11.txt
var input string

type coords = struct {
	row int
	col int
}

func neighbors(pt coords, rowMax, colMax int) []coords {
	//maxes are permissible; max + 1 is oob
	var result []coords
	for r := pt.row - 1; r < pt.row+2 && r <= rowMax; r++ {
		if r < 0 {
			continue
		}
		for c := pt.col - 1; c < pt.col+2 && c <= colMax; c++ {
			if c < 0 || (c == pt.col && r == pt.row) {
				continue
			}
			result = append(result, coords{row: r, col: c})
		}
	}
	return result
}

func getNext(seats []string, pt coords) (byte, bool) {
	seat := (seats)[pt.row][pt.col]
	if seat == '.' {
		return seat, false
	}
	rowMax := len(seats) - 1
	colMax := len((seats)[0]) - 1
	nbrs := neighbors(pt, rowMax, colMax)
	occ_cnt := 0
	for _, v := range nbrs {
		if (seats)[v.row][v.col] == '#' {
			occ_cnt++
		}
	}
	if seat == '#' && occ_cnt >= 4 {
		return 'L', true
	}
	if seat == 'L' && occ_cnt == 0 {
		return '#', true
	}
	return seat, false
}

func single_turn(seats []string) ([]string, bool) {
	next := make([]string, len(seats))
	change := false
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len((seats)[i]); j++ {
			next_seat, test := getNext(seats, coords{row: i, col: j})
			next[i] = next[i] + string(next_seat)
			if !change {
				change = test
			}
		}
	}
	return next, change
}

func multi_turn(seats []string) *[]string {
	result, change := single_turn(seats)
	for change {
		result, change = single_turn(result)
	}
	return &result
}

func d11p1(seats *[]string) string {
	var cnt int
	for _, v := range strings.Join(*multi_turn(*seats), "") {
		if v == '#' {
			cnt++
		}
	}
	return "Part 1 " + fmt.Sprint(cnt)
}

func Day11() string {
	defer perf.Duration(perf.Track("Day11"))
	lines := parse.Lines(strings.TrimRight(input, "\n"))
	return "Day 11 " + d11p1(&lines)
}
