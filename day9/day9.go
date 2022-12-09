package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	deltaX, deltaY int
}

type point struct {
	x, y int
}

type knot struct {
	position point
}

func readData(filename string, instructions chan<- *instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		instructionComponents := strings.Split(line, " ")
		direction := instructionComponents[0]
		steps, _ := strconv.Atoi(instructionComponents[1])
		switch direction {
		case "U":
			instructions <- &instruction{0, steps}
		case "D":
			instructions <- &instruction{0, -steps}
		case "L":
			instructions <- &instruction{-steps, 0}
		case "R":
			instructions <- &instruction{steps, 0}
		default:
			panic("unknown direction")
		}
	}
	close(instructions)
}

func Visited(filename string, numberOfknots int) int {
	instructions := make(chan *instruction, 10)
	go readData(filename, instructions)

	knots := make([]*knot, 0, numberOfknots)
	for i := 0; i < numberOfknots; i++ {
		knots = append(knots, &knot{position: point{0, 0}})
	}

	visited := map[point]bool{{0, 0}: true}

	for inst := range instructions {
		moveHead(inst, knots[0], knots[1:], visited)
	}

	return len(visited)
}

func moveHead(inst *instruction, head *knot, tail []*knot, visited map[point]bool) {
	if inst.deltaX == 0 {
		if inst.deltaY > 0 {
			for y := 1; y <= inst.deltaY; y++ {
				head.position.y = head.position.y + 1
				moveTailKnots(head, tail, visited)
			}
		} else {
			for y := -1; y >= inst.deltaY; y-- {
				head.position.y = head.position.y - 1
				moveTailKnots(head, tail, visited)
			}
		}
	} else if inst.deltaY == 0 {
		if inst.deltaX > 0 {
			for x := 1; x <= inst.deltaX; x++ {
				head.position.x = head.position.x + 1
				moveTailKnots(head, tail, visited)
			}
		} else {
			for x := -1; x >= inst.deltaX; x-- {
				head.position.x = head.position.x - 1
				moveTailKnots(head, tail, visited)
			}
		}
	}
}

func moveTailKnots(head *knot, tail []*knot, visited map[point]bool) {
	if len(tail) == 0 {
		return
	}

	current := tail[0]
	if needsToMoveDiagonal(head, current) {
		moveDiagonal(head, current, visited, len(tail) == 1)
		followPrevious(head, current, visited, len(tail) == 1)
	} else {
		followPrevious(head, current, visited, len(tail) == 1)
	}

	moveTailKnots(current, tail[1:], visited)
}

func needsToMoveDiagonal(head, tail *knot) bool {
	return head.position.x != tail.position.x && head.position.y != tail.position.y
}

func moveDiagonal(head, current *knot, visited map[point]bool, addVisit bool) {
	if head.position.x == current.position.x+2 && head.position.y == current.position.y+2 {
		current.position.x = current.position.x + 1
		current.position.y = current.position.y + 1
	} else if head.position.x == current.position.x+2 && head.position.y == current.position.y-2 {
		current.position.x = current.position.x + 1
		current.position.y = current.position.y - 1
	} else if head.position.x == current.position.x-2 && head.position.y == current.position.y-2 {
		current.position.x = current.position.x - 1
		current.position.y = current.position.y - 1
	} else if head.position.x == current.position.x-2 && head.position.y == current.position.y+2 {
		current.position.x = current.position.x - 1
		current.position.y = current.position.y + 1
	} else if head.position.x == current.position.x+1 && head.position.y > current.position.y+1 {
		current.position.x = head.position.x
		current.position.y = current.position.y + 1
	} else if head.position.x == current.position.x+1 && head.position.y < current.position.y-1 {
		current.position.x = head.position.x
		current.position.y = current.position.y - 1
	} else if head.position.y == current.position.y+1 && head.position.x > current.position.x+1 {
		current.position.x = current.position.x + 1
		current.position.y = head.position.y
	} else if head.position.y == current.position.y+1 && head.position.x < current.position.x-1 {
		current.position.x = current.position.x - 1
		current.position.y = head.position.y
	} else if head.position.x == current.position.x-1 && head.position.y > current.position.y+1 {
		current.position.x = head.position.x
		current.position.y = current.position.y + 1
	} else if head.position.x == current.position.x-1 && head.position.y < current.position.y-1 {
		current.position.x = head.position.x
		current.position.y = current.position.y - 1
	} else if head.position.y == current.position.y-1 && head.position.x > current.position.x+1 {
		current.position.x = current.position.x + 1
		current.position.y = head.position.y
	} else if head.position.y == current.position.y-1 && head.position.x < current.position.x-1 {
		current.position.x = current.position.x - 1
		current.position.y = head.position.y
	}
	addVisitIfNecessary(visited, current.position, addVisit)
}

func followPrevious(head, tail *knot, visited map[point]bool, addVisit bool) {
	if head.position.x == tail.position.x {
		if head.position.y > tail.position.y+1 {
			followHorizontally(tail.position.y+1, head.position.y, 1, tail, visited, addVisit)
		} else if head.position.y < tail.position.y-1 {
			followHorizontally(tail.position.y-1, head.position.y, -1, tail, visited, addVisit)
		}
	} else {
		if head.position.x > tail.position.x+1 {
			followVertically(tail.position.x+1, head.position.x, 1, tail, visited, addVisit)
		} else if head.position.x < tail.position.x-1 {
			followVertically(tail.position.x-1, head.position.x, -1, tail, visited, addVisit)
		}
	}
}

func followHorizontally(start, end, inc int, tail *knot, visited map[point]bool, addVisit bool) {
	for i := start; i != end; i += inc {
		tail.position.y = i
		addVisitIfNecessary(visited, point{tail.position.x, i}, addVisit)
	}
}

func followVertically(start, end, inc int, tail *knot, visited map[point]bool, addVisit bool) {
	for i := start; i != end; i += inc {
		tail.position.x = i
		addVisitIfNecessary(visited, point{i, tail.position.y}, addVisit)
	}
}

func addVisitIfNecessary(visited map[point]bool, key point, add bool) {
	if add {
		visited[key] = true
	}
}
