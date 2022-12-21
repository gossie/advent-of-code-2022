package day19_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day19"
)

func TestPart1(t *testing.T) {
	part1 := day19.Part1("day19_test.txt")
	if part1 != 33 {
		t.Fatalf("part1 = %v", part1)
	}
}
