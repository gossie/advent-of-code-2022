package day14

import "testing"

func TestSand1(t *testing.T) {
	sand := Sand1("day14_test.txt")
	if sand != 24 {
		t.Fatalf("sand = %v", sand)
	}
}

func TestSand2(t *testing.T) {
	sand := Sand2("day14_test.txt")
	if sand != 93 {
		t.Fatalf("sand = %v", sand)
	}
}
