package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calorieItems []int
}

func readData(filename string) []elf {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	lines := make([]int, 0)
	elves := make([]elf, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, elf{lines})
			lines = make([]int, 0)
		} else {
			i, _ := strconv.Atoi(line)
			lines = append(lines, i)
		}
	}
	elves = append(elves, elf{lines})
	return elves
}

func Calories(filename string) int {
	elves := readData(filename)

	max := 0
	for _, elf := range elves {
		sum := 0
		for _, c := range elf.calorieItems {
			sum += c
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func CaloriesTop3(filename string) int {
	elves := readData(filename)

	allCalories := make([]int, 0)

	for _, elf := range elves {
		sum := 0
		for _, calorieItem := range elf.calorieItems {
			sum += calorieItem
		}
		allCalories = append(allCalories, sum)
	}

	sort.Ints(allCalories)

	return allCalories[len(allCalories)-1] + allCalories[len(allCalories)-2] + allCalories[len(allCalories)-3]
}
