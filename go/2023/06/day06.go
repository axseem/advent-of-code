package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseNumbers(line string) []int {
	var numbers []int
	for _, s := range strings.Fields(line) {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return numbers
}

func parseNumber(line string) int {
	n, _ := strconv.Atoi(strings.Join(strings.Fields(line), ""))
	return n
}

func greaterThanDistance(limit, distance int) func(int) bool {
	return func(n int) bool {
		return (limit-n)*n > distance
	}
}

func lessEqualThanDistance(limit, distance int) func(int) bool {
	return func(n int) bool {
		return (limit-n)*n <= distance
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	limits := parseNumbers(lines[0][11:])
	distances := parseNumbers(lines[1][11:])

	ans := 1
	for i := 0; i < len(limits); i++ {
		start := sort.Search(limits[i], greaterThanDistance(limits[i], distances[i]))
		end := sort.Search(limits[i], lessEqualThanDistance(limits[i], distances[i]))
		ans *= end - start
	}

	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	limit := parseNumber(lines[0][11:])
	distance := parseNumber(lines[1][11:])

	start := sort.Search(limit, greaterThanDistance(limit, distance))
	end := sort.Search(limit, lessEqualThanDistance(limit, distance))

	return end - start
}

func main() {
	fmt.Println("--- 2023 day 06 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
