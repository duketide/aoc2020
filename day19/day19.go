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

type loop_state = struct {
	eights, elevens, eight_len, eleven_len, max_len int
}

var inc = 0

func p() int { inc++; return inc }

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
			} else if len(r) == 3 {
				r1, r2, r3 := eval(rules, &r[0]), eval(rules, &r[1]), eval(rules, &r[2])
				for _, s := range r1 {
					for _, t := range r2 {
						for _, u := range r3 {
							result = append(result, s+t+u)
						}
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

func parse_rules(raw_rules *[]string) *map[string]*rule {
	rules := make(map[string]*rule)
	for _, r := range *raw_rules {
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
	return &rules
}

//I could see that all strings returned at 42 and 31 were of length 8
//That was crucial for this @check function

func check(msg *string, set_42, set_31 *map[string]bool) bool {
	matches := true
	second := false
	l := 0
	forty_twos := 0
	thirty_ones := 0
	length := len(*msg)
	for matches && l+8 <= length {
		test := (*msg)[l : l+8]
		if !second {
			matches = (*set_42)[test]
			if matches {
				forty_twos++
			}
			if !matches {
				second = true
			}
		}
		if second {
			matches = (*set_31)[test]
			if matches {
				thirty_ones++
			}
		}
		l += 8
	}
	if thirty_ones >= forty_twos || forty_twos < 2 || !(*set_31)[(*msg)[length-8:length]] {
		matches = false
	}
	return matches
}

func Day19() string {
	defer perf.Duration(perf.Track("Day19"))
	split := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	raw_rules := parse.Lines(split[0])
	rules := parse_rules(&raw_rules)
	messages := parse.Lines(split[1])
	start, forty_two, thirty_one := "0", "42", "31"
	zero_slice := eval(rules, &start)
	forty_twos := eval(rules, &forty_two)
	thirty_ones := eval(rules, &thirty_one)
	zero_set, set_42, set_31 := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	for _, s := range zero_slice {
		zero_set[s] = true
	}
	for _, val := range forty_twos {
		set_42[val] = true
	}
	for _, val := range thirty_ones {
		set_31[val] = true
	}
	p1, p2 := 0, 0
	for _, s := range messages {
		if zero_set[s] {
			p1++
		}
		if check(&s, &set_42, &set_31) {
			p2++
		}
	}
	return "Day 19 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
