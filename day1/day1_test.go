package day1

import (
	"testing"
)

func TestTask1(t *testing.T) {
	calories := Calories("day1_test.txt")
	if calories != 24000 {
		t.Fatalf("calories = %v", calories)
	}
}

func TestTask2(t *testing.T) {
	calories := CaloriesTop3("day1_test.txt")
	if calories != 45000 {
		t.Fatalf("calories = %v", calories)
	}
}
