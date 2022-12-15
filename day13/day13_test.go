package day13

import "testing"

func TestCorrectOrder(t *testing.T) {
	correct := CorrectOrder("day13_test.txt")
	if correct != 13 {
		t.Fatalf("correct = %v", correct)
	}
}

func TestSort(t *testing.T) {
	decoderKey := DecoderKey("day13_test.txt")
	if decoderKey != 140 {
		t.Fatalf("correct = %v", decoderKey)
	}
}
