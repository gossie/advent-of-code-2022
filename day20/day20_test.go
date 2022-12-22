package day20_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day20"
)

func TestPart1(t *testing.T) {
	part1 := day20.Part1("day20_test.txt")
	if part1 != 3 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day20.Part2("day20_test.txt")
	if part2 != 1623178306 {
		t.Fatalf("part2 = %v", part2)
	}
}
