package day7

import "testing"

func TestSumSize(t *testing.T) {
	sum := SumSizes("day7_test.txt")
	if sum != 95437 {
		t.Fatalf("sum = %v", sum)
	}
}

func TestDirectoryDeleteSize(t *testing.T) {
	size := DirectoryDeleteSize("day7_test.txt")
	if size != 24933642 {
		t.Fatalf("sum = %v", size)
	}
}
