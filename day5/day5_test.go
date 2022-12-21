package day5_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day5"
)

func TestPart1(t *testing.T) {
	part1 := day5.Part1("day5_test.txt")
	if part1 != "CMZ" {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day5.Part2("day5_test.txt")
	if part2 != "MCD" {
		t.Fatalf("part2 = %v", part2)
	}
}
