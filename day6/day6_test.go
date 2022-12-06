package day6

import "testing"

func TestPacketMarker(t *testing.T) {
	digits := PacketMarker("day6_test.txt")
	if digits != 7 {
		t.Fatalf("digits = %v", digits)
	}
}

func TestMessageMarker(t *testing.T) {
	digits := MessageMarker("day6_test.txt")
	if digits != 19 {
		t.Fatalf("digits = %v", digits)
	}
}
