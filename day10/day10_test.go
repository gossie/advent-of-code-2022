package day10

import "testing"

func TestSignalStrength(t *testing.T) {
	signalStrength := SignalStrength("day10_test.txt")
	if signalStrength != 13140 {
		t.Fatalf("signal strength = %v", signalStrength)
	}
}

func TestSprite(t *testing.T) {
	Sprite("day10_test.txt")
}
