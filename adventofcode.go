package main

import (
	"fmt"
	"time"

	"github.com/gossie/adventofcode2022/day1"
	"github.com/gossie/adventofcode2022/day10"
	"github.com/gossie/adventofcode2022/day11"
	"github.com/gossie/adventofcode2022/day12"
	"github.com/gossie/adventofcode2022/day13"
	"github.com/gossie/adventofcode2022/day14"
	"github.com/gossie/adventofcode2022/day15"
	"github.com/gossie/adventofcode2022/day16"
	"github.com/gossie/adventofcode2022/day17"
	"github.com/gossie/adventofcode2022/day18"
	"github.com/gossie/adventofcode2022/day19"
	"github.com/gossie/adventofcode2022/day2"
	"github.com/gossie/adventofcode2022/day20"
	"github.com/gossie/adventofcode2022/day3"
	"github.com/gossie/adventofcode2022/day4"
	"github.com/gossie/adventofcode2022/day5"
	"github.com/gossie/adventofcode2022/day6"
	"github.com/gossie/adventofcode2022/day7"
	"github.com/gossie/adventofcode2022/day8"
	"github.com/gossie/adventofcode2022/day9"
)

func main() {

	fmt.Println("\nPerforming tasks of day1")
	startday1Part1 := time.Now().UnixMilli()
	day1Part1 := day1.Part1("day1/day1.txt")
	fmt.Println("day1, task 1: ", day1Part1, ", took", (time.Now().UnixMilli() - startday1Part1), "ms")
	startday1Part2 := time.Now().UnixMilli()
	day1Part2 := day1.Part2("day1/day1.txt")
	fmt.Println("day1, task 2: ", day1Part2, ", took", (time.Now().UnixMilli() - startday1Part2), "ms")

	fmt.Println("\nPerforming tasks of day2")
	startday2Part1 := time.Now().UnixMilli()
	day2Part1 := day2.Part1("day2/day2.txt")
	fmt.Println("day2, task 1: ", day2Part1, ", took", (time.Now().UnixMilli() - startday2Part1), "ms")
	startday2Part2 := time.Now().UnixMilli()
	day2Part2 := day2.Part2("day2/day2.txt")
	fmt.Println("day2, task 2: ", day2Part2, ", took", (time.Now().UnixMilli() - startday2Part2), "ms")

	fmt.Println("\nPerforming tasks of day3")
	startday3Part1 := time.Now().UnixMilli()
	day3Part1 := day3.Part1("day3/day3.txt")
	fmt.Println("day3, task 1: ", day3Part1, ", took", (time.Now().UnixMilli() - startday3Part1), "ms")
	startday3Part2 := time.Now().UnixMilli()
	day3Part2 := day3.Part2("day3/day3.txt")
	fmt.Println("day3, task 2: ", day3Part2, ", took", (time.Now().UnixMilli() - startday3Part2), "ms")

	fmt.Println("\nPerforming tasks of day4")
	startday4Part1 := time.Now().UnixMilli()
	day4Part1 := day4.Part1("day4/day4.txt")
	fmt.Println("day4, task 1: ", day4Part1, ", took", (time.Now().UnixMilli() - startday4Part1), "ms")
	startday4Part2 := time.Now().UnixMilli()
	day4Part2 := day4.Part2("day4/day4.txt")
	fmt.Println("day4, task 2: ", day4Part2, ", took", (time.Now().UnixMilli() - startday4Part2), "ms")

	fmt.Println("\nPerforming tasks of day5")
	startday5Part1 := time.Now().UnixMilli()
	day5Part1 := day5.Part1("day5/day5.txt")
	fmt.Println("day5, task 1: ", day5Part1, ", took", (time.Now().UnixMilli() - startday5Part1), "ms")
	startday5Part2 := time.Now().UnixMilli()
	day5Part2 := day5.Part2("day5/day5.txt")
	fmt.Println("day5, task 2: ", day5Part2, ", took", (time.Now().UnixMilli() - startday5Part2), "ms")

	fmt.Println("\nPerforming tasks of day6")
	startday6Part1 := time.Now().UnixMilli()
	day6Part1 := day6.Part1("day6/day6.txt")
	fmt.Println("day6, task 1: ", day6Part1, ", took", (time.Now().UnixMilli() - startday6Part1), "ms")
	startday6Part2 := time.Now().UnixMilli()
	day6Part2 := day6.Part2("day6/day6.txt")
	fmt.Println("day6, task 2: ", day6Part2, ", took", (time.Now().UnixMilli() - startday6Part2), "ms")

	fmt.Println("\nPerforming tasks of day7")
	startday7Part1 := time.Now().UnixMilli()
	day7Part1 := day7.Part1("day7/day7.txt")
	fmt.Println("day7, task 1: ", day7Part1, ", took", (time.Now().UnixMilli() - startday7Part1), "ms")
	startday7Part2 := time.Now().UnixMilli()
	day7Part2 := day7.Part2("day7/day7.txt")
	fmt.Println("day7, task 2: ", day7Part2, ", took", (time.Now().UnixMilli() - startday7Part2), "ms")

	fmt.Println("\nPerforming tasks of day8")
	startday8Part1 := time.Now().UnixMilli()
	day8Part1 := day8.Part1("day8/day8.txt")
	fmt.Println("day8, task 1: ", day8Part1, ", took", (time.Now().UnixMilli() - startday8Part1), "ms")
	startday8Part2 := time.Now().UnixMilli()
	day8Part2 := day8.Part2("day8/day8.txt")
	fmt.Println("day8, task 2: ", day8Part2, ", took", (time.Now().UnixMilli() - startday8Part2), "ms")

	fmt.Println("\nPerforming tasks of day9")
	startday9Part1 := time.Now().UnixMilli()
	day9Part1 := day9.Visited("day9/day9.txt", 2)
	fmt.Println("day9, task 1: ", day9Part1, ", took", (time.Now().UnixMilli() - startday9Part1), "ms")
	startday9Part2 := time.Now().UnixMilli()
	day9Part2 := day9.Visited("day9/day9.txt", 10)
	fmt.Println("day9, task 2: ", day9Part2, ", took", (time.Now().UnixMilli() - startday9Part2), "ms")

	fmt.Println("\nPerforming tasks of day10")
	startday10Part1 := time.Now().UnixMilli()
	day10Part1 := day10.Part1("day10/day10.txt")
	fmt.Println("day10, task 1: ", day10Part1, ", took", (time.Now().UnixMilli() - startday10Part1), "ms")
	startday10Part2 := time.Now().UnixMilli()
	day10.Part2("day10/day10.txt")
	fmt.Println("day10, task 2 took", (time.Now().UnixMilli() - startday10Part2), "ms")

	fmt.Println("\nPerforming tasks of day11")
	startday11Part1 := time.Now().UnixMilli()
	day11Part1 := day11.Part1("day11/day11.txt")
	fmt.Println("day11, task 1: ", day11Part1, ", took", (time.Now().UnixMilli() - startday11Part1), "ms")
	startday11Part2 := time.Now().UnixMilli()
	day11Part2 := day11.Part2("day11/day11.txt")
	fmt.Println("day11, task 2: ", day11Part2, ", took", (time.Now().UnixMilli() - startday11Part2), "ms")

	fmt.Println("\nPerforming tasks of day12")
	startday12Part1 := time.Now().UnixMilli()
	day12Part1 := day12.Part1("day12/day12.txt")
	fmt.Println("day12, task 1: ", day12Part1, ", took", (time.Now().UnixMilli() - startday12Part1), "ms")
	startday12Part2 := time.Now().UnixMilli()
	day12Part2 := day12.Part2("day12/day12.txt")
	fmt.Println("day12, task 2: ", day12Part2, ", took", (time.Now().UnixMilli() - startday12Part2), "ms")

	fmt.Println("\nPerforming tasks of day13")
	startday13Part1 := time.Now().UnixMilli()
	day13Part1 := day13.Part1("day13/day13.txt")
	fmt.Println("day13, task 1: ", day13Part1, ", took", (time.Now().UnixMilli() - startday13Part1), "ms")
	startday13Part2 := time.Now().UnixMilli()
	day13Part2 := day13.Part2("day13/day13.txt")
	fmt.Println("day13, task 2: ", day13Part2, ", took", (time.Now().UnixMilli() - startday13Part2), "ms")

	fmt.Println("\nPerforming tasks of day14")
	startday14Part1 := time.Now().UnixMilli()
	day14Part1 := day14.Part1("day14/day14.txt")
	fmt.Println("day14, task 1: ", day14Part1, ", took", (time.Now().UnixMilli() - startday14Part1), "ms")
	startday14Part2 := time.Now().UnixMilli()
	day14Part2 := day14.Part2("day14/day14.txt")
	fmt.Println("day14, task 2: ", day14Part2, ", took", (time.Now().UnixMilli() - startday14Part2), "ms")

	fmt.Println("\nPerforming tasks of day15")
	startday15Part1 := time.Now().UnixMilli()
	day15Part1 := day15.Part1("day15/day15.txt", 2000000)
	fmt.Println("day15, task 1: ", day15Part1, ", took", (time.Now().UnixMilli() - startday15Part1), "ms")
	startday15Part2 := time.Now().UnixMilli()
	day15Part2 := day15.Part2("day15/day15.txt", 4000000)
	fmt.Println("day15, task 2: ", day15Part2, ", took", (time.Now().UnixMilli() - startday15Part2), "ms")

	fmt.Println("\nPerforming tasks of day16")
	startday16Part1 := time.Now().UnixMilli()
	day16Part1 := day16.Part1("day16/day16.txt")
	fmt.Println("day16, task 1: ", day16Part1, ", took", (time.Now().UnixMilli() - startday16Part1), "ms")
	startday16Part2 := time.Now().UnixMilli()
	day16Part2 := day16.Part2("day16/day16.txt")
	fmt.Println("day16, task 2: ", day16Part2, ", took", (time.Now().UnixMilli() - startday16Part2), "ms")

	fmt.Println("\nPerforming tasks of day17")
	startday17Part1 := time.Now().UnixMilli()
	day17Part1 := day17.Part1("day17/day17.txt")
	fmt.Println("day17, task 1: ", day17Part1, ", took", (time.Now().UnixMilli() - startday17Part1), "ms")
	startday17Part2 := time.Now().UnixMilli()
	day17Part2 := day17.Part2("day17/day17.txt")
	fmt.Println("day17, task 2: ", day17Part2, ", took", (time.Now().UnixMilli() - startday17Part2), "ms")

	fmt.Println("\nPerforming tasks of day18")
	startday18Part1 := time.Now().UnixMilli()
	day18Part1 := day18.Part1("day18/day18.txt")
	fmt.Println("day18, task 1: ", day18Part1, ", took", (time.Now().UnixMilli() - startday18Part1), "ms")
	// startday18Part2 := time.Now().UnixMilli()
	// day18Part2 := day18.Part2("day18/day18.txt")
	// fmt.Println("day18, task 2: ", day18Part2, ", took", (time.Now().UnixMilli() - startday18Part2), "ms")

	fmt.Println("\nPerforming tasks of day19")
	startday19Part1 := time.Now().UnixMilli()
	day19Part1 := day19.Part1("day19/day19.txt")
	fmt.Println("day19, task 1: ", day19Part1, ", took", (time.Now().UnixMilli() - startday19Part1), "ms")
	startday19Part2 := time.Now().UnixMilli()
	day19Part2 := day19.Part2("day19/day19.txt")
	fmt.Println("day19, task 2: ", day19Part2, ", took", (time.Now().UnixMilli() - startday19Part2), "ms")

	fmt.Println("\nPerforming tasks of day20")
	startday20Part1 := time.Now().UnixMilli()
	day20Part1 := day20.Part1("day20/day20.txt")
	fmt.Println("day20, task 1: ", day20Part1, ", took", (time.Now().UnixMilli() - startday20Part1), "ms")
	startday20Part2 := time.Now().UnixMilli()
	day20Part2 := day20.Part2("day20/day20.txt")
	fmt.Println("day20, task 2: ", day20Part2, ", took", (time.Now().UnixMilli() - startday20Part2), "ms")
}
