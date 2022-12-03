package day3

import (
	"testing"
)

func TestTask1(t *testing.T) {
	sum := PrioSum("day3_test.txt")
	if sum != 157 {
		t.Fatalf("sum = %v", sum)
	}
}

func TestTask2(t *testing.T) {
	sum := BatchSum("day3_test.txt")
	if sum != 70 {
		t.Fatalf("sum = %v", sum)
	}
}
