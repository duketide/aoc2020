package day7

import (
	"aoc2020/ds/set"
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day7.txt
var input string

type void = struct{}

var member void

type str_set = map[string]void

type edges = struct {
	color string
	num   int
}

func dfs_no_num(graph *map[string][]edges, start string, target string) bool {
	visited := make(str_set)
	stack := []string{start}
	for true {
		if len(stack) == 0 {
			return false
		}
		last := len(stack) - 1
		current := stack[last]
		stack = stack[:last]
		if set.IsMember(visited, current) {
			continue
		}
		visited[current] = member
		for _, v := range (*graph)[current] {
			if v.color == target {
				return true
			}
			if set.IsMember(visited, v.color) || v.color == "terminus" {
				continue
			}
			stack = append(stack, v.color)
		}
	}
	panic("yikes")
}

func d7p1(graph *map[string][]edges) string {
	var cnt int
	for k := range *graph {
		if dfs_no_num(graph, k, "shiny gold") {
			cnt++
		}
	}
	return "Part 1 " + fmt.Sprint(cnt)
}

func dfs_num(graph *map[string][]edges, start string) int {
	stack := []string{start}
	var cnt int
	for len(stack) != 0 {
		last := len(stack) - 1
		current := stack[last]
		stack = stack[:last]
		for _, v := range (*graph)[current] {
			if v.color == "terminus" {
				continue
			}
			for i := v.num; i > 0; i-- {
				stack = append(stack, v.color)
				cnt++
			}
		}
	}
	return cnt
}

func d7p2(graph *map[string][]edges) string {
	return " Part 2 " + fmt.Sprint(dfs_num(graph, "shiny gold"))
}

func Day7() string {
	defer perf.Duration(perf.Track("Day7"))
	lines := parse.Lines(input)
	contains := make([][]string, len(lines))
	for i, v := range lines {
		contains[i] = strings.Split(strings.TrimRight(v, "."), " bags contain ")
	}
	graph := make(map[string][]edges)
	for _, v := range contains {
		if v[1] == "no other bags" {
			graph[v[0]] = []edges{{color: "terminus", num: 0}}
			continue
		}
		raw_inside := strings.Split(v[1], ", ")
		words_inside := make([]edges, len(raw_inside))
		for i, str := range raw_inside {
			j := parse.Words(str)
			n, e := strconv.Atoi(j[0])
			err.Check(e)
			words_inside[i] = edges{color: strings.Join(j[1:3], " "), num: n}
			graph[v[0]] = words_inside
		}
	}
	return "Day 7 " + d7p1(&graph) + d7p2(&graph)
}
