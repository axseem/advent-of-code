package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func part1(lines []string) int {
	var sum int
	for lineIndex, line := range lines {
		for i := 6; i < len(line); i++ {
			if line[i] == ':' {
				line = line[i+2:]
				break
			}
		}
		// is faster then:
		// _, line, _ = strings.Cut(line, ": ")

		isPossible := true
		var start int
		for i := 1; i < len(line); i++ {
			if line[i] == ' ' {
				var amount int
				for j := 0; j < len(line[start:i]); j++ {
					amount = amount*10 + int(line[start:i][j]-'0')
				}
				// is faster then:
				// amount, _ := strconv.Atoi(line[start:i])

				var r, g, b int
				switch line[i+1] {
				case 'r':
					r = amount
					i += 6
					start = i

				case 'g':
					g = amount
					i += 8
					start = i

				case 'b':
					b = amount
					i += 7
					start = i
				}

				if r > MAX_RED || g > MAX_GREEN || b > MAX_BLUE {
					isPossible = false
					break
				}
			}
		}

		if isPossible {
			sum += lineIndex + 1
		}
	}

	return sum
}

func part2(lines []string) int {
	var sum int
	for _, line := range lines {
		for i := 6; i < len(line); i++ {
			if line[i] == ':' {
				line = line[i+2:]
				break
			}
		}
		// is faster then:
		// _, line, _ = strings.Cut(line, ": ")

		var start, r, g, b int
		for i := 1; i < len(line); i++ {
			if line[i] == ' ' {
				var amount int
				for j := 0; j < len(line[start:i]); j++ {
					amount = amount*10 + int(line[start:i][j]-'0')
				}
				// is faster then:
				// amount, _ := strconv.Atoi(line[start:i])

				switch line[i+1] {
				case 'r':
					r = max(r, amount)
					i += 6
					start = i

				case 'g':
					g = max(g, amount)
					i += 8
					start = i

				case 'b':
					b = max(b, amount)
					i += 7
					start = i
				}
			}
		}

		sum += r * g * b
	}

	return sum
}

func main() {
	lines := strings.Split(input, "\n")

	fmt.Println("--- 2023 day 02 answer ---")
	fmt.Println("part 1:\t", part1(lines))
	fmt.Println("part 2:\t", part2(lines))
}
