package day7

import (
	"aoc2020/ds/set"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day7.txt
var input string

type void = struct{}

var member void

type str_set = map[string]void

func dfs(graph map[string][]string, start string, target string) bool {
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
		for _, v := range graph[current] {
			if v == target {
				return true
			}
			if set.IsMember(visited, v) || v == "terminus" {
				continue
			}
			stack = append(stack, v)
		}
	}
	panic("yikes")
}

func d7p1(graph map[string][]string) string {
	var cnt int
	for k := range graph {
		if dfs(graph, k, "shiny gold") {
			cnt++
		}
	}
	return "Part 1 " + fmt.Sprint(cnt)
}

func Day7() string {
	defer perf.Duration(perf.Track("Day7"))
	lines := parse.Lines(input)
	contains := make([][]string, len(lines))
	for i, v := range lines {
		contains[i] = strings.Split(strings.TrimRight(v, "."), " bags contain ")
	}
	graph := make(map[string][]string)
	for _, v := range contains {
		if v[1] == "no other bags" {
			graph[v[0]] = []string{"terminus"}
			continue
		}
		raw_inside := strings.Split(v[1], ", ")
		words_inside := make([]string, len(raw_inside))
		for i, str := range raw_inside {
			words_inside[i] = strings.Join(parse.Words(str)[1:3], " ")
			graph[v[0]] = words_inside
		}
	}
	return "Day 7 " + d7p1(graph)
}
