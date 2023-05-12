package day14

import (
	"aoc2020/err"
	"aoc2020/parse"
	"aoc2020/perf"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day14.txt
var input string

func spawn_float(addresses []string) []string {
	var new_addresses []string
	for _, address := range addresses {
		new_addresses = append(new_addresses, address+"0", address+"1")
	}
	return new_addresses
}

func spawn_addresses(mask, add string) []string {
	addresses := make([]string, 1)
	for i := 0; i < len(add); i++ {
		if mask[i] != 'X' {
			for index, str := range addresses {
				if mask[i] == '1' {
					addresses[index] = str + "1"
				} else {
					addresses[index] = str + string(add[i])
				}
			}
		} else {
			addresses = spawn_float(addresses)
		}
	}
	return addresses
}

func Day14() string {
	defer perf.Duration(perf.Track("Day14"))
	lines := parse.Lines(input)
	instructions := make([][]string, len(lines))
	for i, line := range lines {
		instructions[i] = strings.Split(line, " = ")
	}
	memory := make(map[string]string)
	memory2 := make(map[string]int)
	var mask string
	for _, inst := range instructions {
		if inst[0] == "mask" {
			//set mask and restart loop
			mask = inst[1]
			continue
		}

		//convert value to binary
		val, e := strconv.Atoi(inst[1])
		err.Check(e)
		bin_val := strconv.FormatUint(uint64(val), 2)
		for len(bin_val) < 36 {
			bin_val = "0" + bin_val
		}

		//apply bitmask to value for part 1
		new_bin_val := make([]byte, len(bin_val))
		for i := 0; i < len(bin_val); i++ {
			if mask[i] == 'X' {
				new_bin_val[i] = bin_val[i]
			} else {
				new_bin_val[i] = mask[i]
			}
		}

		//add bitmasked value to memory for part 1
		mem := inst[0][4:(len(inst[0]) - 1)]
		memory[mem] = string(new_bin_val)

		//for part 2, convert memory address to binary
		mem_dec, e2 := strconv.Atoi(mem)
		err.Check(e2)
		bin_mem := strconv.FormatUint(uint64(mem_dec), 2)
		for len(bin_mem) < 36 {
			bin_mem = "0" + bin_mem
		}

		//apply part 2 masking rule to address
		for _, address := range spawn_addresses(mask, bin_mem) {
			memory2[address] = val
		}
	}

	//sum values in memory for parts 1 and 2
	var p1 int64
	for _, v := range memory {
		num, e := strconv.ParseInt(v, 2, 64)
		err.Check(e)
		p1 += num
	}
	var p2 int
	for _, v := range memory2 {
		p2 += v
	}
	return "Day 14 Part 1 " + fmt.Sprint(p1) + " Part 2 " + fmt.Sprint(p2)
}
