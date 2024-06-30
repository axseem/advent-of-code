package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func (p Point) Moved(d Point) Point {
	return Point{p.x + d.x, p.y + d.y}
}

func (p Point) Connected(b byte) []Point {
	switch b {
	case '|':
		return []Point{p.Moved(Point{0, -1}), p.Moved(Point{0, 1})}
	case '-':
		return []Point{p.Moved(Point{1, 0}), p.Moved(Point{-1, 0})}
	case 'L':
		return []Point{p.Moved(Point{0, -1}), p.Moved(Point{1, 0})}
	case 'J':
		return []Point{p.Moved(Point{-1, 0}), p.Moved(Point{0, -1})}
	case '7':
		return []Point{p.Moved(Point{0, 1}), p.Moved(Point{-1, 0})}
	case 'F':
		return []Point{p.Moved(Point{1, 0}), p.Moved(Point{0, 1})}
	}
	panic("unreachable")
}

func (p Point) IsInBounds(bounds Point) bool {
	return p.x >= 0 && p.x < bounds.x && p.y >= 0 && p.y < bounds.y
}

func FindStart(lines *[]string) Point {
	for row, line := range *lines {
		for col, r := range line {
			if r == 'S' {
				return Point{col, row}
			}
		}
	}
	panic("unreachable")
}

func NEIGHBOR_SHIFT() [4]Point {
	return [4]Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	bounds := Point{len(lines[0]), len(lines)}
	start := FindStart(&lines)
	current := start
	steps := 1

	for _, shift := range NEIGHBOR_SHIFT() {
		neighbor := start.Moved(shift)
		if neighbor.IsInBounds(bounds) && lines[neighbor.y][neighbor.x] != '.' {
			if slices.Contains(neighbor.Connected(lines[neighbor.y][neighbor.x]), start) {
				current = neighbor
			}
		}
	}

	if current == start {
		panic("unreachable")
	}

	prev := start
	for current != start {
		neighbors := current.Connected(lines[current.y][current.x])
		neighbors = slices.DeleteFunc(neighbors, func(e Point) bool { return e == prev })

		if len(neighbors) != 1 {
			panic("unreachable")
		}

		prev = current
		current = neighbors[0]
		steps++
	}

	return steps / 2
}

func part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--- 2023 day 10 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
