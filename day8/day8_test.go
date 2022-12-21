package day8_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day8"
)

func TestPart1(t *testing.T) {
	part1 := day8.Part1("day8_test.txt")
	if part1 != 21 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day8.Part2("day8_test.txt")
	if part2 != 8 {
		t.Fatalf("part2 = %v", part2)
	}
}
