package day5

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gossie/adventofcode2022/util"
)

type instruction struct {
	from   int
	to     int
	amount int
}

func lineLength(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return len(scanner.Text())
}

func readData(filename string, stacks chan<- []util.Stack[string], instructions chan<- instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	numberOfStacks := (lineLength(filename) + 1) / 4
	importedStacks := make([]util.Stack[string], numberOfStacks)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			instructionComponents := strings.Split(line, " ")
			from, _ := strconv.Atoi(instructionComponents[3])
			to, _ := strconv.Atoi(instructionComponents[5])
			amount, _ := strconv.Atoi(instructionComponents[1])
			instructions <- instruction{from - 1, to - 1, amount}
		} else if line != "" {
			for i := 0; i < numberOfStacks; i++ {
				start := (i * 4) + 1
				end := start + 1
				cargo := line[start:end]
				if strings.TrimSpace(cargo) != "" {
					_, err := strconv.Atoi(cargo)
					if err != nil {
						importedStacks[i].PushLast(cargo)
					}
				}
			}
		}

		if line == "" {
			stacks <- importedStacks
			close(stacks)
		}
	}
	close(instructions)
}

func toResult(stacks []util.Stack[string]) string {
	result := ""
	for _, s := range stacks {
		result += s[0]
	}
	return result
}

func Crates9000(filename string) string {
	stacks := make(chan []util.Stack[string], 1)
	instructions := make(chan instruction, 10)

	go readData(filename, stacks, instructions)

	importedStacks := <-stacks

	images := make([]*image.Paletted, 0)
	images = append(images, createImage(importedStacks))

	for instruct := range instructions {
		for i := 0; i < instruct.amount; i++ {
			cargo, err := importedStacks[instruct.from].Pop()
			if err != nil {
				panic(err)
			}
			importedStacks[instruct.to].Push(cargo)
			images = append(images, createImage(importedStacks))
		}
	}
	createGif("day5-task1.gif", images)

	return toResult(importedStacks)
}

func Crates9001(filename string) string {
	stacks := make(chan []util.Stack[string], 1)
	instructions := make(chan instruction, 1)

	go readData(filename, stacks, instructions)

	importedStacks := <-stacks

	images := make([]*image.Paletted, 0)
	images = append(images, createImage(importedStacks))

	for instruct := range instructions {
		cargo, err := importedStacks[instruct.from].PopMultiple(instruct.amount)
		if err != nil {
			panic(err)
		}
		importedStacks[instruct.to].PushAll(cargo)
		images = append(images, createImage(importedStacks))
	}
	createGif("day5-task2.gif", images)

	return toResult(importedStacks)
}

func createGif(name string, images []*image.Paletted) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("gif could not be created: ", err)
	}
	defer f.Close()

	delays := make([]int, len(images))
	for i := 0; i < len(delays); i++ {
		delays[i] = 5
	}
	delays[len(delays)-1] = 500

	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func createImage(event []util.Stack[string]) *image.Paletted {
	palette := []color.Color{
		color.White,
		color.Black,
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}

	highest := 0
	for _, stack := range event {
		if stack.Len() > highest {
			highest = stack.Len()
		}
	}

	blockWidth := 50
	blockHeight := 50

	img := image.NewPaletted(image.Rect(0, 0, 500, 2000), palette)

	for index, stack := range event {
		length := stack.Len()
		for i := length - 1; i >= 0; i-- {
			x := index * blockWidth
			y := 1999 - (length-i)*blockHeight
			box := image.Rect(x+1, y+1, x+blockWidth-2, y+blockHeight-2)
			//boxImage := img.SubImage(box)
			draw.Draw(img, box, &image.Uniform{color.RGBA{0x00, 0x00, 0xff, 0xff}}, image.Point{}, draw.Src)
			//draw.Draw(img, boxImage.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
		}
	}

	return img
}
