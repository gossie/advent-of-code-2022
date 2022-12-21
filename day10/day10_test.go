package day10_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day10"
)

func TestPart1(t *testing.T) {
	part1 := day10.Part1("day10_test.txt")
	if part1 != 13140 {
		t.Fatalf("part1 = %v", part1)
	}
}
