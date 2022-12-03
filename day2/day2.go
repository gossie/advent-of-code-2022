package day2

import (
	"bufio"
	"os"
	"strings"
)

var myChoiceToStandard = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var otherChoiceToStandard = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

type round struct {
	otherChoice string
	myChoice    string
}

func readData(filename string) []round {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	lines := make([]round, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Split(line, " ")
		lines = append(lines, round{choices[0], choices[1]})
	}
	return lines
}

func roundResult(round round) int {
	if myChoiceToStandard[round.myChoice] == otherChoiceToStandard[round.otherChoice] {
		return 3
	} else if (myChoiceToStandard[round.myChoice] == "R" && otherChoiceToStandard[round.otherChoice] == "S") || (myChoiceToStandard[round.myChoice] == "S" && otherChoiceToStandard[round.otherChoice] == "P") || (myChoiceToStandard[round.myChoice] == "P" && otherChoiceToStandard[round.otherChoice] == "R") {
		return 6
	} else {
		return 0
	}
}

func Points1(filename string) int {
	choiceToPoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	rounds := readData(filename)
	points := 0
	for _, round := range rounds {
		points += choiceToPoints[round.myChoice] + roundResult(round)
	}
	return points
}

func findChoice(round round) int {
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

func Points2(filename string) int {
	resultToPoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	rounds := readData(filename)
	points := 0
	for _, round := range rounds {
		points += resultToPoints[round.myChoice] + findChoice(round)
	}
	return points
}
