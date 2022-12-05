package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/gossie/adventofcode2022/util"
)

type instruction struct {
	from   int
	to     int
	amount int
}

func lineLength(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return len(scanner.Text())
}

func readData(filename string, stacks chan<- []util.Stack[string], instructions chan<- instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	numberOfStacks := (lineLength(filename) + 1) / 4
	importedStacks := make([]util.Stack[string], numberOfStacks)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			instructionComponents := strings.Split(line, " ")
			from, _ := strconv.Atoi(instructionComponents[3])
			to, _ := strconv.Atoi(instructionComponents[5])
			amount, _ := strconv.Atoi(instructionComponents[1])
			instructions <- instruction{from - 1, to - 1, amount}
		} else if line != "" {
			for i := 0; i < numberOfStacks; i++ {
				start := (i * 4) + 1
				end := start + 1
				cargo := line[start:end]
				if strings.TrimSpace(cargo) != "" {
					_, err := strconv.Atoi(cargo)
					if err != nil {
						importedStacks[i].PushLast(cargo)
					}
				}
			}
		}

		if line == "" {
			stacks <- importedStacks
			close(stacks)
		}
	}
	close(instructions)
}

func Crates9000(filename string) string {
	stacks := make(chan []util.Stack[string], 1)
	instructions := make(chan instruction, 10)

	go readData(filename, stacks, instructions)

	importedStacks := <-stacks

	for instruct := range instructions {
		for i := 0; i < instruct.amount; i++ {
			cargo, err := importedStacks[instruct.from].Pop()
			if err != nil {
				panic(err)
			}
			importedStacks[instruct.to].Push(cargo)
		}
	}

	result := ""
	for _, s := range importedStacks {
		result += s[0]
	}

	return result
}

func Crates9001(filename string) string {
	stacks := make(chan []util.Stack[string], 1)
	instructions := make(chan instruction, 10)

	go readData(filename, stacks, instructions)

	importedStacks := <-stacks

	for instruct := range instructions {
		cargo, err := importedStacks[instruct.from].PopMultiple(instruct.amount)
		if err != nil {
			panic(err)
		}
		importedStacks[instruct.to].PushAll(cargo)
	}

	result := ""
	for _, s := range importedStacks {
		result += s[0]
	}

	return result
}
