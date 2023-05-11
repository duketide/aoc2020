package day13

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day13.txt
var input string

func Day13() string {
	defer perf.Duration(perf.Track("Day13"))
	lines := parse.Lines(input)
	arrival, e := strconv.Atoi(lines[0])
	err.Check(e)
	raw_ids := strings.Split(lines[1], ",")
	var ids []int
	var id_map = make(map[int]int)
	for index, id := range raw_ids {
		if id == "x" {
			continue
		}
		id, e2 := strconv.Atoi(id)
		err.Check(e2)
		ids = append(ids, id)
		id_map[id] = index
	}
	var p1 int
	var min_wait = 1000000
	for _, id := range ids {
		if arrival%id == 0 {
			p1 = 0
			break
		}
		total_min := (arrival/id + 1) * id
		wait := total_min - arrival
		if wait < min_wait {
			min_wait = wait
			p1 = wait * id
		}
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] > ids[j] })
	flag := false
	n, c := 0, 1
	var p2 int
	for i := c; !flag; i += c {
		flag = true
		if (i+id_map[ids[n]])%ids[n] == 0 {
			c *= ids[n]
			n++
		}
		for k, v := range id_map {
			if (i+v)%k != 0 {
				flag = false
				break
			}

		}
		if flag {
			p2 = i
		}
	}

	return "Day 13 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
