package day21_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day21"
)

func TestPart1(t *testing.T) {
	part1 := day21.Part1("day21_test.txt")
	if part1 != 152 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day21.Part2("day21_test.txt")
	if part2 != 301 {
		t.Fatalf("part2 = %v", part2)
	}
}
