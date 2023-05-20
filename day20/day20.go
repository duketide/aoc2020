package day20

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed day20.txt
var input string

type square = map[int]bool

func getBits(s string) (uint, uint) {
	l := len(s)
	var result, rev_result float64
	for i, bit := range s {
		if bit == '.' {
			continue
		}
		power := float64(l - i - 1)
		result += math.Pow(2, power)
		rev_result += math.Pow(2, float64(i))
	}
	return uint(result), uint(rev_result)
}

func getSquare(lines []string) []uint {
	up := lines[0]
	down := lines[len(lines)-1]
	var right, left string
	for _, line := range lines {
		left = left + string(line[0])
		right = right + string(line[len(line)-1])
	}
	result := make([]uint, 8)
	for i, side := range []string{up, right, down, left} {
		m, n := getBits(side)
		result[i*2], result[i*2+1] = m, n
	}
	return result
}

func Day20() string {
	defer perf.Duration(perf.Track("Day20"))
	raw_tiles := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	tiles := make(map[int][]uint)
	border_map := make(map[uint][]int)
	for _, tile := range raw_tiles {
		lines := parse.Lines(tile)
		name_line := parse.Words(lines[0])
		n := strings.TrimRight(name_line[1], ":")
		num, e := strconv.Atoi(n)
		err.Check(e)
		tiles[num] = getSquare(lines[1:])
	}
	for tile, borders := range tiles {
		for _, border := range borders {
			border_map[border] = append(border_map[border], tile)
		}
	}
	//assuming intra-tile borders are unique
	singles := make(map[int]int)
	for _, tile_list := range border_map {
		if len(tile_list) == 1 {
			singles[tile_list[0]] += 1
		}
	}
	p1 := 1
	for k, v := range singles {
		if v == 4 {
			p1 *= k
		}
	}
	return "Day 20 Part 1 " + fmt.Sprint(p1)
}
