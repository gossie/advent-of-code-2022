package day15

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) manhattenDistanceTo(other point) int {
	return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)))
}

type beacon struct {
	position point
}

type sensor struct {
	position      point
	closestBeacon *beacon
}

func (s *sensor) manhattenDistanceToBeacon() int {
	p1 := s.position
	p2 := s.closestBeacon.position
	return p1.manhattenDistanceTo(p2)
}

type beaconFreeArray struct {
	center point
	radius int
}

func (bfa *beaconFreeArray) containsYCoordinate(y int) bool {
	return y <= bfa.center.y+bfa.radius && y >= bfa.center.y-bfa.radius
}

func (bfa *beaconFreeArray) countBeaconFreePostions(y int, beacons map[point]*beacon, foundPositions map[point]bool) {
	var start, end point
	yDifference := y - bfa.center.y

	switch {
	case yDifference > 0:
		start = point{(bfa.center.x - bfa.radius) + int(math.Abs(float64(yDifference))), bfa.center.y + yDifference}
		end = point{(bfa.center.x + bfa.radius) - int(math.Abs(float64(yDifference))), bfa.center.y + yDifference}
	case yDifference < 0:
		start = point{(bfa.center.x - bfa.radius) + int(math.Abs(float64(yDifference))), bfa.center.y - yDifference}
		end = point{(bfa.center.x + bfa.radius) - int(math.Abs(float64(yDifference))), bfa.center.y - yDifference}
	case yDifference == 0:
		start = point{bfa.center.x - bfa.radius, bfa.center.y}
		end = point{bfa.center.x + bfa.radius, bfa.center.y}
	}

	for i := start.x; i <= end.x; i++ {
		position := point{i, y}
		if _, isBeacon := beacons[position]; !isBeacon {
			if _, alreadyFound := foundPositions[position]; !alreadyFound {
				foundPositions[position] = true
			}
		}
	}
}

func (bfa *beaconFreeArray) includes(p point) (bool, int) {
	distance := bfa.center.manhattenDistanceTo(p)
	if distance > bfa.radius {
		return false, -1
	}

	yDifference := p.y - bfa.center.y
	var start, end point
	switch {
	case yDifference > 0:
		start = point{(bfa.center.x - bfa.radius) + int(math.Abs(float64(yDifference))), bfa.center.y + yDifference}
		end = point{(bfa.center.x + bfa.radius) - int(math.Abs(float64(yDifference))), bfa.center.y + yDifference}
	case yDifference < 0:
		start = point{(bfa.center.x - bfa.radius) + int(math.Abs(float64(yDifference))), bfa.center.y - yDifference}
		end = point{(bfa.center.x + bfa.radius) - int(math.Abs(float64(yDifference))), bfa.center.y - yDifference}
	case yDifference == 0:
		start = point{bfa.center.x - bfa.radius, bfa.center.y}
		end = point{bfa.center.x + bfa.radius, bfa.center.y}
	}
	return p.x >= start.x && p.x <= end.x, end.x
}

func toPoint(sensorPosition []string) point {
	x, _ := strconv.Atoi(strings.Split(sensorPosition[0], "=")[1])
	y, _ := strconv.Atoi(strings.Split(sensorPosition[1], "=")[1])
	return point{x, y}
}

func readData(filename string) ([]*sensor, map[point]*beacon) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sensors := make([]*sensor, 0)
	beacons := make(map[point]*beacon)

	for scanner.Scan() {
		line := scanner.Text() // Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		colonIndex := strings.IndexRune(line, ':')
		sensorPosition := toPoint(strings.Split(line[10:colonIndex], ", "))
		beaconPosition := toPoint(strings.Split(line[colonIndex+1:][21:], ", "))

		b, found := beacons[beaconPosition]
		if !found {
			b = &beacon{beaconPosition}
			beacons[beaconPosition] = b
		}
		sensors = append(sensors, &sensor{sensorPosition, b})
	}

	return sensors, beacons
}

func NumberOfPositionsWithoutBeacon(filename string, y int) int {
	sensors, beacons := readData(filename)
	beaconFreeAreas := make([]beaconFreeArray, 0)

	for _, sensor := range sensors {
		beaconFreeAreas = append(beaconFreeAreas, beaconFreeArray{sensor.position, sensor.manhattenDistanceToBeacon()})
	}

	positions := make(map[point]bool)
	for _, area := range beaconFreeAreas {
		if area.containsYCoordinate(y) {
			area.countBeaconFreePostions(y, beacons, positions)
		}
	}

	return len(positions)
}

func TuningFrequency(filename string, maxXY int) uint64 {
	sensors, _ := readData(filename)
	beanonFreeAreas := make([]*beaconFreeArray, 0)

	for _, sensor := range sensors {
		beanonFreeAreas = append(beanonFreeAreas, &beaconFreeArray{sensor.position, sensor.manhattenDistanceToBeacon()})
	}

	for y := 0; y <= maxXY; y++ {
	xloop:
		for x := 0; x <= maxXY; x++ {
			p := point{x, y}
			for _, bfa := range beanonFreeAreas {
				if found, newX := bfa.includes(p); found {
					x = newX
					continue xloop
				}
			}
			return uint64(p.x)*4000000 + uint64(p.y)
		}
	}

	return 0
}
