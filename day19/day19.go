package day19

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func createKey(s *state) string {
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", s.blueprint.id, s.timeLeft, s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots, s.ore, s.clay, s.obsidian, s.geode, s.allowedToBuild)
}

type costs struct {
	ore, clay, obsidian int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type blueprint struct {
	id            int
	oreRobot      costs
	clayRobot     costs
	obsidianRobot costs
	geodeRobot    costs
}

func (b *blueprint) maxOreCosts() int {
	return max(max(max(b.oreRobot.ore, b.clayRobot.ore), b.obsidianRobot.ore), b.geodeRobot.ore)
}

func (b *blueprint) maxClayCosts() int {
	return max(max(max(b.oreRobot.clay, b.clayRobot.clay), b.obsidianRobot.clay), b.geodeRobot.clay)
}

func (b *blueprint) maxObsidianCosts() int {
	return max(max(max(b.oreRobot.obsidian, b.clayRobot.obsidian), b.obsidianRobot.obsidian), b.geodeRobot.obsidian)
}

type state struct {
	blueprint                  *blueprint
	timeLeft                   int
	ore, clay, obsidian, geode int
	oreRobots                  int
	clayRobots                 int
	obsidianRobots             int
	geodeRobots                int
	allowedToBuild             int
}

func (s *state) subtract(c *costs) {
	s.ore -= c.ore
	s.clay -= c.clay
	s.obsidian -= c.obsidian
}

func (s *state) collect() {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
}

func readData(filename string) []*blueprint {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	blueprints := make([]*blueprint, 0)
	blueprintId := 1

	for scanner.Scan() {
		line := scanner.Text()
		costsOfTheRest := line[strings.Index(line, ":")+2:]

		dotIndex := strings.Index(costsOfTheRest, ".")
		oreRobotCostsStr := costsOfTheRest[:dotIndex]
		ore, _ := strconv.Atoi(strings.TrimSpace(oreRobotCostsStr[20 : len(oreRobotCostsStr)-3]))
		oreRobotCosts := costs{ore: ore}
		costsOfTheRest = costsOfTheRest[dotIndex+2:]

		dotIndex = strings.Index(costsOfTheRest, ".")
		clayRobotCostsStr := costsOfTheRest[:dotIndex]
		ore, _ = strconv.Atoi(strings.TrimSpace(clayRobotCostsStr[21 : len(clayRobotCostsStr)-3]))
		clayRobotCosts := costs{ore: ore}
		costsOfTheRest = costsOfTheRest[dotIndex+2:]

		dotIndex = strings.Index(costsOfTheRest, ".")
		obsidianRobotCostsStr := costsOfTheRest[:dotIndex]
		ore, _ = strconv.Atoi(strings.TrimSpace(obsidianRobotCostsStr[25:strings.Index(obsidianRobotCostsStr, "ore")]))
		clay, _ := strconv.Atoi(strings.TrimSpace(obsidianRobotCostsStr[strings.Index(obsidianRobotCostsStr, "ore")+7 : len(obsidianRobotCostsStr)-4]))
		obsidianRobotCosts := costs{ore: ore, clay: clay}
		costsOfTheRest = costsOfTheRest[dotIndex+2:]

		dotIndex = strings.Index(costsOfTheRest, ".")
		geodeRobotCostsStr := costsOfTheRest[:dotIndex]
		ore, _ = strconv.Atoi(strings.TrimSpace(geodeRobotCostsStr[22:strings.Index(geodeRobotCostsStr, "ore")]))
		obsidian, _ := strconv.Atoi(strings.TrimSpace(geodeRobotCostsStr[strings.Index(geodeRobotCostsStr, "ore")+7 : len(geodeRobotCostsStr)-8]))
		geodeRobotCosts := costs{ore: ore, obsidian: obsidian}

		blueprints = append(blueprints, &blueprint{id: blueprintId, oreRobot: oreRobotCosts, clayRobot: clayRobotCosts, obsidianRobot: obsidianRobotCosts, geodeRobot: geodeRobotCosts})

		blueprintId++
	}

	return blueprints
}

func canAfford(s *state, c *costs) bool {
	return s.ore >= c.ore && s.clay >= c.clay && s.obsidian >= c.obsidian
}

// func isItPossibleToBuildAnotherGeodeRobot(s *state) bool {
// 	if s.timeLeft == 0 {
// 		return false
// 	}

// 	if s.timeLeft == 1 && s.blueprint.geodeRobot.obsidian > s.obsidian+s.obsidianRobots {
// 		return false
// 	}

// 	return true
// }

func numberOfGeodes(current state) int {
	if current.timeLeft == 0 {
		return current.geode
	}

	key := createKey(&current)
	if value, hit := cache[key]; hit {
		return value
	}

	max := 0
	allowedToBuildInNextCollectState := 7

	if canAfford(&current, &current.blueprint.geodeRobot) {
		newState4 := current
		newState4.timeLeft--

		newState4.subtract(&newState4.blueprint.geodeRobot)
		newState4.collect()
		newState4.geodeRobots++
		newState4.allowedToBuild = 7
		max = int(math.Max(float64(max), float64(numberOfGeodes(newState4))))
	} else {
		if canAfford(&current, &current.blueprint.oreRobot) && current.allowedToBuild&4 > 0 && current.oreRobots*current.timeLeft+current.ore < current.timeLeft*current.blueprint.maxOreCosts() {
			allowedToBuildInNextCollectState &= 3

			newState1 := current
			newState1.timeLeft--
			newState1.subtract(&newState1.blueprint.oreRobot)
			newState1.collect()
			newState1.oreRobots++
			newState1.allowedToBuild = 7
			max = int(math.Max(float64(max), float64(numberOfGeodes(newState1))))
		}

		if canAfford(&current, &current.blueprint.clayRobot) && current.allowedToBuild&2 > 0 && current.clayRobots*current.timeLeft+current.clay < current.timeLeft*current.blueprint.maxClayCosts() {
			allowedToBuildInNextCollectState &= 5

			newState2 := current
			newState2.timeLeft--
			newState2.subtract(&newState2.blueprint.clayRobot)
			newState2.collect()
			newState2.clayRobots++
			newState2.allowedToBuild = 7
			max = int(math.Max(float64(max), float64(numberOfGeodes(newState2))))
		}

		if canAfford(&current, &current.blueprint.obsidianRobot) && current.allowedToBuild&1 > 0 && current.obsidianRobots*current.timeLeft+current.obsidian < current.timeLeft*current.blueprint.maxObsidianCosts() {
			allowedToBuildInNextCollectState |= 6

			newState3 := current
			newState3.timeLeft--
			newState3.subtract(&newState3.blueprint.obsidianRobot)
			newState3.collect()
			newState3.obsidianRobots++
			newState3.allowedToBuild = 7
			max = int(math.Max(float64(max), float64(numberOfGeodes(newState3))))
		}

		newState5 := current
		newState5.allowedToBuild = allowedToBuildInNextCollectState
		newState5.timeLeft--
		newState5.collect()
		max = int(math.Max(float64(max), float64(numberOfGeodes(newState5))))
	}

	cache[key] = max

	return max
}

func Part1(filename string) int {
	blueprints := readData(filename)

	result := 0
	for _, bp := range blueprints {
		s := state{blueprint: bp, timeLeft: 24, oreRobots: 1, allowedToBuild: 7}
		result += bp.id * numberOfGeodes(s)
	}
	return result
}

func Part2(filename string) int {
	blueprints := readData(filename)

	result := 1
	for i := 0; i < 3; i++ {
		s := state{blueprint: blueprints[i], timeLeft: 32, oreRobots: 1, allowedToBuild: 7}
		result *= numberOfGeodes(s)
	}
	return result
}
