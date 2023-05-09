package main

import (
	"aoc2020/day1"
	"aoc2020/day10"
	"aoc2020/day11"
	"aoc2020/day2"
	"aoc2020/day3"
	"aoc2020/day4"
	"aoc2020/day5"
	"aoc2020/day6"
	"aoc2020/day7"
	"aoc2020/day8"
	"aoc2020/day9"
	"aoc2020/perf"
	"fmt"
)

func main() {
	defer perf.Duration(perf.Track("Total"))
	fmt.Println(day1.Day1())
	fmt.Println(day2.Day2())
	fmt.Println(day3.Day3())
	fmt.Println(day4.Day4())
	fmt.Println(day5.Day5())
	fmt.Println(day6.Day6())
	fmt.Println(day7.Day7())
	fmt.Println(day8.Day8())
	fmt.Println(day9.Day9())
	fmt.Println(day10.Day10())
	fmt.Println(day11.Day11())
}
