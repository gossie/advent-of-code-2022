package day2

import (
	"testing"
)

func TestTask1(t *testing.T) {
	score := Points1("day2_test.txt")
	if score != 15 {
		t.Fatalf("score = %v", score)
	}
}

func TestTask2(t *testing.T) {
	score := Points2("day2_test.txt")
	if score != 12 {
		t.Fatalf("score = %v", score)
	}
}
