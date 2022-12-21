package day2

import (
	"bufio"
	"os"
	"strings"
)

var myChoiceNormalized = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var otherChoiceNormalized = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

type round struct {
	otherChoice string
	myChoice    string
}

func readData(filename string, rounds chan<- round) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Split(line, " ")
		rounds <- round{choices[0], choices[1]}
	}
	close(rounds)
}

func roundResult(round round) uint32 {
	if myChoiceNormalized[round.myChoice] == otherChoiceNormalized[round.otherChoice] {
		return 3
	} else if (myChoiceNormalized[round.myChoice] == "R" && otherChoiceNormalized[round.otherChoice] == "S") || (myChoiceNormalized[round.myChoice] == "S" && otherChoiceNormalized[round.otherChoice] == "P") || (myChoiceNormalized[round.myChoice] == "P" && otherChoiceNormalized[round.otherChoice] == "R") {
		return 6
	} else {
		return 0
	}
}

func Part1(filename string) uint32 {
	choiceToPoints := map[string]uint32{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	rounds := make(chan round, 10)
	go readData(filename, rounds)
	points := uint32(0)
	for round := range rounds {
		points += choiceToPoints[round.myChoice] + roundResult(round)
	}
	return points
}

func findChoice(round round) uint32 {
	if round.otherChoice == "A" {
		if round.myChoice == "X" {
			return 3
		} else if round.myChoice == "Y" {
			return 1
		} else {
			return 2
		}
	} else if round.otherChoice == "B" {
		if round.myChoice == "X" {
			return 1
		} else if round.myChoice == "Y" {
			return 2
		} else {
			return 3
		}
	} else {
		if round.myChoice == "X" {
			return 2
		} else if round.myChoice == "Y" {
			return 3
		} else {
			return 1
		}
	}
}

func Part2(filename string) uint32 {
	resultToPoints := map[string]uint32{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	rounds := make(chan round, 10)
	go readData(filename, rounds)
	points := uint32(0)
	for round := range rounds {
		points += resultToPoints[round.myChoice] + findChoice(round)
	}
	return points
}
