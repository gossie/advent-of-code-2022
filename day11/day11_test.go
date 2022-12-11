package day11

import "testing"

func TestMonkeyBusiness(t *testing.T) {
	monkeyBusiness := MonkeyBusiness("day11_test.txt")
	if monkeyBusiness != 10605 {
		t.Fatalf("monkey business = %v", monkeyBusiness)
	}
}

func TestMonkeyBusinessWithoutRelief(t *testing.T) {
	monkeyBusiness := MonkeyBusinessWithoutRelief("day11_test.txt")
	if monkeyBusiness != 2713310158 {
		t.Fatalf("monkey business = %v", monkeyBusiness)
	}
}
