package day14

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gossie/aoc-utils/geometry"
)

func parsePoint(pointAsString string) geometry.Point2d {
	components := strings.Split(pointAsString, ",")
	x, err := strconv.Atoi(components[0])
	if err != nil {
		panic(components[0] + " cannot be parsed into coordinate")
	}

	y, err := strconv.Atoi(components[1])
	if err != nil {
		panic(components[0] + " cannot be parsed into coordinate")
	}

	return geometry.CreatePoint2d(x, y)
}

func addPointsInBetween(points map[geometry.Point2d]bool, p1, p2 geometry.Point2d) {
	if p1.X == p2.X {
		if p1.Y < p2.Y {
			for y := p1.Y + 1; y < p2.Y; y++ {
				points[geometry.CreatePoint2d(p1.X, y)] = true
			}
		} else {
			for y := p2.Y + 1; y < p1.Y; y++ {
				points[geometry.CreatePoint2d(p1.X, y)] = true
			}
		}
	} else {
		if p1.X < p2.X {
			for x := p1.X + 1; x < p2.X; x++ {
				points[geometry.CreatePoint2d(x, p1.Y)] = true
			}
		} else {
			for x := p2.X + 1; x < p1.X; x++ {
				points[geometry.CreatePoint2d(x, p1.Y)] = true
			}
		}
	}
}

func readData(filename string) map[geometry.Point2d]bool {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	points := make(map[geometry.Point2d]bool)

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

func moveDown(points map[geometry.Point2d]bool, sand *geometry.Point2d) (*geometry.Point2d, error) {
	next := geometry.CreatePoint2d(sand.X, sand.Y+1)
	if _, found := points[next]; found {
		return nil, errors.New("point taken")
	}
	return &next, nil
}

func moveLeft(points map[geometry.Point2d]bool, sand *geometry.Point2d) (*geometry.Point2d, error) {
	next := geometry.CreatePoint2d(sand.X-1, sand.Y+1)
	if _, found := points[next]; found {
		return nil, errors.New("point taken")
	}
	return &next, nil
}

func moveRight(points map[geometry.Point2d]bool, sand *geometry.Point2d) (*geometry.Point2d, error) {
	next := geometry.CreatePoint2d(sand.X+1, sand.Y+1)
	if _, found := points[next]; found {
		return nil, errors.New("point taken")
	}
	return &next, nil
}

func move(points map[geometry.Point2d]bool, sand *geometry.Point2d, highestY int) *geometry.Point2d {
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
	if next.Y > highestY {
		return nil
	}
	return move(points, next, highestY)
}

func moveWithFloor(points map[geometry.Point2d]bool, sand *geometry.Point2d, highestY int) *geometry.Point2d {
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
	if next.Y > highestY {
		return sand
	}
	return moveWithFloor(points, next, highestY)
}

func highestY(points map[geometry.Point2d]bool) int {
	heighestY := 0
	for p := range points {
		heighestY = int(math.Max(float64(heighestY), float64(p.Y)))
	}
	return heighestY
}

func Part1(filename string) int {
	points := readData(filename)

	highestY := highestY(points)
	sandDrops := 0

	for {
		newPoint := geometry.CreatePoint2d(500, 0)
		rested := move(points, &newPoint, highestY)
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
		newPoint := geometry.CreatePoint2d(500, 0)
		rested := moveWithFloor(points, &newPoint, highestY)
		if rested.X == 500 && rested.Y == 0 {
			break
		}
		sandDrops++
		//rested.Y--
		points[*rested] = true
	}

	return sandDrops + 1
}
