package day1

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calorieItems []int
}

func (e elf) totalCalories() int {
	sum := 0
	for _, calorieItem := range e.calorieItems {
		sum += calorieItem
	}
	return sum
}

func readData(filename string, elves chan<- elf) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	lines := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves <- elf{lines}
			lines = make([]int, 0)
		} else {
			i, _ := strconv.Atoi(line)
			lines = append(lines, i)
		}
	}
	elves <- elf{lines}
	close(elves)
}

func Part1(filename string) int {
	elves := make(chan elf, 10)
	go readData(filename, elves)

	max := 0
	for elf := range elves {
		sum := 0
		for _, c := range elf.calorieItems {
			sum += c
		}
		max = int(math.Max(float64(max), float64(sum)))
	}
	return max
}

func Part2(filename string) int {
	elves := make(chan elf, 10)
	go readData(filename, elves)

	allCalories := make([]int, 0)

	for elf := range elves {
		allCalories = append(allCalories, elf.totalCalories())
	}

	sort.Slice(allCalories, func(i, j int) bool { return allCalories[i] > allCalories[j] })

	return allCalories[0] + allCalories[1] + allCalories[2]
}
