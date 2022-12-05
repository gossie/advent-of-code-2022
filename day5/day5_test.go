package day5

import (
	"testing"
)

func TestTask1(t *testing.T) {
	crates := Crates9000("day5_test.txt")
	if crates != "CMZ" {
		t.Fatalf("crates = %v", crates)

	}
}

func TestTask2(t *testing.T) {
	crates := Crates9001("day5_test.txt")
	if crates != "MCD" {
		t.Fatalf("crates = %v", crates)

	}
}
