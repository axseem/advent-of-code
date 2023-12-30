package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseValues(values []int, line string) {
	line += " "
	var start, vIndex int
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			n, _ := strconv.Atoi(line[start:i])
			values[vIndex] = n
			start = i + 1
			vIndex++
		}
	}
}

func nextValue(v []int) int {
	var nv int
	var az bool
	for !az {
		az = true
		nv += v[len(v)-1]
		for i := 1; i < len(v); i++ {
			v[i-1] = v[i] - v[i-1]

			if v[i-1] != 0 {
				az = false
			}
		}
		v = v[:len(v)-1]
	}

	return nv
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	length := len(strings.Split(lines[0], " "))
	values := make([]int, length)

	var sum int
	for _, line := range lines {
		parseValues(values, line)
		sum += nextValue(values)
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	length := len(strings.Split(lines[0], " "))
	values := make([]int, length)

	var sum int
	for _, line := range lines {
		parseValues(values, line)
		slices.Reverse(values)
		sum += nextValue(values)
	}

	return sum
}

func main() {
	fmt.Println("--- 2023 day 09 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
