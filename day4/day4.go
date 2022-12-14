package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type assignmentRange struct {
	lowerBound int32
	upperBound int32
}

func (ar assignmentRange) subsumes(other assignmentRange) bool {
	return ar.lowerBound <= other.lowerBound && ar.upperBound >= other.upperBound
}

func (ar assignmentRange) intersectsWith(other assignmentRange) bool {
	return (ar.lowerBound >= other.lowerBound && ar.lowerBound <= other.upperBound) || (other.lowerBound >= ar.lowerBound && other.lowerBound <= ar.upperBound)
}

type pair struct {
	first  assignmentRange
	second assignmentRange
}

func (p pair) subsumes() bool {
	return p.first.subsumes(p.second) || p.second.subsumes(p.first)
}

func (p pair) intersects() bool {
	return p.first.intersectsWith(p.second)
}

type imageEvent struct {
	pair pair
}

func readData(filename string, pairs chan<- pair) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		firstRange := strings.Split(ranges[0], "-")
		firstLower, _ := strconv.Atoi(firstRange[0])
		firstUpper, _ := strconv.Atoi(firstRange[1])
		firstAssignemtRange := assignmentRange{int32(firstLower), int32(firstUpper)}

		secondRange := strings.Split(ranges[1], "-")
		secondLower, _ := strconv.Atoi(secondRange[0])
		secondUpper, _ := strconv.Atoi(secondRange[1])
		secondAssignemtRange := assignmentRange{int32(secondLower), int32(secondUpper)}

		pairs <- pair{firstAssignemtRange, secondAssignemtRange}
	}
	close(pairs)
}

func SubsumingPairs(filename string) int32 {
	pairs := make(chan pair, 10)
	// events := make(chan imageEvent, 10)
	// done := make(chan bool, 1)

	go readData(filename, pairs)
	// go createGif("day4-task1.gif", events, done)

	sum := int32(0)
	for p := range pairs {
		if p.subsumes() {
			sum++
		}
		// events <- imageEvent{p}
	}
	// close(events)

	// if <-done {
	// 	println("finshed rendering the gif")
	// } else {
	// 	println("gif could not be rendered")
	// }

	return sum
}

func OverlappingPairs(filename string) int32 {
	pairs := make(chan pair, 10)
	// events := make(chan imageEvent, 10)
	// done := make(chan bool, 1)

	go readData(filename, pairs)
	// go createGif("day4-task1.gif", events, done)

	sum := int32(0)
	for p := range pairs {
		if p.intersects() {
			sum++
		}
		// events <- imageEvent{p}
	}
	// close(events)

	// if <-done {
	// 	println("finshed rendering the gif")
	// } else {
	// 	println("gif could not be rendered")
	// }

	return sum
}

// func createGif(name string, events <-chan imageEvent, done chan<- bool) {
// 	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Println("gif could not be created: ", err)
// 		done <- false
// 	}
// 	defer f.Close()

// 	images := make([]*image.Paletted, 0)
// 	delays := make([]int, 0)
// 	for event := range events {
// 		images = append(images, createImage(event))
// 		delays = append(delays, 10)
// 	}

// 	gif.EncodeAll(f, &gif.GIF{
// 		Image: images,
// 		Delay: delays,
// 	})

// 	done <- true
// }

// func createImage(event imageEvent) *image.Paletted {
// 	palette := []color.Color{
// 		color.RGBA{0xff, 0xff, 0xff, 0xff},
// 		color.RGBA{0x00, 0x00, 0xff, 0xff},
// 		color.RGBA{0x00, 0xff, 0xff, 0xff},
// 		color.RGBA{0xff, 0x00, 0x00, 0xff},
// 	}

// 	height := 50
// 	img := image.NewPaletted(image.Rect(0, 0, 100, height), palette)

// 	for x := int(event.pair.first.lowerBound); x <= int(event.pair.first.upperBound); x++ {
// 		for y := 0; y < height; y++ {
// 			img.Set(x, y, color.RGBA{uint8(0), uint8(0), uint8(255), 0xff})
// 		}
// 	}

// 	for x := int(event.pair.second.lowerBound); x <= int(event.pair.second.upperBound); x++ {
// 		for y := 0; y < height; y++ {
// 			img.Set(x, y, color.RGBA{uint8(0), uint8(255), uint8(255), 0xff})
// 		}
// 	}

// 	for x := math.Max(float64(event.pair.second.lowerBound), float64(event.pair.first.lowerBound)); x <= math.Min(float64(event.pair.first.upperBound), float64(event.pair.second.upperBound)); x++ {
// 		for y := 0; y < height; y++ {
// 			img.Set(int(x), y, color.RGBA{uint8(255), uint8(0), uint8(0), 0xff})
// 		}
// 	}

// 	return img
// }
