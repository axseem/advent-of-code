package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test2.txt
var inputTest2 string

func TestPart1(t *testing.T) {
	answer := part1(inputTest)
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(inputTest2)
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart1Input(t *testing.T) {
	answer := part1(input)
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func TestPart2Input(t *testing.T) {
	answer := part2(input)
	expected := 0
	if answer != expected {
		t.Errorf("wrong answer, got: %d, expected: %d.", answer, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
