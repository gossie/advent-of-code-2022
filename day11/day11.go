package day11

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

type item struct {
	level         int
	levelForPrime map[int]int
}

type operation struct {
	op     string
	factor int
}

func (o *operation) exec(item *item) {
	for index, prime := range primes {
		switch o.op {
		case "+":
			if o.factor < 0 {
				item.levelForPrime[prime] = (item.levelForPrime[prime] * 2) % prime
				if index == 0 {
					item.level += item.level
				}
			} else {
				item.levelForPrime[prime] = (item.levelForPrime[prime] + o.factor) % prime
				if index == 0 {
					item.level += o.factor
				}
			}
		case "*":
			if o.factor < 0 {
				item.levelForPrime[prime] = (item.levelForPrime[prime] * item.levelForPrime[prime]) % prime
				if index == 0 {
					item.level *= item.level
				}
			} else {
				item.levelForPrime[prime] = (item.levelForPrime[prime] * o.factor) % prime
				if index == 0 {
					item.level *= o.factor
				}
			}
		default:
			panic("unknown operator: " + o.op)
		}
	}
}

func (o *operation) test(item *item) bool {
	if o.op != "%" {
		panic("cannot perform test with " + o.op)
	}
	return item.level%o.factor == 0
}

func (o *operation) testWithoutRelief(item *item) bool {
	if o.op != "%" {
		panic("cannot perform test with " + o.op)
	}
	return item.levelForPrime[o.factor] == 0
}

type monkey struct {
	nr                   int
	items                []*item
	operation            *operation
	test                 *operation
	follower1, follower2 int
	inspectedItems       int
}

func readData(filename string) []*monkey {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	monkeys := make([]*monkey, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var currentMonkey *monkey = nil

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Monkey ") {
			if currentMonkey != nil {
				monkeys = append(monkeys, currentMonkey)
			}

			end := strings.IndexByte(line, ':')
			nr, _ := strconv.Atoi(line[7:end])
			currentMonkey = &monkey{nr: nr}
		}

		if strings.HasPrefix(line, "  Starting items: ") {
			startingItems := strings.Split(line[18:], ",")

			for _, i := range startingItems {
				newItem := item{levelForPrime: make(map[int]int)}
				level, _ := strconv.Atoi(strings.TrimSpace(i))
				newItem.level = level
				for _, prime := range primes {
					newItem.levelForPrime[prime] = newItem.level % prime
				}
				currentMonkey.items = append(currentMonkey.items, &newItem)
			}
		}

		if strings.HasPrefix(line, "  Operation: ") {
			function := line[13:]
			if strings.ContainsRune(function, '+') {
				summand, err := strconv.Atoi(strings.TrimSpace(strings.Split(function, "+")[1]))
				if err != nil {
					summand = -1
				}
				currentMonkey.operation = &operation{op: "+", factor: summand}
			} else if strings.ContainsRune(function, '*') {
				factor, err := strconv.Atoi(strings.TrimSpace(strings.Split(function, "*")[1]))
				if err != nil {
					factor = -1
				}
				currentMonkey.operation = &operation{op: "*", factor: factor}
			} else {
				panic("cannot parse operation")
			}
		}

		if strings.HasPrefix(line, "  Test: divisible by ") {
			test, _ := strconv.Atoi(line[21:])
			currentMonkey.test = &operation{op: "%", factor: test}
			scanner.Scan()
			x := strings.Split(scanner.Text(), " ")
			follower1, _ := strconv.Atoi(x[len(x)-1])

			scanner.Scan()
			x = strings.Split(scanner.Text(), " ")
			follower2, _ := strconv.Atoi(x[len(x)-1])

			currentMonkey.follower1 = follower1
			currentMonkey.follower2 = follower2
		}
	}
	monkeys = append(monkeys, currentMonkey)
	return monkeys
}

func MonkeyBusiness(filename string) int {
	monkeys := readData(filename)

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.operation.exec(item)
				monkey.inspectedItems++
				item.level /= 3
				if monkey.test.test(item) {
					monkeys[monkey.follower1].items = append(monkeys[monkey.follower1].items, item)
				} else {
					monkeys[monkey.follower2].items = append(monkeys[monkey.follower2].items, item)
				}
			}
			monkey.items = make([]*item, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems < monkeys[j].inspectedItems
	})

	return monkeys[len(monkeys)-1].inspectedItems * monkeys[len(monkeys)-2].inspectedItems
}

func MonkeyBusinessWithoutRelief(filename string) uint64 {
	monkeys := readData(filename)

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.operation.exec(item)
				monkey.inspectedItems++
				if monkey.test.testWithoutRelief(item) {
					monkeys[monkey.follower1].items = append(monkeys[monkey.follower1].items, item)
				} else {
					monkeys[monkey.follower2].items = append(monkeys[monkey.follower2].items, item)
				}
			}
			monkey.items = make([]*item, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems < monkeys[j].inspectedItems
	})

	return uint64(monkeys[len(monkeys)-1].inspectedItems) * uint64(monkeys[len(monkeys)-2].inspectedItems)
}
