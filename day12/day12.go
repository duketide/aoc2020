package day12

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"math"
	"strconv"
)

//go:embed day12.txt

var input string

type facing = int

const (
	E facing = 0
	S facing = 90
	W facing = 180
	N facing = 270
)

var face_map = map[byte]facing{'E': E, 'S': S, 'W': W, 'N': N}

type state = struct {
	inst       []string
	facing     facing
	E, S, W, N int
}

type state_2 = struct {
	inst             []string
	facing           facing
	E, S, W, N, X, Y int
}

func rotate(inst string, f facing) facing {
	dir := inst[0]
	deg, e := strconv.Atoi(inst[1:])
	err.Check(e)

	//right is positive, left is negative
	if dir == 'R' {
		return (f + deg) % 360
	}
	raw_deg := f - deg
	if raw_deg < 0 {
		return 360 + raw_deg
	}
	return raw_deg
}

func single_turn(s *state) *state {
	inst := s.inst[0]
	s.inst = s.inst[1:]
	num, e := strconv.Atoi(inst[1:])
	err.Check(e)
	switch inst[0] {
	case 'L', 'R':
		s.facing = rotate(inst, s.facing)
		return s
	case 'N':
		s.N += num
		return s
	case 'E':
		s.E += num
		return s
	case 'S':
		s.S += num
		return s
	case 'W':
		s.W += num
		return s
	case 'F':
		switch s.facing {
		case E:
			s.E += num
			return s
		case S:
			s.S += num
			return s
		case W:
			s.W += num
			return s
		case N:
			s.N += num
			return s
		default:
			panic(s.facing)
		}
	default:
		panic("ouch")
	}
}

func multi_turn(s *state) {
	for len(s.inst) > 0 {
		s = single_turn(s)
	}
}

func d12p1(inst []string) string {
	s := state{inst: inst}
	multi_turn(&s)
	return "Part 1 " + fmt.Sprint(math.Abs(float64(s.N-s.S))+math.Abs(float64(s.E-s.W)))
}

func move_waypoint(s *state_2, d facing, num int) {
	f := s.facing
	diff := f - d
	if diff < 0 {
		diff = 360 + diff
	}
	switch diff {
	case 0:
		s.Y += num
	case 180:
		s.Y -= num
	case 90:
		s.X -= num
	case 270:
		s.X += num
	default:
		panic("waypoint issue")
	}
}

func mover(s *state_2, f facing, dist int) {
	if dist < 0 {
		dist = dist * -1
	}
	switch f {
	case E:
		s.E += dist
		return
	case S:
		s.S += dist
		return
	case W:
		s.W += dist
		return
	case N:
		s.N += dist
		return
	default:
		panic("mover" + fmt.Sprint(f))
	}
}

func move_boat(s *state_2, mult int) {
	x_move, y_move := s.X*mult, s.Y*mult
	var x_dir, y_dir facing = (s.facing + 90) % 360, s.facing
	if x_move < 0 {
		x_dir = s.facing - 90
		if x_dir < 0 {
			x_dir += 360
		}
	}
	mover(s, x_dir, x_move)
	if y_move < 0 {
		y_dir = s.facing - 180
		if y_dir < 0 {
			y_dir += 360
		}
	}
	mover(s, y_dir, y_move)
}

func single_turn_2(s *state_2) {
	inst := s.inst[0]
	s.inst = s.inst[1:]
	ltr := inst[0]
	if ltr == 'L' || inst[0] == 'R' {
		s.facing = rotate(inst, s.facing)
		return
	}
	num, e := strconv.Atoi(inst[1:])
	err.Check(e)
	switch ltr {
	case 'F':
		move_boat(s, num)
		return
	case 'E', 'S', 'W', 'N':
		move_waypoint(s, face_map[ltr], num)
		return
	default:
		panic("single turn 2 issue")
	}
}

func multi_turn_2(s *state_2) {
	for len(s.inst) > 0 {
		single_turn_2(s)
	}
}

func d12p2(inst []string) string {
	s := state_2{inst: inst, X: -1, Y: 10}
	multi_turn_2(&s)
	return " Part 2 " + fmt.Sprint(math.Abs(float64(s.N-s.S))+math.Abs(float64(s.E-s.W)))
}

func Day12() string {
	defer perf.Duration(perf.Track("Day12"))
	lines := parse.Lines(input)
	return "Day 12 " + d12p1(lines) + d12p2(lines)
}
