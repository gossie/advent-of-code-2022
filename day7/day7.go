package day7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size uint64
}

type directory struct {
	name        string
	files       []*file
	directories []*directory
	parent      *directory
}

func (d *directory) subDirectory(name string) *directory {
	for _, sub := range d.directories {
		if sub.name == name {
			return sub
		}
	}
	return nil
}

func (d *directory) size() uint64 {
	sum := uint64(0)
	for _, f := range d.files {
		sum += f.size
	}
	for _, subD := range d.directories {
		sum += subD.size()
	}
	return sum
}

func readData(filename string) *directory {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var currentDirectory *directory = nil
	var rootDirectory *directory = nil

	skipScan := false
	for skipScan || scanner.Scan() {
		skipScan = false
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd ") {
			parameter := line[5:]
			if parameter == ".." {
				if currentDirectory.parent != nil {
					currentDirectory = currentDirectory.parent
				}
			} else {
				if currentDirectory != nil {
					sub := currentDirectory.subDirectory(parameter)
					if sub != nil {
						currentDirectory = sub
					} else {
						currentDirectory = &directory{name: parameter, parent: currentDirectory}
					}
				} else {
					currentDirectory = &directory{name: parameter}
					rootDirectory = currentDirectory
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			subDirectories, files := readDirectoryContent(scanner, currentDirectory)
			currentDirectory.files = files
			currentDirectory.directories = subDirectories

			skipScan = true // beacause in readDirectoryContent a scan already happened
		}
	}
	return rootDirectory
}

func readDirectoryContent(scanner *bufio.Scanner, parent *directory) ([]*directory, []*file) {
	subDirectories := make([]*directory, 0)
	files := make([]*file, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			break
		} else if strings.HasPrefix(line, "dir ") {
			subDirectories = append(subDirectories, &directory{name: line[4:], parent: parent})
		} else {
			fileInfo := strings.Split(line, " ")
			name := fileInfo[1]
			size, _ := strconv.Atoi(fileInfo[0])
			files = append(files, &file{name, uint64(size)})
		}
	}
	return subDirectories, files
}

func Part1(filename string) uint64 {
	root := readData(filename)
	sum := uint64(0)

	queue := make([]*directory, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = append(queue[1:], current.directories...)
		size := current.size()
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func Part2(filename string) uint64 {
	root := readData(filename)
	var result uint64 = math.MaxUint64

	diskSpace := uint64(70000000)
	neededSpace := uint64(30000000)

	availableSpace := diskSpace - root.size()

	queue := make([]*directory, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = append(queue[1:], current.directories...)
		size := current.size()
		if (availableSpace+size) > neededSpace && size < result {
			result = size
		}
	}

	return result
}
