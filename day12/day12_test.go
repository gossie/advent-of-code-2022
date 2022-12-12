package day12

import "testing"

func TestShortestClimb(t *testing.T) {
	steps := ShortestClimb("day12_test.txt")
	if steps != 31 {
		t.Fatalf("steps = %v", steps)
	}
}

func TestShortestStartingPoint(t *testing.T) {
	steps := ShortestStartingPoint("day12_test.txt")
	if steps != 29 {
		t.Fatalf("steps = %v", steps)
	}
}
