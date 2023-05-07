package day2

import (
	"aoc2020/err"
	"aoc2020/parse"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day2.txt
var input string

func parse_input(l []string) (int, int, byte, string) {
	rng := strings.Split(l[0], "-")
	min, errMin := strconv.Atoi(rng[0])
	max, errMax := strconv.Atoi(rng[1])
	err.Check(errMin)
	err.Check(errMax)
	ltr := l[1][0]
	str := l[2]
	return min, max, ltr, str
}

func d2p1(l [][]string) string {
	cnt := 0
	for _, v := range l {
		min, max, ltr, str := parse_input(v)
		subCnt := 0
		for _, v2 := range str {
			if v2 == rune(ltr) {
				subCnt++
			}
		}
		if subCnt >= min && subCnt <= max {
			cnt++
		}
		subCnt = 0

	}
	return "Part 1 " + fmt.Sprint(cnt)
}

func d2p2(l [][]string) string {
	cnt := 0
	for _, v := range l {
		i1, i2, ltr, str := parse_input(v)
		subCnt := 0
		if str[i1-1] == ltr {
			subCnt++
		}
		if str[i2-1] == ltr {
			subCnt++
		}
		if subCnt == 1 {
			cnt++
		}
		subCnt = 0
	}
	return " Part 2 " + fmt.Sprint(cnt)
}

func Day2() string {
	lines := parse.Lines(input)
	words := make([][]string, len(lines))
	for index, val := range lines {
		words[index] = parse.Words(val)
	}
	return "Day 2 " + d2p1(words) + d2p2(words)
}
