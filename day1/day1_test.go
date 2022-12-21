package day1_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day1"
)

func TestPart1(t *testing.T) {
	part1 := day1.Part1("day1_test.txt")
	if part1 != 24000 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day1.Part2("day1_test.txt")
	if part2 != 45000 {
		t.Fatalf("part2 = %v", part2)
	}
}
