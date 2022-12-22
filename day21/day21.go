package day21

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type value struct {
	op          string
	left, right *value
	number      int
}

func (v *value) String() string {
	switch v.op {
	case "num":
		return fmt.Sprintf("%d", v.number)
	case "var":
		return "x"
	default:
		return fmt.Sprintf("%v %v %v", v.left.String(), v.op, v.right.String())
	}
}

func (v *value) execute() (int, error) {
	switch v.op {
	default:
		return -1, errors.New("here is a var")
	case "num":
		return v.number, nil
	case "+":
		l, lErr := v.left.execute()
		if lErr != nil {
			return -1, errors.New("here is a var")
		}
		r, rErr := v.right.execute()
		if rErr != nil {
			return -1, errors.New("here is a var")
		}
		return l + r, nil
	case "*":
		l, lErr := v.left.execute()
		if lErr != nil {
			return -1, errors.New("here is a var")
		}
		r, rErr := v.right.execute()
		if rErr != nil {
			return -1, errors.New("here is a var")
		}
		return l * r, nil
	case "/":
		l, lErr := v.left.execute()
		if lErr != nil {
			return -1, errors.New("here is a var")
		}
		r, rErr := v.right.execute()
		if rErr != nil {
			return -1, errors.New("here is a var")
		}
		return l / r, nil
	case "-":
		l, lErr := v.left.execute()
		if lErr != nil {
			return -1, errors.New("here is a var")
		}
		r, rErr := v.right.execute()
		if rErr != nil {
			return -1, errors.New("here is a var")
		}
		return l - r, nil
	}
}

func (v *value) solveTo(result int) int {
	// fmt.Println("solving", v.String(), " to ", result)
	switch v.op {
	default:
		panic("implement me")
	case "var":
		return result
	case "num":
		if v.number != result {
			panic(fmt.Sprintf("%d cannot be %d", v.number, result))
		}
		return v.number
	case "+":
		l, lErr := v.left.execute()
		if lErr != nil {
			r, _ := v.right.execute()
			return v.left.solveTo(result - r)
		} else {
			return v.right.solveTo(result - l)
		}
	case "*":
		l, lErr := v.left.execute()
		if lErr != nil {
			r, _ := v.right.execute()
			return v.left.solveTo(result / r)
		} else {
			return v.right.solveTo(result / l)
		}
	case "/":
		l, lErr := v.left.execute()
		if lErr != nil {
			r, _ := v.right.execute()
			return v.left.solveTo(result * r)
		} else {
			return v.right.solveTo(l / result)
		}
	case "-":
		l, lErr := v.left.execute()
		if lErr != nil {
			r, _ := v.right.execute()
			return v.left.solveTo(result + r)
		} else {
			return v.right.solveTo(l - result)
		}
	}
}

func buildTree(nodeName string, valuesMap map[string]string, part2 bool) *value {
	valueStr := valuesMap[nodeName]
	v, err := strconv.Atoi(valueStr)
	if err != nil {
		opArr := strings.Split(valueStr, " ")
		return &value{op: opArr[1], left: buildTree(opArr[0], valuesMap, part2), right: buildTree(opArr[2], valuesMap, part2)}
	} else {
		if part2 && nodeName == "humn" {
			return &value{op: "var", number: v}
		}
		return &value{op: "num", number: v}
	}
}

func readData(filename string, part2 bool) *value {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	valuesMap := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, ": ")
		valuesMap[val[0]] = val[1]
	}

	return buildTree("root", valuesMap, part2)
}

func Part1(filename string) int {
	root := readData(filename, false)
	result, _ := root.execute()
	return result
}

func Part2(filename string) int {
	root := readData(filename, true)
	root.op = "="
	left, lErr := root.left.execute()
	if lErr != nil {
		right, _ := root.right.execute()
		return root.left.solveTo(right)
	} else {
		return root.right.solveTo(left)
	}
}
