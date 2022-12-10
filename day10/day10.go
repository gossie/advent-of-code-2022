package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction interface {
	offsets() (int, int)
}

type noop struct{}

func (n noop) offsets() (int, int) {
	return 1, 0
}

type addX struct {
	n int
}

func (a addX) offsets() (int, int) {
	return 2, a.n
}

func readData(filename string, instructions chan instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "noop" {
			instructions <- noop{}
		} else {
			n, _ := strconv.Atoi(strings.Split(line, " ")[1])
			instructions <- addX{n}
		}
	}
	close(instructions)
}

func SignalStrength(filename string) int64 {
	instructions := make(chan instruction, 10)

	signalStrength := int64(0)

	x := 1

	go readData(filename, instructions)

	cycles := 0
	for inst := range instructions {
		cyclesToAdd, addedDeltaX := inst.offsets()
		for i := 0; i < cyclesToAdd; i++ {
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				newStrength := int64(cycles * x)
				signalStrength += newStrength
			}
		}

		x += addedDeltaX
	}

	return signalStrength
}

func Sprite(filename string) {
	instructions := make(chan instruction, 10)

	go readData(filename, instructions)

	screen := make([]string, 6*40)
	x := 1
	cycles := 0
	screen[cycles] = "#"

	for inst := range instructions {
		cyclesToAdd, addedDeltaX := inst.offsets()
		for i := 0; i < cyclesToAdd; i++ {
			toCompare := cycles % 40
			if x-1 == toCompare || x == toCompare || x+1 == toCompare {
				screen[cycles] = "#"
			} else {
				screen[cycles] = "."
			}
			cycles++
		}

		x += addedDeltaX
	}

	for i := 0; i < 6; i++ {
		fmt.Println(screen[i*40 : i*40+40])
	}
}
