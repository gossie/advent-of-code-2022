package day13_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day13"
)

func TestPart1(t *testing.T) {
	part1 := day13.Part1("day13_test.txt")
	if part1 != 13 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day13.Part2("day13_test.txt")
	if part2 != 140 {
		t.Fatalf("part2 = %v", part2)
	}
}
