package day22

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day22.txt
var input string

type won = struct {
	winner int
	hand   []int
}

func hand(p1, p2 *[]int) {
	//assume no ties possible
	c1, c2 := (*p1)[0], (*p2)[0]
	if c1 > c2 {
		(*p1) = append((*p1)[1:], c1, c2)
		(*p2) = (*p2)[1:]
	}
	if c2 > c1 {
		(*p2) = append((*p2)[1:], c2, c1)
		(*p1) = (*p1)[1:]
	}
}

func stringify(p1, p2 *[]int) string {
	str := ""
	for _, num := range *p1 {
		str = str + fmt.Sprint(num) + ","
	}
	str = str + "|"
	for _, num := range *p2 {
		str = str + fmt.Sprint(num) + ","
	}
	return str
}

func game(p1, p2 []int) won {
	memo := make(map[string]bool)
	for true {
		s := stringify(&p1, &p2)
		if memo[s] {
			return won{1, p1}
		}
		memo[s] = true
		hand_result := rec_hand(p1, p2)
		if hand_result == 1 {
			p1 = append(p1[1:], p1[0], p2[0])
			p2 = p2[1:]
		} else if hand_result == 2 {
			p2 = append(p2[1:], p2[0], p1[0])
			p1 = p1[1:]
		} else {
			panic("oof")
		}
		if len(p1) == 0 {
			return won{2, p2}
		}
		if len(p2) == 0 {
			return won{1, p1}
		}
	}
	panic("ouch")
}

func rec_hand(p1, p2 []int) int {
	if len(p1) >= p1[0]+1 && len(p2) >= p2[0]+1 {
		new1, new2 := make([]int, p1[0]), make([]int, p2[0])
		copy(new1, p1[1:p1[0]+1])
		copy(new2, p2[1:p2[0]+1])
		return game(new1, new2).winner
	} else if p1[0] > p2[0] {
		return 1
	} else {
		return 2
	}
}

func Day22() string {
	defer perf.Duration(perf.Track("Day22"))
	decks := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	pl1, pl2 := parse.Lines(decks[0]), parse.Lines(decks[1])
	var player1, player2 []int
	for i := 1; i < len(pl1); i++ {
		num, e := strconv.Atoi(pl1[i])
		err.Check(e)
		player1 = append(player1, num)
	}
	for i := 1; i < len(pl2); i++ {
		num, e := strconv.Atoi(pl2[i])
		err.Check(e)
		player2 = append(player2, num)
	}
	player1B, player2B := make([]int, len(player1)), make([]int, len(player2))
	copy(player1B, player1)
	copy(player2B, player2)
	for len(player1)*len(player2) != 0 {
		hand(&player1, &player2)
	}
	var winner []int
	if len(player1) > 0 {
		winner = player1
	} else {
		winner = player2
	}
	part1 := 0
	for i, card := range winner {
		part1 += (len(winner) - i) * card
	}
	var winnerB []int = game(player1B, player2B).hand
	part2 := 0
	for i, card := range winnerB {
		part2 += (len(winnerB) - i) * card
	}
	return "Day 22 Part 1 " + fmt.Sprint(part1) + " Part 2 " + fmt.Sprint(part2)
}
