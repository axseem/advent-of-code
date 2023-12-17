package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func findStart(line string) int {
	for i := 6; i < len(line); i++ {
		if line[i] == ':' {
			return i + 2
		}
	}
	return 0
}

func findMidline(line string) int {
	for i := 9; i < len(line); i++ {
		if line[i] == '|' {
			return i
		}
	}
	return 0
}

func amountOfMatches(line string, start, midline int) int {
	var amount int
	for i := start; i < midline-1; i += 3 {
		for j := midline + 2; j < len(line); j += 3 {
			if line[i] == line[j] && line[i+1] == line[j+1] {
				amount++
				break
			}
		}
	}

	return amount
}

func pow2(n int) int {
	if n > 0 {
		return 1 << n
	} else if n == 0 {
		return 1
	}
	return 0
}

func part1(lines []string) int {
	start := findStart(lines[0])
	midline := findMidline(lines[0])

	var sum int
	for _, line := range lines {
		amount := amountOfMatches(line, start, midline)
		sum += pow2(amount - 1)
	}

	return sum
}

func part2(lines []string) int {
	start := findStart(lines[0])
	midline := findMidline(lines[0])

	var sum int
	cards := make([]int, len(lines))
	for lineIndex, line := range lines {
		amount := amountOfMatches(line, start, midline)

		for i := 0; i < amount; i++ {
			cards[lineIndex+1+i] += cards[lineIndex] + 1
		}
		sum += cards[lineIndex] + 1
	}

	return sum
}

func main() {
	lines := strings.Split(input, "\n")

	fmt.Println("--- 2023 day 04 answer ---")
	fmt.Println("part 1:\t", part1(lines))
	fmt.Println("part 2:\t", part2(lines))
}
