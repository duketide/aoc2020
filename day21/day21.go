package day21

import (
	"aoc2020/ds/set"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed day21.txt
var input string

type recipe = struct {
	ingredients, allergens map[string]bool
}

func get_singles(a map[string](map[string]bool)) map[string]bool {
	result := make(map[string]bool)
	for _, is := range a {
		if len(is) == 1 {
			for i := range is {
				result[i] = true
			}
		}
	}
	return result
}

func Day21() string {
	defer perf.Duration(perf.Track("Day21"))
	lines := parse.Lines(input)
	recipes := make([]recipe, len(lines))
	for i, line := range lines {
		n_line := strings.Split(line, " (contains ")
		ing := parse.Words(n_line[0])
		alrg := strings.Split(strings.TrimRight(n_line[1], ")"), ", ")
		ingredients, allergens := make(map[string]bool), make(map[string]bool)
		for _, i := range ing {
			ingredients[i] = true
		}
		for _, a := range alrg {
			allergens[a] = true
		}
		recipes[i] = recipe{ingredients, allergens}
	}
	alrg_map := make(map[string](map[string]bool))
	for _, rec := range recipes {
		for a := range rec.allergens {
			if set.IsMember(alrg_map, a) {
				for i := range alrg_map[a] {
					if !rec.ingredients[i] {
						delete(alrg_map[a], i)
					}
				}
			} else {
				alrg_map[a] = make(map[string]bool)
				for i := range rec.ingredients {
					alrg_map[a][i] = true
				}
			}
		}
	}
	possible_alrg := make(map[string]bool)
	for _, is := range alrg_map {
		for i := range is {
			possible_alrg[i] = true
		}
	}
	p1 := 0
	for _, rec := range recipes {
		for i := range rec.ingredients {
			if !possible_alrg[i] {
				p1++
			}
		}
	}
	unfinished := true
	for unfinished {
		unfinished = false
		singles := get_singles(alrg_map)
		for a, is := range alrg_map {
			if len(is) > 1 {
				unfinished = true
				for i := range is {
					if singles[i] {
						delete(alrg_map[a], i)
					}
				}
			}
		}
	}
	var p2_slice []string
	for a, v := range alrg_map {
		for k := range v {
			p2_slice = append(p2_slice, a+"|"+k)
		}
	}
	sort.Strings(p2_slice)
	p2 := ""
	for _, s := range p2_slice {
		p2 = p2 + strings.Split(s, "|")[1] + ","
	}
	return "Day 21 Part 1 " + fmt.Sprint(p1) + " Part 2 " + strings.TrimRight(fmt.Sprint(p2), ",")
}
