package day8

import (
	"bufio"
	"math"
	"os"
)

func readData(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, heigth := range line {
			row = append(row, int(heigth))
		}

		grid = append(grid, row)
	}

	return grid
}

func Part1(filename string) int {
	grid := readData(filename)

	visible := len(grid)*2 + (len(grid[0])-2)*2
	for y := 1; y < len(grid)-1; y++ {
		row := grid[y]
		for x := 1; x < len(row)-1; x++ {
			heigth := row[x]
			if isVisible(grid, y, x, heigth) {
				visible++
			}
		}

	}

	return visible
}

func isVisible(grid [][]int, posY, posX, heigth int) bool {
	return visibleFromTop(grid, posY, posX, heigth) || visibleFromBottom(grid, posY, posX, heigth) || visibleFromLeft(grid, posY, posX, heigth) || visibleFromRight(grid, posY, posX, heigth)
}

func visibleFromTop(grid [][]int, posY, posX, heigth int) bool {
	for y := posY - 1; y >= 0; y-- {
		if grid[y][posX] >= heigth {
			return false
		}
	}
	return true
}

func visibleFromBottom(grid [][]int, posY, posX, heigth int) bool {
	for y := posY + 1; y < len(grid); y++ {
		if grid[y][posX] >= heigth {
			return false
		}
	}
	return true
}

func visibleFromLeft(grid [][]int, posY, posX, heigth int) bool {
	row := grid[posY]
	for x := posX - 1; x >= 0; x-- {
		if row[x] >= heigth {
			return false
		}
	}
	return true
}

func visibleFromRight(grid [][]int, posY, posX, heigth int) bool {
	row := grid[posY]
	for x := posX + 1; x < len(row); x++ {
		if row[x] >= heigth {
			return false
		}
	}
	return true
}

func Part2(filename string) int {
	grid := readData(filename)

	highestScenicScore := 0

	for y := 1; y < len(grid)-1; y++ {
		row := grid[y]
		for x := 1; x < len(row)-1; x++ {
			currentScore := scenicScoreFromTop(grid, y, x, row[x]) * scenicScoreFromBottom(grid, y, x, row[x]) * scenicScoreFromLeft(grid, y, x, row[x]) * scenicScoreFromRight(grid, y, x, row[x])
			highestScenicScore = int(math.Max(float64(highestScenicScore), float64(currentScore)))
		}

	}

	return highestScenicScore
}

func scenicScoreFromTop(grid [][]int, posY, posX, heigth int) int {
	score := 0
	for y := posY - 1; y >= 0; y-- {
		if grid[y][posX] < heigth {
			score++
		} else {
			score++
			break
		}
	}

	return score
}

func scenicScoreFromBottom(grid [][]int, posY, posX, heigth int) int {
	score := 0
	for y := posY + 1; y < len(grid); y++ {
		if grid[y][posX] < heigth {
			score++
		} else {
			score++
			break
		}
	}

	return score
}

func scenicScoreFromLeft(grid [][]int, posY, posX, heigth int) int {
	score := 0
	row := grid[posY]
	for x := posX - 1; x >= 0; x-- {
		if row[x] < heigth {
			score++
		} else {
			score++
			break
		}
	}

	return score
}

func scenicScoreFromRight(grid [][]int, posY, posX, heigth int) int {
	score := 0
	row := grid[posY]
	for x := posX + 1; x < len(row); x++ {
		if row[x] < heigth {
			score++
		} else {
			score++
			break
		}
	}

	return score
}
