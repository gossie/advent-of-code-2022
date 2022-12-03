package day3

import (
	"bufio"
	"os"
	"strings"
)

type rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func readData(filename string) []rucksack {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	lines := make([]rucksack, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		compartment1 := line[0 : len(line)/2]
		compartment2 := line[len(line)/2:]
		lines = append(lines, rucksack{compartment1, compartment2})
	}
	return lines
}

func PrioSum(filename string) int32 {
	rucksacks := readData(filename)
	sum := int32(0)
	for _, r := range rucksacks {
		for _, letter := range r.firstCompartment {
			if strings.ContainsRune(r.secondCompartment, letter) {
				sum += letterToPrio(letter)
				break
			}
		}
	}
	return sum
}

func BatchSum(filename string) int32 {
	rucksacks := readData(filename)
	sum := int32(0)
	for i := 0; i < len(rucksacks); i += 3 {
		r0 := rucksacks[i].firstCompartment + rucksacks[i].secondCompartment
		r1 := rucksacks[i+1].firstCompartment + rucksacks[i+1].secondCompartment
		r2 := rucksacks[i+2].firstCompartment + rucksacks[i+2].secondCompartment
		for _, letter := range r0 {
			if strings.ContainsRune(r1, letter) && strings.ContainsRune(r2, letter) {
				sum += letterToPrio(letter)
				break
			}
		}
	}
	return sum
}

func letterToPrio(letter rune) int32 {
	if letter >= 65 && letter <= 90 {
		return int32(letter - 38)
	}
	return int32(letter - 96)
}
