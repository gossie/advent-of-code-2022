package day19_test

import (
	"testing"

	"github.com/gossie/adventofcode2022/day19"
)

func TestPart1(t *testing.T) {
	result := day19.Part1("day19_test.txt")
	if result != 33 {
		t.Fatalf("result = %v", result)
	}
}

func TestPart2(t *testing.T) {
	// result := day19.Part2("day19_test.txt")
	// if result != 62 {
	// 	t.Fatalf("result = %v", result)
	// }
}
