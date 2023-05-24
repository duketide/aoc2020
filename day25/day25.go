package day25

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed day25.txt
var input string

func Day25() string {
	defer perf.Duration(perf.Track("Day25"))
	lines := parse.Lines(input)
	a, e := strconv.Atoi(lines[0])
	err.Check(e)
	b, e2 := strconv.Atoi(lines[1])
	err.Check(e2)
	const sub_num = 7
	val := 1
	a_loop := 0
	b_loop := 0
	for i := 1; a_loop*b_loop == 0; i++ {
		val *= sub_num
		val %= 20201227
		if val == a {
			a_loop = i
		}
		if val == b {
			b_loop = i
		}
	}
	val = 1
	for i := 0; i < b_loop; i++ {
		val *= a
		val %= 20201227
	}
	return "Day 25 " + fmt.Sprint(val)
}
