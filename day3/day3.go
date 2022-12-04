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

func readData(filename string, rucksacks chan rucksack) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		compartment1 := line[0 : len(line)/2]
		compartment2 := line[len(line)/2:]
		rucksacks <- rucksack{compartment1, compartment2}
	}
	close(rucksacks)
}

func PrioSum(filename string) int32 {
	rucksacks := make(chan rucksack, 1)
	go readData(filename, rucksacks)
	sum := int32(0)
	for r := range rucksacks {
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
	rucksacks := make(chan rucksack, 3)
	go readData(filename, rucksacks)
	sum := int32(0)
	for ruck0 := range rucksacks {
		ruck1 := <-rucksacks
		ruck2 := <-rucksacks

		r0 := ruck0.firstCompartment + ruck0.secondCompartment
		r1 := ruck1.firstCompartment + ruck1.secondCompartment
		r2 := ruck2.firstCompartment + ruck2.secondCompartment
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
