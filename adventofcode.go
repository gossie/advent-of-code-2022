package main

import (
	"fmt"
	"time"

	"github.com/gossie/adventofcode2022/day19"
)

func main() {
	// fmt.Println("Performing tasks of day 1")
	// fmt.Println("Day 1, task 1: ", day1.Calories("day1/day1.txt"))
	// fmt.Println("Day 1, task 2: ", day1.CaloriesTop3("day1/day1.txt"))

	// fmt.Println("\nPerforming tasks of day 2")
	// fmt.Println("Day 2, task 1: ", day2.Points1("day2/day2.txt"))
	// fmt.Println("Day 2, task 2: ", day2.Points2("day2/day2.txt"))

	// fmt.Println("\nPerforming tasks of day 3")
	// fmt.Println("Day 3, task 1: ", day3.PrioSum("day3/day3.txt"))
	// fmt.Println("Day 3, task 2: ", day3.BatchSum("day3/day3.txt"))

	// fmt.Println("\nPerforming tasks of day 4")
	// fmt.Println("Day 4, task 1: ", day4.SubsumingPairs("day4/day4.txt"))
	// fmt.Println("Day 4, task 2: ", day4.OverlappingPairs("day4/day4.txt"))

	// fmt.Println("\nPerforming tasks of day 5")
	// fmt.Println("Day 5, task 1: ", day5.Crates9000("day5/day5.txt"))
	// fmt.Println("Day 5, task 2: ", day5.Crates9001("day5/day5.txt"))

	// fmt.Println("\nPerforming tasks of day 6")
	// fmt.Println("Day 6, task 1: ", day6.PacketMarker("day6/day6.txt"))
	// fmt.Println("Day 6, task 2: ", day6.MessageMarker("day6/day6.txt"))

	// fmt.Println("\nPerforming tasks of day 7")
	// fmt.Println("Day 7, task 1: ", day7.SumSizes("day7/day7.txt"))
	// fmt.Println("Day 7, task 2: ", day7.DirectoryDeleteSize("day7/day7.txt"))

	// fmt.Println("\nPerforming tasks of day 8")
	// fmt.Println("Day 8, task 1: ", day8.VisibleTrees("day8/day8.txt"))
	// fmt.Println("Day 8, task 2: ", day8.ScenicScore("day8/day8.txt"))

	// fmt.Println("\nPerforming tasks of day 9")
	// fmt.Println("Day 9, task 1: ", day9.Visited("day9/day9.txt", 2))
	// fmt.Println("Day 9, task 2: ", day9.Visited("day9/day9.txt", 10))

	// fmt.Println("\nPerforming tasks of day 10")
	// fmt.Println("Day 10, task 1: ", day10.SignalStrength("day10/day10.txt"))
	// fmt.Println("Day 10, task 2: ")
	// day10.Sprite("day10/day10.txt")

	// fmt.Println("\nPerforming tasks of day 11")
	// fmt.Println("Day 11, task 1: ", day11.MonkeyBusiness("day11/day11.txt"))
	// fmt.Println("Day 11, task 2: ", day11.MonkeyBusinessWithoutRelief("day11/day11.txt"))

	// fmt.Println("\nPerforming tasks of day 12")
	// fmt.Println("Day 12, task 1: ", day12.ShortestClimb("day12/day12.txt"))
	// fmt.Println("Day 12, task 2: ", day12.ShortestStartingPoint("day12/day12.txt"))

	// fmt.Println("\nPerforming tasks of day 13")
	// fmt.Println("Day 13, task 1: ", day13.CorrectOrder("day13/day13.txt"))
	// fmt.Println("Day 13, task 1: ", day13.DecoderKey("day13/day13.txt"))

	// fmt.Println("\nPerforming tasks of day 14")
	// fmt.Println("Day 14, task 1: ", day14.Sand1("day14/day14.txt"))
	// fmt.Println("Day 14, task 1: ", day14.Sand2("day14/day14.txt"))

	// fmt.Println("\nPerforming tasks of day 15")
	// fmt.Println("Day 15, task 1: ", day15.NumberOfPositionsWithoutBeacon("day15/day15.txt", 2000000))
	// fmt.Println("Day 15, task 1: ", day15.TuningFrequency("day15/day15.txt", 4000000))

	// fmt.Println("\nPerforming tasks of day 16")
	// fmt.Println("Day 16, task 1: ", day16.PressureReleased("day16/day16.txt"))
	// fmt.Println("Day 16, task 2: ", day16.PressureReleasedWithElephant("day16/day16.txt"))

	// fmt.Println("\nPerforming tasks of day 18")
	// fmt.Println("Day 18, task 1: ", day18.Part1("day18/day18.txt"))

	fmt.Println("\nPerforming tasks of day 19")
	startDay19Part1 := time.Now().UnixMilli()
	day19Part1 := day19.Part1("day19/day19_test.txt")
	fmt.Println("Day 19, task 1: ", day19Part1, ", took", (time.Now().UnixMilli() - startDay19Part1), "ms")
	// startDay19Part2 := time.Now().UnixMilli()
	// day19Part2 := day19.Part2("day19/day19.txt")
	// fmt.Println("Day 19, task 2: ", day19Part2, ", took", (time.Now().UnixMilli() - startDay19Part2), "ms")

}
