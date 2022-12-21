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
