package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseInstructions(line string) []int {
	instructions := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		if line[i] == 'L' {
			instructions[i] = 0
		} else {
			instructions[i] = 1
		}
	}
	return instructions
}

func lettersToIndex(s string) int {
	return int(s[0]-'A')*26*26 + int(s[1]-'A')*26 + int(s[2]-'A')
}

func parseNetwork(lines []string) ([][2]int, []int) {
	netwrok := make([][2]int, 26*26*26)
	var endWithA []int
	for _, line := range lines {
		index := lettersToIndex(line[:3])
		netwrok[index] = [2]int{
			lettersToIndex(line[7:10]),
			lettersToIndex(line[12:15]),
		}

		if line[2] == 'A' {
			endWithA = append(endWithA, index)
		}
	}

	return netwrok, endWithA
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(n []int) int {
	result := n[0] * n[1] / GCD(n[0], n[1])
	for i := 2; i < len(n); i++ {
		result = LCM([]int{result, n[i]})
	}
	return result
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	instructions := parseInstructions(lines[0])
	network, _ := parseNetwork(lines[2:])

	current := 0
	var step int
	for current != lettersToIndex("ZZZ") {
		instructionIndex := step % len(instructions)
		current = network[current][instructions[instructionIndex]]
		step++
	}

	return step
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	instructions := parseInstructions(lines[0])
	network, starts := parseNetwork(lines[2:])

	var steps []int
	for _, current := range starts {
		var step int
		for (current+1)%26 != 0 {
			instructionIndex := step % len(instructions)
			current = network[current][instructions[instructionIndex]]
			step++
		}
		steps = append(steps, step)
	}

	return LCM(steps)
}

func main() {
	fmt.Println("--- 2023 day 08 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
