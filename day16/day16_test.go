package day16

import "testing"

func TestPressureReleased(t *testing.T) {
	pressure := PressureReleased("day16_test.txt")
	if pressure != 1651 {
		t.Fatalf("pressure released = %v", pressure)
	}
}

func TestPressureReleasedWithElephant(t *testing.T) {
	pressure := PressureReleasedWithElephant("day16_test.txt")
	if pressure != 1707 {
		t.Fatalf("pressure released = %v", pressure)
	}
}
