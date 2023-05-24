package day24

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day24.txt
var input string

type pair = struct {
	x, y int
}

func stringify(x, y int) string {
	return fmt.Sprint(x) + "|" + fmt.Sprint(y)
}

func destringify(s string) (int, int) {
	ints := strings.Split(s, "|")
	x, e := strconv.Atoi(ints[0])
	err.Check(e)
	y, e2 := strconv.Atoi(ints[1])
	err.Check(e2)
	return x, y
}

func Day24() string {
	defer perf.Duration(perf.Track("Day24"))
	lines := parse.Lines(input)
	tracker := make(map[string]int)
	for _, line := range lines {
		var x, y int = 0, 0
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'w':
				x -= 2
				continue
			case 'e':
				x += 2
				continue
			case 's':
				y--
				if line[i+1] == 'w' {
					x--
				} else {
					x++
				}
				i++
				continue
			case 'n':
				y++
				if line[i+1] == 'w' {
					x--
				} else {
					x++
				}
				i++
				continue
			default:
				panic("busted direction")
			}
		}
		tracker[stringify(x, y)]++
	}
	p1 := 0
	p2_tracker := make(map[string]bool)
	for k, v := range tracker {
		if v%2 == 1 {
			p1++
			p2_tracker[k] = true
		}
	}
	/*var adjacents []pair

	turns := 100
	for turns >0 {
	}*/

	return "Day 24 Part 1 " + fmt.Sprint(p1)
}
