package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func getNumber(line *string, index int) int {
	var start int
	for i := index - 1; i >= 0; i-- {
		if (*line)[i] < '0' || (*line)[i] > '9' {
			start = i + 1
			break
		}
	}

	var number int
	newLine := []byte(*line)
	for i := start; i < len(*line) && (*line)[i] >= '0' && (*line)[i] <= '9'; i++ {
		number = number*10 + int((*line)[i]-'0')
		newLine[i] = '.'
	}
	*line = string(newLine)

	return number
}

func part1(lines []string) int {
	var adjacent = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	var sum int
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {

			c := lines[y][x]
			if (c != '.' && c < '0') || c > '9' {
				for i := 0; i < len(adjacent); i++ {
					nX := x + adjacent[i][0]
					nY := y + adjacent[i][1]
					c := lines[nY][nX]
					if c >= '0' && c <= '9' {
						sum += getNumber(&lines[nY], nX)
					}
				}
			}

		}
	}

	return sum
}

func part2(lines []string) int {
	var adjacent = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	var sum int
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {

			c := lines[y][x]
			if (c != '.' && c < '0') || c > '9' {
				gearRatio := 1
				var amount int
				for i := 0; i < len(adjacent); i++ {
					if amount > 2 {
						break
					}

					nX := x + adjacent[i][0]
					nY := y + adjacent[i][1]
					c := lines[nY][nX]
					if c >= '0' && c <= '9' {
						amount++
						gearRatio *= getNumber(&lines[nY], nX)
					}
				}

				if amount == 2 {
					sum += gearRatio
				}
			}

		}
	}

	return sum
}

func main() {
	lines := strings.Split(input, "\n")

	fmt.Println("--- 2023 day 03 answer ---")
	fmt.Println("part 1:\t", part1(lines))
	fmt.Println("part 2:\t", part2(lines))
}
