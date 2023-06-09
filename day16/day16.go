package day16

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

//go:embed day16.txt
var input string

func in_a_range(ranges *[][]int, n *int) bool {
	for _, rs := range *ranges {
		if (*n) >= rs[0] && (*n) <= rs[1] {
			return true
		}
	}
	return false
}

func get_field_index(num *int, tix *[][]int, f_ranges *[][][]int, seen *[][]int, perm_seen *[][]int, unavailable *[]int) int {
	for i := range *f_ranges {
		seen_pass := false
		for _, n := range append((*seen)[*num], (*perm_seen)[*num]...) {
			if n == i {
				seen_pass = true
				break
			}
		}
		if seen_pass {
			continue
		}
		for _, n := range *unavailable {
			if n == i {
				seen_pass = true
				break
			}
		}
		if seen_pass {
			continue
		}
		return i
	}
	return -1
}

func Day16() string {
	defer perf.Duration(perf.Track("Day16"))

	//parse input
	segments := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	fields := parse.Lines(segments[0])
	pfields := make([][]string, len(fields))
	for i, line := range fields {
		pfields[i] = strings.Split(line, ": ")
	}
	field_names, field_ranges := make([]string, len(pfields)), make([]string, len(pfields))
	for i, field := range pfields {
		field_names[i] = field[0]
		field_ranges[i] = field[1]
	}
	pfield_ranges := make([][]string, len(field_ranges))
	for i, f_range := range field_ranges {
		pfield_ranges[i] = strings.Split(f_range, " or ")
	}

	//need bare_field_ranges for part 2
	bare_field_ranges := make([][][]int, len(pfield_ranges))
	for i, f_ranges := range pfield_ranges {
		for _, f_range := range f_ranges {
			strs := strings.Split(f_range, "-")
			nums := make([]int, len(strs))
			for k, str := range strs {
				num, e := strconv.Atoi(str)
				err.Check(e)
				nums[k] = num
			}
			bare_field_ranges[i] = append(bare_field_ranges[i], nums)
		}
	}

	//need just_ranges for part 1
	var just_ranges [][]int
	for _, x := range bare_field_ranges {
		for _, y := range x {
			just_ranges = append(just_ranges, y)
		}
	}

	//add my ticket to end of list of nearby tickets, convert to []int
	my_ticket := parse.Lines(segments[1])[1:][0]
	n_tickets := append(parse.Lines(segments[2])[1:], my_ticket)
	s_tickets := make([][]int, len(n_tickets))
	for i, t := range n_tickets {
		x := strings.Split(t, ",")
		y := make([]int, len(x))
		for j, n := range x {
			num, e := strconv.Atoi(n)
			err.Check(e)
			y[j] = num
		}
		s_tickets[i] = y
	}

	//need invalid_nums for part 1, valid_tix for part 2
	invalid_nums := make(map[int]int)
	var valid_tix [][]int

	//main loop for part 1 (though valid_tix is only for part 2)
	//exclude my ticket from this loop per instructions
	for _, t := range s_tickets[:len(s_tickets)-1] {
		invalid := false
		for _, num := range t {
			if !in_a_range(&just_ranges, &num) {
				invalid_nums[num]++
				invalid = true
			}
		}
		if !invalid {
			valid_tix = append(valid_tix, t)
		}
	}

	//add my ticket back to the valid ticket list for part 2
	my_int_ticket := s_tickets[len(s_tickets)-1]
	valid_tix = append(valid_tix, my_int_ticket)

	//sum invalid numbers to get answer for part 1
	var p1 int
	for k, v := range invalid_nums {
		p1 += k * v
	}
	t_length := len(s_tickets[0])

	//because index_in_fields will be used to determine which fields
	//remain to be assigned, it needs to be initialized with a value
	//that won't give bad results (e.g., -1)
	index_in_fields := make([]int, t_length)
	for i := range index_in_fields {
		index_in_fields[i] = -1
	}

	//temp_seen is a mutable list of what's been tried at each position
	//it can reset due to backtracking
	temp_seen := make([][]int, t_length)

	//perm_seen is a position_by_position list of invalid fields
	//it should not be changed
	perm_seen := make([][]int, t_length)
	curr := 0
	for curr < t_length {
		for i, f := range bare_field_ranges {
			for _, t := range valid_tix {
				if !in_a_range(&f, &t[curr]) {
					perm_seen[curr] = append(perm_seen[curr], i)
					break
				}
			}
		}
		curr++
	}

	//to optimize, we want to evaluate ticket positions from
	//fewest # of valid possibilities to greatest
	//this optimization changed run time from >1m to <1s
	field_test_order := make([]int, t_length)
	for i := range field_test_order {
		field_test_order[i] = i
	}
	sort.Slice(field_test_order, func(i, j int) bool { return len(perm_seen[i]) > len(perm_seen[j]) })

	//main loop for part 2
	curr = 0
	for curr < t_length {
		tick_index := field_test_order[curr]
		test := get_field_index(&tick_index, &valid_tix, &bare_field_ranges, &temp_seen, &perm_seen, &index_in_fields)
		index_in_fields[tick_index] = test

		//backtrack
		if test == -1 {
			temp_seen[tick_index] = make([]int, 0)
			curr--
			tick_index := field_test_order[curr]
			temp_seen[tick_index] = append(temp_seen[tick_index], index_in_fields[tick_index])
			continue
		}
		curr++
	}

	//get answer for part 2
	p2 := 1
	for i, val := range index_in_fields {
		if val <= 5 {
			p2 *= my_int_ticket[i]
		}
	}
	return "Day 16 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
