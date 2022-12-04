package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calorieItems []uint32
}

func (e elf) totalCalories() uint32 {
	sum := uint32(0)
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

	lines := make([]uint32, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves <- elf{lines}
			lines = make([]uint32, 0)
		} else {
			i, _ := strconv.Atoi(line)
			lines = append(lines, uint32(i))
		}
	}
	elves <- elf{lines}
	close(elves)
}

func Calories(filename string) uint32 {
	elves := make(chan elf, 10)
	go readData(filename, elves)

	max := uint32(0)
	for elf := range elves {
		sum := uint32(0)
		for _, c := range elf.calorieItems {
			sum += c
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func CaloriesTop3(filename string) uint32 {
	elves := make(chan elf, 10)
	go readData(filename, elves)

	allCalories := make([]uint32, 0)

	for elf := range elves {
		allCalories = append(allCalories, elf.totalCalories())
	}

	sort.Slice(allCalories, func(i, j int) bool { return allCalories[i] > allCalories[j] })

	return allCalories[0] + allCalories[1] + allCalories[2]
}
