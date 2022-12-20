package day18

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type point struct {
	x, y, z int
}

type side struct {
	points []point
}

func (s *side) equals(other *side) bool {
	if len(s.points) != len(other.points) {
		return false
	}

	for _, p := range s.points {
		if !slices.Contains(other.points, p) {
			return false
		}
	}
	return true
}

func (s *side) contains(p1, p2 point) bool {
	return (s.points[0] == p1 && s.points[1] == p2) ||
		(s.points[1] == p1 && s.points[0] == p2) ||
		(s.points[1] == p1 && s.points[2] == p2) ||
		(s.points[2] == p1 && s.points[1] == p2) ||
		(s.points[2] == p1 && s.points[3] == p2) ||
		(s.points[3] == p1 && s.points[2] == p2) ||
		(s.points[3] == p1 && s.points[0] == p2) ||
		(s.points[0] == p1 && s.points[3] == p2)
}

type cube struct {
	id        int
	start     point
	freeSides []side
	sides     []side
}

func readData(filename string) []*cube {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	cubes := make([]*cube, 0)
	cubeId := 1

	for scanner.Scan() {
		line := scanner.Text()
		c := strings.Split(line, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		z, _ := strconv.Atoi(c[2])
		start := point{x, y, z}
		cubes = append(cubes, &cube{cubeId, start, allSidesForCube(start), allSidesForCube(start)})
		cubeId++
	}

	return cubes
}

func allSidesForCube(p point) []side {
	return []side{
		{[]point{p, {p.x, p.y + 1, p.z}, {p.x, p.y, p.z + 1}, {p.x, p.y + 1, p.z + 1}}},
		{[]point{p, {p.x + 1, p.y, p.z}, {p.x, p.y, p.z + 1}, {p.x + 1, p.y, p.z + 1}}},
		{[]point{p, {p.x + 1, p.y, p.z}, {p.x, p.y + 1, p.z}, {p.x + 1, p.y + 1, p.z}}},

		{[]point{{p.x + 1, p.y + 1, p.z + 1}, {p.x + 1, p.y, p.z + 1}, {p.x + 1, p.y + 1, p.z}, {p.x + 1, p.y, p.z}}},
		{[]point{{p.x + 1, p.y + 1, p.z + 1}, {p.x, p.y + 1, p.z + 1}, {p.x + 1, p.y + 1, p.z}, {p.x, p.y + 1, p.z}}},
		{[]point{{p.x + 1, p.y + 1, p.z + 1}, {p.x, p.y + 1, p.z + 1}, {p.x + 1, p.y, p.z + 1}, {p.x, p.y, p.z + 1}}},
	}
}

func removeOverlappingSide(c1, c2 *cube) {
	i := 0
	for _, s1 := range c1.freeSides {
		toRemoveFromS2 := -1
		for j, s2 := range c2.freeSides {
			if s1.equals(&s2) {
				toRemoveFromS2 = j
				break
			}
		}

		if toRemoveFromS2 >= 0 {
			c2.freeSides = append(c2.freeSides[0:toRemoveFromS2], c2.freeSides[toRemoveFromS2+1:]...)
		} else {
			c1.freeSides[i] = s1
			i++
		}
	}

	c1.freeSides = c1.freeSides[0:i]
}

func Part1(filename string) int {
	cubes := readData(filename)

	for i, c1 := range cubes {
		for j := i + 1; j < len(cubes); j++ {
			c2 := cubes[j]
			removeOverlappingSide(c1, c2)
		}
	}

	total := 0
	for _, c := range cubes {
		total += len(c.freeSides)
	}

	return total
}

func areCubesConnected(c1, c2 *cube) bool {
	for _, s1 := range c1.sides {
		for _, s2 := range c2.sides {
			if areSidesConnected(s1, s2) {
				return true
			}
		}
	}
	return false
}

func areSidesConnected(s1, s2 side) bool {
	return s1.contains(s2.points[0], s2.points[1]) || s1.contains(s2.points[1], s2.points[2]) || s1.contains(s2.points[2], s2.points[3]) || s1.contains(s2.points[3], s2.points[0])
}

func createKey(s side) string {
	key := ""
	for _, p := range s.points {
		key += strconv.Itoa(p.x) + strconv.Itoa(p.y) + strconv.Itoa(p.z)
	}
	return key
}

var cache map[*cube]map[int]*cube = make(map[*cube]map[int]*cube)

func findOutsideCubes(c *cube, cubes []*cube) map[int]*cube {
	key := c // createKey(s)

	if value, hit := cache[key]; hit {
		return value
	}

	connectedCubes := map[int]*cube{c.id: c}

	for i, other := range cubes {
		if c == other || len(c.freeSides) == 0 || len(other.freeSides) == 0 {
			continue
		}
		if areCubesConnected(c, other) {
			connectedCubes[other.id] = other
			for k, v := range findOutsideCubes(other, append(cubes[0:i], cubes[i+1:]...)) {
				connectedCubes[k] = v
			}
		}
	}

	cache[key] = connectedCubes

	return connectedCubes
}

func inBound(p, min, max point) bool {
	return p.x >= min.x && p.x <= max.x && p.y >= min.y && p.y <= max.y && p.z >= min.z && p.z < max.z
}

func Part2(filename string) int {
	cubes := readData(filename)

	for i, c1 := range cubes {
		for j := i + 1; j < len(cubes); j++ {
			c2 := cubes[j]
			removeOverlappingSide(c1, c2)
		}
	}

	start := cubes[0]
	toDelete := 0
	for i := 1; i < len(cubes); i++ {
		c := cubes[i]
		if c.start.x < start.start.x {
			start = c
			toDelete = i
		}
	}

	total := 0
	for _, c := range findOutsideCubes(start, append(cubes[:toDelete], cubes[toDelete+1:]...)) {
		total += len(c.freeSides)
	}

	return total
}
