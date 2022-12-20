package day17

import (
	"bufio"
	"os"
	"sort"
)

type shape [][]int

type point struct {
	x, y int
}

type rock struct {
	topLeft point
	shape   *shape
}

func (r *rock) width() int {
	return len((*r.shape)[0])
}

var shape1 shape = [][]int{{1, 1, 1, 1}}
var shape2 shape = [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
var shape3 shape = [][]int{{0, 0, 1}, {0, 0, 1}, {1, 1, 1}}
var shape4 shape = [][]int{{1}, {1}}
var shape5 shape = [][]int{{1}, {1}}

var allShapes []shape = []shape{shape1, shape2, shape3, shape4, shape5}

func readData(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return scanner.Text()
}

func calculateStepsTillCollision(r *rock, landedRocks []*rock) int {
	if len(landedRocks) == 0 {
		return 3
	}

	xStart := r.topLeft.x
	xEnd := r.topLeft.x + r.width() - 1

	for _, landedRock := range landedRocks {
		if landedRock.topLeft.y < r.topLeft.y {
			landedXStart := landedRock.topLeft.x
			landedXEnd := landedRock.topLeft.x + r.width() - 1

			if (landedXStart <= xEnd && landedXStart >= xStart) || (landedXEnd <= xEnd && landedXEnd >= xStart) {
				return (r.topLeft.y + len(*r.shape)) - landedRock.topLeft.y
			}
		}
	}

	return r.topLeft.y - len(*r.shape)
}

func fall(r *rock, jetPattern string, jetPatternIndex *int, landedRocks []*rock) {
	yDiffTillCollision := calculateStepsTillCollision(r, landedRocks)

	for yDiffTillCollision >= 0 {
		dir := jetPattern[*jetPatternIndex%len(jetPattern)]
		*jetPatternIndex++

		switch dir {
		case '>':
			if (r.topLeft.x+r.width())-1 < 6 {
				r.topLeft.x++
			}
		case '<':
			if r.topLeft.x > 0 {
				r.topLeft.x--
			}
		default:
			panic("unknown direction: " + string(dir))
		}

		newYDiffTillCollision := calculateStepsTillCollision(r, landedRocks)
		if yDiffTillCollision > 0 {
			r.topLeft.y--
		}

		yDiffTillCollision = newYDiffTillCollision
	}
}

func Height(filename string) int {
	jetPattern := readData(filename)

	landedRocks := make([]*rock, 0)
	jetPatternIndex := 0
	for i := 0; i < 2022; i++ {
		shapeIndex := i % len(allShapes)
		nextShape := &allShapes[shapeIndex]
		y := 3
		if len(landedRocks) > 0 {
			y = landedRocks[0].topLeft.y + 3 + len(*nextShape)
		}
		startPoint := point{2, y}

		r := &rock{startPoint, nextShape}
		fall(r, jetPattern, &jetPatternIndex, landedRocks)
		landedRocks = append(landedRocks, r)
		sort.Slice(landedRocks, func(i, j int) bool {
			return landedRocks[i].topLeft.y > landedRocks[j].topLeft.y
		})
	}

	return landedRocks[0].topLeft.y
}
