package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}

func main() {
	lines := strings.Split(input, "\n")

	fmt.Println("--- 0000 day 00 answer ---")
	fmt.Println("part 1:\t", part1(lines))
	fmt.Println("part 2:\t", part2(lines))
}
