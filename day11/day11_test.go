package day11_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day11"
)

func TestPart1(t *testing.T) {
	part1 := day11.Part1("day11_test.txt")
	if part1 != 10605 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day11.Part2("day11_test.txt")
	if part2 != 2713310158 {
		t.Fatalf("part2 = %v", part2)
	}
}
