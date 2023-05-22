package day20

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed day20.txt
var input string

//go:embed day20monster.txt
var monster string

type m_bit = struct {
	row, col int
}

type si = int

const (
	first  si = 0
	second si = 1
	third  si = 2
	fourth si = 3
)

type orientation = struct {
	flipped bool
	start   si
}

type square = struct {
	id                 int
	unflipped, flipped []int
	orientation        orientation
}

type tile = []string

func reverse10(n int) int {
	bin := strconv.FormatUint(uint64(n), 2)
	offset := len(bin)
	var result float64 = 0
	for i, d := range bin {
		if d == '1' {
			result += math.Pow(2, float64(10-offset+i))
		}
	}
	return int(result)
}

func rev_string(s string) string {
	result := ""
	for _, ltr := range s {
		result = string(ltr) + result
	}
	return result
}

func rotate_ctr_clk(t tile) tile {
	//function assumes square tiles
	var result tile
	for i := len(t[0]) - 1; i >= 0; i-- {
		s := ""
		for j := 0; j < len(t); j++ {
			s = s + string(t[j][i])
		}
		result = append(result, s)
	}
	return result
}

func flip(t tile) tile {
	var result tile
	for _, line := range t {
		result = append(result, rev_string(line))
	}
	return result
}

func flip_and_rotate(s square, t tile) tile {
	o := s.orientation
	var result tile
	if o.flipped && o.start == 2 {
		for _, line := range t {
			result = append(tile{line}, result...)
		}
		return result
	}
	if o.flipped {
		result = flip(t)
	} else {
		result = t
	}
	for i := 0; i < o.start; i++ {
		result = rotate_ctr_clk(result)
	}
	//fix this return
	return result

}

func getBits(s string) (int, int) {
	l := len(s)
	var result, rev_result float64
	for i, bit := range s {
		if bit == '.' {
			continue
		}
		power := float64(l - i - 1)
		result += math.Pow(2, power)
		rev_result += math.Pow(2, float64(i))
	}
	return int(result), int(rev_result)
}

func getSquare(lines []string) []int {
	// this function reads bits clockwise
	// the second half of the returned array is the flipped square
	// so it swaps the left and right sides (indices 4 and 6)
	up := lines[0]
	down := lines[len(lines)-1]
	var right, left string
	for _, line := range lines {
		left = left + string(line[0])
		right = right + string(line[len(line)-1])
	}
	unflipped := make([]int, 8)
	for i, side := range []string{left, up, right, down} {
		m, n := getBits(side)
		if i == 1 || i == 2 {
			unflipped[i], unflipped[i+4] = m, n
		} else {
			unflipped[i], unflipped[i+4] = n, m
		}
	}
	flipped := make([]int, 8)
	for i, side := range unflipped {
		if i == 4 {
			flipped[6] = side
		} else if i == 6 {
			flipped[4] = side
		} else {
			flipped[i] = side
		}
	}
	return flipped
}

func monster_counter(sea *[]string, m_bits *[]m_bit) int {
	count := 0
	for i, r := range *sea {
		if i > len(*sea)-3 {
			break
		}
		for j := range r {
			if j > len(r)-20 {
				break
			}
			m := true
			for _, bit := range *m_bits {
				if (*sea)[i+bit.row][j+bit.col] != '#' {
					m = false
					break
				}
			}
			if m {
				count++
			}
		}
	}
	return count
}

