package day12

import (
	"bufio"
	"container/heap"
	"math"
	"os"
)

type point struct {
	x      int
	y      int
	height int
}

type path struct {
	currentPosition           *point
	totalSteps                int
	totalStepsPlusAirDistance int
}

type priorityQueue []*path

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*path)
	*pq = append(*pq, item)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].totalStepsPlusAirDistance < pq[j].totalStepsPlusAirDistance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Len() int { return len(*pq) }

func (pq *priorityQueue) contains(p *path) bool {
	for i := range *pq {
		if (*pq)[i].currentPosition == p.currentPosition && (*pq)[i].totalSteps == p.totalSteps {
			return true
		}
	}
	return false
}

func readData(filename string) ([][]*point, *point, *point) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	field := make([][]*point, 0)
	var start, dest *point

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := make([]*point, 0)
		line := scanner.Text()
		for x, r := range line {
			var p *point
			if r == 'S' {
				p = &point{x: x, y: len(field), height: int('a')}
				start = p
			} else if r == 'E' {
				p = &point{x: x, y: len(field), height: int('z')}
				dest = p
			} else {
				p = &point{x: x, y: len(field), height: int(r)}
			}

			row = append(row, p)
		}
		field = append(field, row)
	}

	return field, start, dest
}

func airDistance(p, dest *point) int {
	return int(math.Sqrt(math.Pow(math.Abs(float64(dest.x-p.x)), 2) + math.Pow(math.Abs(float64(dest.y-p.y)), 2)))
}

func appendPathIfNecessary(newPaths []*path, currentPath *path, p, dest *point, visitedRiskMapping map[*point]int) []*path {
	if value, present := visitedRiskMapping[p]; !present || value > currentPath.totalSteps+1 {
		newPaths = append(newPaths, &path{currentPosition: p, totalSteps: currentPath.totalSteps + 1, totalStepsPlusAirDistance: currentPath.totalSteps + 1 + airDistance(p, dest)})
	}
	return newPaths
}

func neighborIsReachable(currentPath *path, point *point) bool {
	return point.height <= currentPath.currentPosition.height+1
}

func pathsToNeighbors(field [][]*point, currentPath *path, dest *point, visitedRiskMapping map[*point]int) []*path {
	x, y := currentPath.currentPosition.x, currentPath.currentPosition.y
	newPaths := make([]*path, 0)
	if y-1 >= 0 && neighborIsReachable(currentPath, field[y-1][x]) {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y-1][x], dest, visitedRiskMapping)
	}
	if x+1 <= len(field[0])-1 && neighborIsReachable(currentPath, field[y][x+1]) {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y][x+1], dest, visitedRiskMapping)
	}
	if y+1 <= len(field)-1 && neighborIsReachable(currentPath, field[y+1][x]) {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y+1][x], dest, visitedRiskMapping)
	}
	if x-1 >= 0 && neighborIsReachable(currentPath, field[y][x-1]) {
		newPaths = appendPathIfNecessary(newPaths, currentPath, field[y][x-1], dest, visitedRiskMapping)
	}
	return newPaths
}

func shortestRouteLength(field [][]*point, start, dest *point) int {
	priorityQueue := make(priorityQueue, 0)
	heap.Init(&priorityQueue)

	visitedRiskMapping := map[*point]int{start: 0}

	// images := []*image.Paletted{createFirstImage(len(field[0]), len(field), dest)}

	bestPath := path{currentPosition: start, totalSteps: 0}
	for bestPath.currentPosition != dest {
		// images = append(images, createImage(bestPath.currentPosition))
		for _, p := range pathsToNeighbors(field, &bestPath, dest, visitedRiskMapping) {
			if !priorityQueue.contains(p) {
				heap.Push(&priorityQueue, p)
			}
		}
		if priorityQueue.Len() == 0 {
			return math.MaxInt
		}
		bestPath = *heap.Pop(&priorityQueue).(*path)
		visitedRiskMapping[bestPath.currentPosition] = bestPath.totalSteps
	}

	// createGif("day12.gif", images)

	return bestPath.totalSteps
}

func Part1(filename string) int {
	field, start, dest := readData(filename)
	return shortestRouteLength(field, start, dest)
}

func Part2(filename string) int {
	field, _, dest := readData(filename)

	steps := math.MaxInt
	for _, row := range field {
		for _, p := range row {
			if p.height == int('a') {
				steps = int(math.Min(float64(steps), float64(shortestRouteLength(field, p, dest))))
			}
		}
	}
	return steps
}

// func createGif(name string, images []*image.Paletted) {
// 	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Println("gif could not be created: ", err)
// 	}
// 	defer f.Close()

// 	delays := make([]int, len(images))
// 	for i := 0; i < len(delays); i++ {
// 		delays[i] = 1
// 	}

// 	gif.EncodeAll(f, &gif.GIF{
// 		Image: images,
// 		Delay: delays,
// 	})
// }

// func createFirstImage(width, height int, dest *point) *image.Paletted {
// 	palette := []color.Color{
// 		color.White,
// 		color.Black,
// 		color.RGBA{0x00, 0x00, 0xff, 0xff},
// 		color.RGBA{0xff, 0x00, 0x00, 0xff},
// 	}

// 	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

// 	img.Set(dest.x, dest.y+1, color.RGBA{0xff, 0x00, 0x00, 0xff})
// 	img.Set(dest.x, dest.y, color.RGBA{0xff, 0x00, 0x00, 0xff})
// 	img.Set(dest.x+1, dest.y+1, color.RGBA{0xff, 0x00, 0x00, 0xff})
// 	img.Set(dest.x+1, dest.y, color.RGBA{0xff, 0x00, 0x00, 0xff})

// 	return img
// }

// func createImage(p *point) *image.Paletted {
// 	palette := []color.Color{
// 		color.White,
// 		color.Black,
// 		color.RGBA{0x00, 0x00, 0xff, 0xff},
// 	}

// 	img := image.NewPaletted(image.Rect(p.x, p.y, p.x+1, p.y+1), palette)

// 	x := p.x
// 	y := p.y
// 	box := image.Rect(x, y, x+1, y+1)
// 	//boxImage := img.SubImage(box)
// 	draw.Draw(img, box, &image.Uniform{color.RGBA{0x00, 0x00, 0xff, 0xff}}, image.Point{}, draw.Src)
// 	//draw.Draw(img, boxImage.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

// 	return img
// }
