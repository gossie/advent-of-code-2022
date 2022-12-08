package day8

import "testing"

func TestVisibleTrees(t *testing.T) {
	visibleTrees := VisibleTrees("day8_test.txt")
	if visibleTrees != 21 {
		t.Fatalf("visibleTrees = %v", visibleTrees)
	}
}

func TestScenicScore(t *testing.T) {
	scenicScore := ScenicScore("day8_test.txt")
	if scenicScore != 8 {
		t.Fatalf("scenicScore = %v", scenicScore)
	}
}
