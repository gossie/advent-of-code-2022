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

func PacketMarker(filename string) int {
	input := readData(filename)

	length := 4
	for index := range input {
		if areDisjunct(input, index, length) {
			return index + length
		}
	}
	panic("found nothing")
}

func MessageMarker(filename string) int {
	input := readData(filename)

	length := 14
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
