package day19

import (
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day19.txt
var input string

type rule = struct {
	kind   string
	letter string
	single []string
	either [][]string
}

func eval(rules *map[string]*rule, rule_num *string) []string {
	rule := (*rules)[*rule_num]
	switch rule.kind {
	case "letter":
		return []string{rule.letter}
	case "single":
		if len(rule.single) == 1 {
			return eval(rules, &rule.single[0])
		} else if len(rule.single) == 2 {
			r1, r2 := eval(rules, &rule.single[0]), eval(rules, &rule.single[1])
			var result []string
			for _, s := range r1 {
				for _, t := range r2 {
					result = append(result, s+t)
				}
			}
			return result
		} else if len(rule.single) == 3 {
			r1, r2, r3 := eval(rules, &rule.single[0]), eval(rules, &rule.single[1]), eval(rules, &rule.single[2])
			var result []string
			for _, s := range r1 {
				for _, t := range r2 {
					for _, u := range r3 {
						result = append(result, s+t+u)
					}
				}
			}
			return result
		} else {
			panic("too many rules in a single")
		}
	case "either":
		var result []string
		for _, r := range rule.either {
			if len(r) == 1 {
				result = append(result, eval(rules, &r[0])...)
			} else if len(r) == 2 {
				r1, r2 := eval(rules, &r[0]), eval(rules, &r[1])
				for _, s := range r1 {
					for _, t := range r2 {
						result = append(result, s+t)
					}
				}
			} else {
				panic("problem in an either")
			}
		}
		return result
	default:
		panic("reached default")
	}
}

func Day19() string {
	defer perf.Duration(perf.Track("Day19"))
	split := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	raw_rules := parse.Lines(split[0])
	rules := make(map[string]*rule)
	for _, r := range raw_rules {
		x := strings.Split(r, ": ")
		num := x[0]
		val := x[1]
		if val[0] == '"' {
			rules[num] = &rule{kind: "letter", letter: string(val[1])}
			continue
		}
		pipes := strings.Split(val, " | ")
		if len(pipes) == 1 {
			words := parse.Words(pipes[0])
			rules[num] = &rule{
				kind:   "single",
				single: words,
			}
			continue
		}
		doub := make([][]string, 2)
		for i, val := range pipes {
			doub[i] = parse.Words(val)
		}
		rules[num] = &rule{
			kind:   "either",
			either: doub,
		}

	}
	start := "0"
	zero_slice := eval(&rules, &start)
	zero_set := make(map[string]bool)
	for _, s := range zero_slice {
		zero_set[s] = true
	}
	p1 := 0
	for _, s := range parse.Lines(split[1]) {
		if zero_set[s] {
			p1++
		}
	}
	return "Day 19 Part 1 " + fmt.Sprint(p1)
}
