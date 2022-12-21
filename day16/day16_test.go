package day16_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day16"
)

func TestPart1(t *testing.T) {
	part1 := day16.Part1("day16_test.txt")
	if part1 != 1651 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day16.Part2("day16_test.txt")
	if part2 != 1707 {
		t.Fatalf("part2 = %v", part2)
	}
}
