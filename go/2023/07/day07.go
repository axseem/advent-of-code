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

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

type hand struct {
	cards [5]int
	hType handType
	bid   int
}

func cardValue(card byte) int {
	if card >= '2' && card <= '9' {
		return int(card - 50)
	}

	switch card {
	case 'T':
		return 8
	case 'J':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}

	return 0
}

func cardValueJoker(card byte) int {
	if card >= '2' && card <= '9' {
		return int(card - 49)
	}

	switch card {
	case 'J':
		return 0
	case 'T':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}

	return 0
}

func getHandType(hand [5]int) handType {
	amount := make(map[int]int)
	for _, v := range hand {
		amount[v]++
	}

	var firstMax, secondMax int
	for _, v := range amount {
		if v > firstMax {
			secondMax = firstMax
			firstMax = v
		} else if v > secondMax {
			secondMax = v
		}
	}

	switch {
	case firstMax == 5:
		return fiveKind

	case firstMax == 4:
		return fourKind

	case firstMax == 3 && secondMax == 2:
		return fullHouse

	case firstMax == 3:
		return threeKind

	case firstMax == 2 && secondMax == 2:
		return twoPair

	case firstMax == 2:
		return onePair

	default:
		return highCard
	}
}

func getHandTypeJoker(hand [5]int) handType {
	amount := make(map[int]int)
	for _, v := range hand {
		amount[v]++
	}

	var firstMax, secondMax int
	for k, v := range amount {
		if k != 0 {
			if v > firstMax {
				secondMax = firstMax
				firstMax = v
			} else if v > secondMax {
				secondMax = v
			}
		}
	}

	firstMax += amount[0]

	switch {
	case firstMax == 5:
		return fiveKind

	case firstMax == 4:
		return fourKind

	case firstMax == 3 && secondMax == 2:
		return fullHouse

	case firstMax == 3:
		return threeKind

	case firstMax == 2 && secondMax == 2:
		return twoPair

	case firstMax == 2:
		return onePair

	default:
		return highCard
	}
}

func parseSet(input string) []hand {
	lines := strings.Split(input, "\n")
	set := make([]hand, len(lines))

	for i, line := range lines {
		set[i].cards = [5]int{
			cardValue(line[0]),
			cardValue(line[1]),
			cardValue(line[2]),
			cardValue(line[3]),
			cardValue(line[4]),
		}
		bid, _ := strconv.Atoi(line[6:])
		set[i].hType = getHandType(set[i].cards)
		set[i].bid = bid
	}

	return set
}

func parseSetJoker(input string) []hand {
	lines := strings.Split(input, "\n")
	set := make([]hand, len(lines))

	for i, line := range lines {
		set[i].cards = [5]int{
			cardValueJoker(line[0]),
			cardValueJoker(line[1]),
			cardValueJoker(line[2]),
			cardValueJoker(line[3]),
			cardValueJoker(line[4]),
		}
		bid, _ := strconv.Atoi(line[6:])
		set[i].hType = getHandTypeJoker(set[i].cards)
		set[i].bid = bid
	}

	return set
}

func part1(input string) int {
	set := parseSet(input)

	sort.Slice(set, func(i, j int) bool {
		if set[i].hType != set[j].hType {
			return set[i].hType < set[j].hType
		}
		for l := 0; l < len(set[i].cards); l++ {
			if set[i].cards[l] != set[j].cards[l] {
				return set[i].cards[l] < set[j].cards[l]
			}
		}
		return false
	})

	var sum int
	for i, part := range set {
		sum += part.bid * (i + 1)
	}

	return sum
}

func part2(input string) int {
	set := parseSetJoker(input)

	sort.Slice(set, func(i, j int) bool {
		if set[i].hType != set[j].hType {
			return set[i].hType < set[j].hType
		}
		for l := 0; l < len(set[i].cards); l++ {
			if set[i].cards[l] != set[j].cards[l] {
				return set[i].cards[l] < set[j].cards[l]
			}
		}
		return false
	})

	var sum int
	for i, part := range set {
		sum += part.bid * (i + 1)
	}

	return sum
}

func main() {
	fmt.Println("--- 2023 day 07 answer ---")
	fmt.Println("part 1:\t", part1(input))
	fmt.Println("part 2:\t", part2(input))
}
