package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
	lines := strings.Split(input, "\n")

	var first, last, sum int
	for _, line := range lines {
		for i := range line {
			if line[i] >= '1' && line[i] <= '9' {
				first = int(line[i] - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '1' && line[i] <= '9' {
				last = int(line[i] - '0')
				break
			}
		}

		sum += first*10 + last
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var first, last, sum int
	for _, line := range lines {
	FirstLoop:
		for i := range line {
			if line[i] >= '1' && line[i] <= '9' {
				first = int(line[i] - '0')
				break
			}
			for j, digit := range digits {
				for k := 0; k < len(digit) && i+k < len(line) && line[i+k] == digit[k]; k++ {
					if len(digit)-1 == k {
						first = j + 1
						break FirstLoop
					}
				}
			}
		}

	LastLoop:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '1' && line[i] <= '9' {
				last = int(line[i] - '0')
				break
			}
			for j, digit := range digits {
				for k := 0; k < len(digit) && i-k >= 0 && line[i-k] == digit[len(digit)-1-k]; k++ {
					if len(digit)-1 == k {
						last = j + 1
						break LastLoop
					}
				}
			}
		}

		sum += first*10 + last
	}

	return sum
}

func main() {
	fmt.Println("--- 2023 day 01 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
