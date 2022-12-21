package day9_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day9"
)

func TestPart1(t *testing.T) {
	part1 := day9.Visited("day9_test.txt", 2)
	if part1 != 13 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day9.Visited("day9_test_larger.txt", 10)
	if part2 != 36 {
		t.Fatalf("part2 = %v", part2)
	}
}
