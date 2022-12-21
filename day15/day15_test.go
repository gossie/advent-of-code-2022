package day15_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day15"
)

func TestPart1(t *testing.T) {
	part1 := day15.Part1("day15_test.txt", 10)
	if part1 != 26 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day15.Part2("day15_test.txt", 20)
	if part2 != 56000011 {
		t.Fatalf("part2 = %v", part2)
	}
}
