package day17

import (
	"aoc2020/ds/set"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
)

//go:embed day17.txt
var input string

type coords = struct {
	x int
	y int
	z int
}

type coords_4 = struct {
	x int
	y int
	z int
	w int
}

func get_touches(cube *coords) *[]coords {
	result := make([]coords, 0)
	l := []int{-1, 0, 1}
	for _, i := range l {
		for _, j := range l {
			for _, k := range l {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				result = append(result, coords{x: cube.x + i, y: cube.y + j, z: cube.z + k})
			}
		}
	}
	return &result
}

func get_touches_4(cube *coords_4) *[]coords_4 {
	result := make([]coords_4, 0)
	l := []int{-1, 0, 1}
	for _, h := range l {
		for _, i := range l {
			for _, j := range l {
				for _, k := range l {
					if h == 0 && i == 0 && j == 0 && k == 0 {
						continue
					}
					result = append(result, coords_4{x: cube.x + i, y: cube.y + j, z: cube.z + k, w: cube.w + h})
				}
			}
		}
	}
	return &result
}

func Day17() string {
	defer perf.Duration(perf.Track("Day17"))
	lines := parse.Lines(input)
	grid := make(map[coords]bool)
	grid_4 := make(map[coords_4]bool)
	for x, line := range lines {
		for y, val := range line {
			if val == '#' {
				grid[coords{x: x, y: y, z: 0}] = true
				grid_4[coords_4{x: x, y: y, z: 0, w: 0}] = true
			}
		}
	}
	cycle := 0
	for cycle < 6 {
		new_grid := make(map[coords]bool)
		touches := make(map[coords]int)
		for cube := range grid {
			for _, touch := range *get_touches(&cube) {
				touches[touch]++
			}
		}
		for k, v := range touches {
			if v == 3 || (v == 2 && set.IsMember(grid, k)) {
				new_grid[k] = true
			}
		}
		grid = new_grid
		cycle++
	}
	cycle_4 := 0
	for cycle_4 < 6 {
		new_grid := make(map[coords_4]bool)
		touches := make(map[coords_4]int)
		for cube := range grid_4 {
			for _, touch := range *get_touches_4(&cube) {
				touches[touch]++
			}
		}
		for k, v := range touches {
			if v == 3 || (v == 2 && set.IsMember(grid_4, k)) {
				new_grid[k] = true
			}
		}
		grid_4 = new_grid
		cycle_4++
	}
	return "Day 17 Part 1 " + fmt.Sprint(len(grid)) + " Part 2 " + fmt.Sprint(len(grid_4))
}
