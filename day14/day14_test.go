package day14_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day14"
)

func TestPart1(t *testing.T) {
	part1 := day14.Part1("day14_test.txt")
	if part1 != 24 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day14.Part2("day14_test.txt")
	if part2 != 93 {
		t.Fatalf("part2 = %v", part2)
	}
}