func Day20() string {
	defer perf.Duration(perf.Track("Day20"))
	raw_tiles := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	tiles := make(map[int][]int)
	tile_literals := make(map[int]tile)
	border_map := make(map[int][]int)
	for _, tile := range raw_tiles {
		lines := parse.Lines(tile)
		name_line := parse.Words(lines[0])
		n := strings.TrimRight(name_line[1], ":")
		num, e := strconv.Atoi(n)
		err.Check(e)
		tiles[num] = getSquare(lines[1:])
		lit_tile := lines[2 : len(lines)-1]
		for i, line := range lit_tile {
			lit_tile[i] = line[1 : len(line)-1]
		}
		tile_literals[num] = lit_tile
	}
	for tile, borders := range tiles {
		for _, border := range borders {
			border_map[border] = append(border_map[border], tile)
		}
	}
	//assuming intra-tile borders are unique
	singles := make(map[int]int)
	for _, tile_list := range border_map {
		if len(tile_list) == 1 {
			singles[tile_list[0]] += 1
		}
	}
	p1 := 1
	for k, v := range singles {
		if v == 4 {

			p1 *= k
		}
	}
	var corners []square
	for k, v := range singles {
		if v == 4 {
			vals := tiles[k]
			start := -1
			for i := 0; i < 4; i++ {
				ok1 := len(border_map[vals[i]]) == 1
				ok2 := len(border_map[vals[(i+1)%4]]) == 1
				if ok1 && ok2 {
					start = i
					break
				}
			}
			if start == -1 {
				panic("failed to find start")
			}
			corners = append(corners, square{
				id:        k,
				unflipped: vals[:4],
				flipped:   vals[4:],
				orientation: orientation{
					flipped: false,
					start:   start,
				},
			})
		}
	}
	//square is 12 by 12, so a gross
	gross := make([][]square, 12)
	gross[0] = append(gross[0], corners[0])
	for i := 1; i < len(gross); i++ {
		match := gross[i-1][0]
		match_id := match.id
		match_flip := match.orientation.flipped
		match_start := match.orientation.start
		var match_set []int
		if match_flip {
			match_set = match.flipped
		} else {
			match_set = match.unflipped
		}
		match_bottom := reverse10(match_set[(3+match_start)%4])
		var next_tile_id int
		var next_tile []int
		for _, tile := range border_map[match_bottom] {
			if tile != match_id {
				next_tile_id = tile
				next_tile = tiles[tile]
			}
		}
		start_index := -1
		next_flip := false
		for i, v := range next_tile {
			if v == match_bottom {
				start_index = (i - 1) % 4
				if start_index == -1 {
					start_index = 3
				}
				next_flip = i > 3
			}
		}
		gross[i] = append(gross[i], square{
			id:        next_tile_id,
			unflipped: next_tile[:4],
			flipped:   next_tile[4:],
			orientation: orientation{
				flipped: next_flip,
				start:   start_index,
			},
		})
	}
	for i := range gross {
		for j := 1; j < 12; j++ {
			match := gross[i][j-1]
			match_id := match.id
			match_flip := match.orientation.flipped
			match_start := match.orientation.start
			var match_set []int
			if match_flip {
				match_set = match.flipped
			} else {
				match_set = match.unflipped
			}
			match_right := reverse10(match_set[(2+match_start)%4])
			var next_tile_id int
			var next_tile []int
			for _, tile := range border_map[match_right] {
				if tile != match_id {
					next_tile_id = tile
					next_tile = tiles[tile]
				}
			}
			start_index := -1
			next_flip := false
			for i, v := range next_tile {
				if v == match_right {
					start_index = i
					next_flip = i > 3
				}
			}
			gross[i] = append(gross[i], square{
				id:        next_tile_id,
				unflipped: next_tile[:4],
				flipped:   next_tile[4:],
				orientation: orientation{
					flipped: next_flip,
					start:   start_index,
				},
			})
		}
	}
	tile_slice := make([][]tile, 12)
	for i, line := range gross {
		for _, tile := range line {
			tile_slice[i] = append(tile_slice[i], flip_and_rotate(tile, tile_literals[tile.id]))
		}
	}
	raw_map := make([]string, 96)
	for i, row_of_tiles := range tile_slice {
		for _, tile := range row_of_tiles {
			for k, row := range tile {
				raw_map[i*8+k] = raw_map[i*8+k] + row
			}
		}
	}
	m_slice := parse.Lines(monster)
	var m_bits []m_bit
	for i, line := range m_slice {
		for j, bit := range line {
			if bit == '#' {
				m_bits = append(m_bits, m_bit{row: i, col: j})
			}
		}
	}
	total_hash := 0
	for _, str := range raw_map {
		for _, bit := range str {
			if bit == '#' {
				total_hash++
			}
		}
	}
	m_count := 0
	for i := 0; i < 8; i++ {
		test := monster_counter(&raw_map, &m_bits)
		if test > 0 {
			m_count = test
			break
		}
		if i == 3 {
			raw_map = flip(raw_map)
		} else {
			raw_map = rotate_ctr_clk(raw_map)
		}
	}
	return "Day 20 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(total_hash-(m_count*len(m_bits)))
}
