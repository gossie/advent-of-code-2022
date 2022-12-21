package day20

import (
	"bufio"
	"os"
	"strconv"
)

func readData(filename string) []*int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := make([]*int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, &n)
	}

	return numbers
}

func actualIndex(numbers []*int, pointer *int) int {
	for i, n := range numbers {
		if n == pointer {
			return i
		}
	}
	panic("did not find index")
}

func mix(initial, numbers []*int, index int) []*int {
	if index == len(numbers) {
		return numbers
	}

	n := initial[index]
	actualIndex := actualIndex(numbers, n)
	newIndex := actualIndex + *n
	if *n < 0 {
		newIndex--
	}

	if newIndex >= 0 {
		newIndex %= len(numbers)
	} else {
		for newIndex < 0 {
			newIndex = len(numbers) + newIndex
		}
	}

	if actualIndex == newIndex {
		return mix(initial, numbers, index+1)
	} else if actualIndex < newIndex {
		mixed := append(append(append(append(make([]*int, 0, len(numbers)), numbers[:actualIndex]...), numbers[actualIndex+1:newIndex+1]...), n), numbers[newIndex+1:]...)
		return mix(initial, mixed, index+1)
	} else {
		mixed := append(append(append(append(make([]*int, 0, len(numbers)), numbers[:newIndex]...), numbers[newIndex+1:actualIndex+1]...), n), numbers[actualIndex+1:]...)
		return mix(initial, mixed, index+1)
	}

}

func Part1(filename string) int {
	numbers := readData(filename)
	numbers = mix(numbers, append(make([]*int, 0, len(numbers)), numbers...), 0)
	return *numbers[1000&len(numbers)] + *numbers[2000&len(numbers)] + *numbers[3000&len(numbers)]
}
