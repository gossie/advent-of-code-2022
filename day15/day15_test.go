package day15

import "testing"

func TestNumberOfPositionsWithoutBeacon(t *testing.T) {
	positions := NumberOfPositionsWithoutBeacon("day15_test.txt", 10)
	if positions != 26 {
		t.Fatalf("positions = %v", positions)
	}
}

func TestTuneFrequency(t *testing.T) {
	tuningFrequency := TuningFrequency("day15_test.txt", 20)
	if tuningFrequency != 56000011 {
		t.Fatalf("tuning frequency = %v", tuningFrequency)
	}
}
