package day6

import (
	"bufio"
	"os"
)

func readData(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}

func Part1(filename string) int {
	input := readData(filename)
	return determineMarker(input, 4)
}

func Part2(filename string) int {
	input := readData(filename)
	return determineMarker(input, 14)
}

func determineMarker(input string, length int) int {
	for index := range input {
		if areDisjunct(input, index, length) {
			return index + length
		}
	}
	panic("found nothing")
}

func areDisjunct(input string, start, length int) bool {
	for i := start; i < start+length-1; i++ {
		for j := i + 1; j < start+length; j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
