package day14

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func parsePoint(pointAsString string) point {
	components := strings.Split(pointAsString, ",")
	x, err := strconv.Atoi(components[0])
	if err != nil {
		panic(components[0] + " cannot be parsed into coordinate")
	}

	y, err := strconv.Atoi(components[1])
	if err != nil {
		panic(components[0] + " cannot be parsed into coordinate")
	}

	return point{x, y}
}

func addPointsInBetween(points map[point]bool, p1, p2 point) {
	if p1.x == p2.x {
		if p1.y < p2.y {
			for y := p1.y + 1; y < p2.y; y++ {
				points[point{p1.x, y}] = true
			}
		} else {
			for y := p2.y + 1; y < p1.y; y++ {
				points[point{p1.x, y}] = true
			}
		}
	} else {
		if p1.x < p2.x {
			for x := p1.x + 1; x < p2.x; x++ {
				points[point{x, p1.y}] = true
			}
		} else {
			for x := p2.x + 1; x < p1.x; x++ {
				points[point{x, p1.y}] = true
			}
		}
	}
}

func readData(filename string) map[point]bool {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	points := make(map[point]bool)

	for scanner.Scan() {
		line := scanner.Text()
		pointsAsString := strings.Split(line, " -> ")
		startAsString := pointsAsString[0]
		for i := 1; i < len(pointsAsString); i++ {
			endAsString := pointsAsString[i]

			p1 := parsePoint(startAsString)
			p2 := parsePoint(endAsString)

			points[p1] = true
			points[p2] = true

			addPointsInBetween(points, p1, p2)

			startAsString = endAsString

			tmp := parsePoint(startAsString)
			_, ok := points[tmp]
			if !ok {
				panic("point not found")
			}

		}
	}
	return points
}

func moveDown(points map[point]bool, sand *point) (*point, error) {
	next := point{sand.x, sand.y + 1}
	if _, found := points[next]; found {
		return nil, errors.New("Point taken")
	}
	return &next, nil
}

func moveLeft(points map[point]bool, sand *point) (*point, error) {
	next := point{sand.x - 1, sand.y + 1}
	if _, found := points[next]; found {
		return nil, errors.New("Point taken")
	}
	return &next, nil
}

func moveRight(points map[point]bool, sand *point) (*point, error) {
	next := point{sand.x + 1, sand.y + 1}
	if _, found := points[next]; found {
		return nil, errors.New("Point taken")
	}
	return &next, nil
}

func move(points map[point]bool, sand *point, highestY int) *point {
	next, err := moveDown(points, sand)
	if err != nil {
		next, err = moveLeft(points, sand)
		if err != nil {
			next, err = moveRight(points, sand)
			if err != nil {
				return sand
			}
		}
	}
	if next.y > highestY {
		return nil
	}
	return move(points, next, highestY)
}

func moveWithFloor(points map[point]bool, sand *point, highestY int) *point {
	next, err := moveDown(points, sand)
	if err != nil {
		next, err = moveLeft(points, sand)
		if err != nil {
			next, err = moveRight(points, sand)
			if err != nil {
				return sand
			}
		}
	}
	if next.y > highestY {
		return sand
	}
	return moveWithFloor(points, next, highestY)
}

func highestY(points map[point]bool) int {
	heighestY := 0
	for p := range points {
		heighestY = int(math.Max(float64(heighestY), float64(p.y)))
	}
	return heighestY
}

func Part1(filename string) int {
	points := readData(filename)

	highestY := highestY(points)
	sandDrops := 0

	for {
		rested := move(points, &point{500, 0}, highestY)
		if rested == nil {
			break
		}
		sandDrops++
		points[*rested] = true
	}

	return sandDrops
}

func Part2(filename string) int {
	points := readData(filename)

	highestY := highestY(points) + 1
	sandDrops := 0

	for {
		rested := moveWithFloor(points, &point{500, 0}, highestY)
		if rested.x == 500 && rested.y == 0 {
			break
		}
		sandDrops++
		//rested.y--
		points[*rested] = true
	}

	return sandDrops + 1
}
