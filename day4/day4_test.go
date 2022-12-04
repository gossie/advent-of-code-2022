package day4

import (
	"testing"
)

func TestTask1(t *testing.T) {
	pairs := SubsumingPairs("day4_test.txt")
	if pairs != 2 {
		t.Fatalf("pairs = %v", pairs)

	}
}

func TestTask2(t *testing.T) {
	pairs := OverlappingPairs("day4_test.txt")
	if pairs != 4 {
		t.Fatalf("pairs = %v", pairs)

	}
}
