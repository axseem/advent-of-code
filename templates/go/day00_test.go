package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test2.txt
var inputTest2 string

func TestPart1(t *testing.T) {
	answer := part1(strings.Split(inputTest, "\n"))
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(strings.Split(inputTest2, "\n"))
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart1Input(t *testing.T) {
	answer := part1(strings.Split(input, "\n"))
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart2Input(t *testing.T) {
	answer := part2(strings.Split(input, "\n"))
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := strings.Split(input, "\n")
	for i := 0; i < b.N; i++ {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := strings.Split(input, "\n")
	for i := 0; i < b.N; i++ {
		part2(lines)
	}
}
