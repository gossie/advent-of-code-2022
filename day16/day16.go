package day16

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var resultCache map[string]int

func createKey(s *state) string {
	return s.current.name + strconv.Itoa(s.openValves) + strconv.Itoa(s.timeLeft) + strconv.Itoa(s.numberOfElephants)
}

type valve struct {
	id              int
	name            string
	flowRate        int
	connectedValves []string
}

type state struct {
	current                                                   *valve
	allValves                                                 map[string]*valve
	openValves, timeLeft, numberOfElephants, pressureReleased int
}

func readData(filename string) map[string]*valve {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	id := 1
	valves := make(map[string]*valve)

	for scanner.Scan() {
		line := scanner.Text()
		name := line[6:8]

		index1 := strings.Index(line, "=")
		index2 := strings.Index(line, ";")

		flowRate, _ := strconv.Atoi(line[index1+1 : index2])

		connections := strings.Split(line, ", ")
		connections[0] = connections[0][len(connections[0])-2:]

		valve := &valve{id: id, name: name, flowRate: flowRate, connectedValves: connections}
		id *= 2
		valves[name] = valve
	}

	return valves
}

func possiblePaths(current *state) int {
	if current.timeLeft == 0 {
		if current.numberOfElephants > 0 {
			newState := &state{current.allValves["AA"], current.allValves, current.openValves, 26, current.numberOfElephants - 1, current.pressureReleased}
			return possiblePaths(newState)
		} else {
			return 0
		}
	}

	key := createKey(current)
	if cached, hit := resultCache[key]; hit {
		return cached
	}

	pressure := 0

	for _, name := range current.current.connectedValves {
		newState := &state{current.allValves[name], current.allValves, current.openValves, current.timeLeft - 1, current.numberOfElephants, current.pressureReleased}
		pressure = int(math.Max(float64(pressure), float64(possiblePaths(newState))))
	}

	alreadyOpen := current.openValves & current.current.id
	if current.current.flowRate > 0 && alreadyOpen == 0 {
		newOpenValves := current.openValves | current.current.id
		newState := &state{current.current, current.allValves, newOpenValves, current.timeLeft - 1, current.numberOfElephants, current.pressureReleased}
		pressure = int(math.Max(float64(pressure), float64((current.timeLeft-1)*current.current.flowRate+possiblePaths(newState))))
	}

	resultCache[key] = pressure

	return pressure
}

func Part1(filename string) int {
	valves := readData(filename)
	resultCache = make(map[string]int)
	return possiblePaths(&state{valves["AA"], valves, 0, 30, 0, 0})
}

func Part2(filename string) int {
	valves := readData(filename)
	resultCache = make(map[string]int)
	return possiblePaths(&state{valves["AA"], valves, 0, 26, 1, 0})
}
