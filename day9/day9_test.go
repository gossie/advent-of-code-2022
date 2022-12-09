package day9

import (
	"testing"
)

func TestVisited1(t *testing.T) {
	visited := Visited("day9_test.txt", 2)
	if visited != 13 {
		t.Fatalf("visited = %v", visited)
	}
}

func TestVisited9(t *testing.T) {
	visited := Visited("day9_test_larger.txt", 10)
	if visited != 36 {
		t.Fatalf("visited = %v", visited)
	}
}
