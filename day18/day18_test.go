package day18_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day18"
)

func TestPart1(t *testing.T) {
	part1 := day18.Part1("day18_test.txt")
	if part1 != 64 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	// part2 := day18.Part2("day18_test.txt")
	// if part2 != 58 {
	// 	t.Fatalf("part2 = %v", part2)
	// }
}
