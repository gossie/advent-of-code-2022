package day13

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
)

type pair struct {
	left, right []any
	index       int
}

func readData(filename string, pairs chan<- pair) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	pairIndex := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			left := make([]any, 0)
			right := make([]any, 0)
			err := json.Unmarshal([]byte(line), &left)
			if err != nil {
				panic("could not parse left")
			}
			scanner.Scan()
			line = scanner.Text()
			err = json.Unmarshal([]byte(line), &right)
			if err != nil {
				panic("could not parse left")
			}

			pairs <- pair{left, right, pairIndex}
			pairIndex++
		}
	}

	close(pairs)
}

func compareNumbers(left, right int) int {
	return left - right
}

func compareLists(left, right []any) int {
	for leftIndex, l := range left {
		lIsAList := false
		rIsAList := false

		if leftIndex >= len(right) {
			return 1
		}

		var ll []any
		ln, ok := l.(float64)
		if !ok {
			lIsAList = true
			x, ok := l.([]any)
			if !ok {
				panic("l ist not a list")
			}
			ll = x
		}

		r := right[leftIndex]

		var rl []any
		rn, ok := r.(float64)
		if !ok {
			rIsAList = true
			x, ok := r.([]any)
			if !ok {
				panic("r ist not a list")
			}
			rl = x
		}

		var comparisonResult int
		switch {
		case lIsAList && rIsAList:
			comparisonResult = compareLists(ll, rl)
		case !lIsAList && !rIsAList:
			comparisonResult = compareNumbers(int(ln), int(rn))
		case lIsAList && !rIsAList:
			comparisonResult = compareLists(ll, []any{rn})
		case !lIsAList && rIsAList:
			comparisonResult = compareLists([]any{ln}, rl)
		}

		if comparisonResult == 0 {
			continue
		}
		return comparisonResult
	}
	return len(left) - len(right)
}

func CorrectOrder(filename string) int {
	pairs := make(chan pair, 10)

	go readData(filename, pairs)

	correct := 0

	for p := range pairs {
		if compareLists(p.left, p.right) <= 0 {
			correct += p.index
		}
	}

	return correct
}

func DecoderKey(filename string) int {
	pairs := make(chan pair, 10)

	go readData(filename, pairs)

	dividerPacket1 := []any{[]any{float64(2)}}
	dividerPacket2 := []any{[]any{float64(6)}}

	allPackets := make([]any, 0)
	allPackets = append(allPackets, dividerPacket1)
	allPackets = append(allPackets, dividerPacket2)

	for p := range pairs {
		allPackets = append(allPackets, p.left)
		allPackets = append(allPackets, p.right)
	}

	sort.Slice(allPackets, func(i, j int) bool {
		x, ok1 := allPackets[i].([]any)
		y, ok2 := allPackets[j].([]any)
		if !ok1 || !ok2 {
			panic("expecting lists")
		}
		return compareLists(x, y) < 0
	})

	index1 := -1
	index2 := -1

	for i := 0; i < len(allPackets); i++ {
		p := allPackets[i].([]any)

		if compareLists(p, dividerPacket1) == 0 {
			index1 = i + 1
		}
		if compareLists(p, dividerPacket2) == 0 {
			index2 = i + 1
		}
	}

	return index1 * index2
}
