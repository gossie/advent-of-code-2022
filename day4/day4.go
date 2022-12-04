package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type assignmentRange struct {
	lowerBound int32
	upperBound int32
}

func (ar assignmentRange) subsumes(other assignmentRange) bool {
	return ar.lowerBound <= other.lowerBound && ar.upperBound >= other.upperBound
}

func (ar assignmentRange) intersectsWith(other assignmentRange) bool {
	return (ar.lowerBound >= other.lowerBound && ar.lowerBound <= other.upperBound) || (other.lowerBound >= ar.lowerBound && other.lowerBound <= ar.upperBound)
}

type pair struct {
	first  assignmentRange
	second assignmentRange
}

func (p pair) subsumes() bool {
	return p.first.subsumes(p.second) || p.second.subsumes(p.first)
}

func (p pair) intersects() bool {
	return p.first.intersectsWith(p.second)
}

func readData(filename string, pairs chan pair) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		firstRange := strings.Split(ranges[0], "-")
		firstLower, _ := strconv.Atoi(firstRange[0])
		firstUpper, _ := strconv.Atoi(firstRange[1])
		firstAssignemtRange := assignmentRange{int32(firstLower), int32(firstUpper)}

		secondRange := strings.Split(ranges[1], "-")
		secondLower, _ := strconv.Atoi(secondRange[0])
		secondUpper, _ := strconv.Atoi(secondRange[1])
		secondAssignemtRange := assignmentRange{int32(secondLower), int32(secondUpper)}

		pairs <- pair{firstAssignemtRange, secondAssignemtRange}
	}
	close(pairs)
}

func SubsumingPairs(filename string) int32 {
	pairs := make(chan pair)
	go readData(filename, pairs)

	sum := int32(0)
	for p := range pairs {
		if p.subsumes() {
			sum++
		}
	}
	return sum
}

func OverlappingPairs(filename string) int32 {
	pairs := make(chan pair)
	go readData(filename, pairs)

	sum := int32(0)
	for p := range pairs {
		if p.intersects() {
			sum++
		}
	}
	return sum
}
