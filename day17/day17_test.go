package day17_test

import (
    "testing"

    "github.com/gossie/adventofcode2022/day17"
)

func TestPart1(t *testing.T) {
    part1 := day17.Part1("day17_test.txt")
    if part1 != 0 {
        t.Fatalf("part1 = %v", part1)
    }
}

func TestPart2(t *testing.T) {
    part2 := day17.Part2("day17_test.txt")
    if part2 != 0 {
        t.Fatalf("part2 = %v", part2)
    }
}
