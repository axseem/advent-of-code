package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type rule struct {
	start int
	end   int
	shift int
}

func parseSeeds(line string) []int {
	var seeds []int
	for _, s := range strings.Split(line[7:], " ") {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}

	return seeds
}

func parseSeedRanges(line string) [][2]int {
	parts := strings.Split(line[7:], " ")
	var seeds [][2]int
	for i := 0; i < len(parts); i += 2 {
		n1, _ := strconv.Atoi(parts[i])
		n2, _ := strconv.Atoi(parts[i+1])
		seeds = append(seeds, [2]int{n1, n1 + n2 - 1})
	}

	return seeds
}

func parseMaps(lines []string) [7][]rule {
	var maps [7][]rule
	var mapIndex int
	for _, line := range lines {
		if len(line) == 0 {
			mapIndex++
		} else if line[0] >= '0' && line[0] <= '9' {
			var numbers [3]int
			for i, s := range strings.Split(line, " ") {
				numbers[i], _ = strconv.Atoi(s)
			}
			maps[mapIndex] = append(maps[mapIndex], rule{
				start: numbers[1],
				end:   numbers[1] + numbers[2],
				shift: numbers[0] - numbers[1],
			})
		}
	}

	return maps
}

func convertRanges(ranges [][2]int, rules []rule) [][2]int {
	var shifted [][2]int
	for _, r := range rules {
		for i := 0; i < len(ranges); i++ {
			n := &ranges[i]
			startInside := n[0] >= r.start && n[0] <= r.end
			endInside := n[1] >= r.start && n[1] <= r.end
			opositeSides := n[0] < r.start && n[1] > r.end

			if startInside && endInside {
				shifted = append(shifted, [2]int{n[0] + r.shift, n[1] + r.shift})
				ranges = append(ranges[:i], ranges[i+1:]...)
				i--

			} else if opositeSides {
				shifted = append(shifted, [2]int{r.start + r.shift, r.end + r.shift})
				ranges = append(ranges, [2]int{r.end + 1, n[1]})
				*n = [2]int{n[0], r.start - 1}

			} else if startInside {
				shifted = append(shifted, [2]int{n[0] + r.shift, r.end + r.shift})
				*n = [2]int{r.end + 1, n[1]}

			} else if endInside {
				shifted = append(shifted, [2]int{r.start + r.shift, n[1] + r.shift})
				*n = [2]int{n[0], r.start - 1}
			}
		}
	}

	return append(shifted, ranges...)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines[0])
	maps := parseMaps(lines[3:])

	lowest := math.MaxInt
	for _, s := range seeds {
		for _, m := range maps {
			for _, r := range m {
				if s > r.start && s < r.end {
					s += r.shift
					break
				}
			}
		}

		lowest = min(lowest, s)
	}

	return lowest
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	seeds := parseSeedRanges(lines[0])
	maps := parseMaps(lines[3:])

	for _, m := range maps {
		seeds = convertRanges(seeds, m)
	}

	lowest := seeds[0][0]
	for _, s := range seeds {
		lowest = min(lowest, s[0])
	}

	return lowest
}

func main() {
	fmt.Println("--- 2023 day 05 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
