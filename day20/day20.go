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
	if *n == 0 {
		return mix(initial, numbers, index+1)
	}

	actualIndex := actualIndex(numbers, n)
	mixed := append(append(make([]*int, 0, len(numbers)-1), numbers[:actualIndex]...), numbers[actualIndex+1:]...)

	newIndex := actualIndex + *n

	if newIndex >= 0 {
		newIndex %= len(mixed)
	} else {
		newIndex = len(mixed) + (newIndex % len(mixed))
	}

	if actualIndex < newIndex {
		mixed = append(append(append(make([]*int, 0, len(numbers)), mixed[:newIndex]...), n), mixed[newIndex:]...)
		return mix(initial, mixed, index+1)
	} else {
		mixed = append(append(append(make([]*int, 0, len(numbers)), mixed[:newIndex]...), n), mixed[newIndex:]...)
		return mix(initial, mixed, index+1)
	}

}

func Part1(filename string) int {
	numbers := readData(filename)
	numbers = mix(numbers, append(make([]*int, 0, len(numbers)), numbers...), 0)

	for index, p := range numbers {
		if *p == 0 {
			return *numbers[(1000+index)%len(numbers)] + *numbers[(2000+index)%len(numbers)] + *numbers[(3000+index)%len(numbers)]
		}
	}
	panic("no zero found")
}

func Part2(filename string) int {
	initial := readData(filename)

	for _, p := range initial {
		*p = *p * 811589153
	}

	mixed := append(make([]*int, 0, len(initial)), initial...)
	for i := 0; i < 10; i++ {
		mixed = mix(initial, mixed, 0)
	}

	for index, p := range mixed {
		if *p == 0 {
			return *mixed[(1000+index)%len(mixed)] + *mixed[(2000+index)%len(mixed)] + *mixed[(3000+index)%len(mixed)]
		}
	}
	panic("no zero found")
}
