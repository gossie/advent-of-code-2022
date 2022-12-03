package main

import (
	"fmt"

	"github.com/gossie/adventofcode2022/day1"
	"github.com/gossie/adventofcode2022/day2"
	"github.com/gossie/adventofcode2022/day3"
)

func main() {
	fmt.Println("Performing tasks of day 1")
	fmt.Println("Day 1, task 1: ", day1.Calories("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.CaloriesTop3("day1/day1.txt"))

	fmt.Println("\nPerforming tasks of day 2")
	fmt.Println("Day 2, task 1: ", day2.Points1("day2/day2.txt"))
	fmt.Println("Day 2, task 2: ", day2.Points2("day2/day2.txt"))

	fmt.Println("\nPerforming tasks of day 3")
	fmt.Println("Day 3, task 1: ", day3.PrioSum("day3/day3.txt"))
	fmt.Println("Day 3, task 2: ", day3.BatchSum("day3/day3.txt"))
}
