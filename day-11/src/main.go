package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type WorryReductionMethod int

const (
	DivBy3 WorryReductionMethod = iota
	ModuloBySumOfAllTestDivisors
)

type Operator int

const (
	Add Operator = iota
	Times
)

type Operand interface {
	isOperand()
}

type Old struct{}

func (o Old) isOperand() {}

type Number struct {
	value uint64
}

func (o Number) isOperand() {}

type Monkey struct {
	items        []uint64
	op           Operator
	opa          Operand
	testDiv      uint64
	trueDest     uint64
	falseDest    uint64
	inspectCount uint64
}

func solvePart1(input string) uint64 {
	return solve(input, 20, DivBy3, false)
}

func solvePart2(input string) uint64 {
	return solve(input, 10000, ModuloBySumOfAllTestDivisors, false)
}

func solve(input string, numRounds uint64, worryReductionMethod WorryReductionMethod, shouldLog bool) uint64 {
	monkeys := make([]Monkey, 0)

	monkeyStrings := strings.Split(input, "\n\n")
	for _, monkeyStr := range monkeyStrings {
		monkey, _ := parseMonkey(monkeyStr)
		monkeys = append(monkeys, monkey)
	}

	for round := uint64(1); round <= numRounds; round++ {
		for mi := range monkeys {
			if shouldLog {
				fmt.Printf("Monkey %d:\n", mi)
			}

			reverseItems(&monkeys[mi].items)

			for len(monkeys[mi].items) > 0 {
				i := monkeys[mi].items[len(monkeys[mi].items)-1]
				monkeys[mi].items = monkeys[mi].items[:len(monkeys[mi].items)-1]

				if shouldLog {
					fmt.Printf("  Monkey inspect an item with a worry level of %d.\n", i)
				}
				monkeys[mi].inspectCount++

				opa := resolveOperand(monkeys[mi].opa)
				higherWorry := uint64(0)
				if monkeys[mi].op == Add {
					newI := i + opa
					if shouldLog {
						fmt.Printf("    Worry level is increased by %d to %d.\n", opa, newI)
					}
					higherWorry = newI
				} else if monkeys[mi].op == Times {
					newI := i * opa
					if shouldLog {
						fmt.Printf("    Worry level is multiplied by %d to %d.\n", opa, newI)
					}
					higherWorry = newI
				}

				maybeLowerWorry := uint64(0)
				if worryReductionMethod == DivBy3 {
					div := uint64(3)
					temp := higherWorry / div
					if shouldLog {
						fmt.Printf("    Monkey gets bored with item. Worry level is divided by %d to %d.\n", div, temp)
					}
					maybeLowerWorry = temp
				} else if worryReductionMethod == ModuloBySumOfAllTestDivisors {
					modulo := calculateModulo(monkeys)
					temp := higherWorry % modulo
					if shouldLog {
						fmt.Printf("    Monkey gets bored with item. Worry level is moduloed by %d to %d.\n", modulo, temp)
					}
					maybeLowerWorry = temp
				}

				if maybeLowerWorry == 0 {
					panic("Invalid maybeLowerWorry")
				}

				isDiv := maybeLowerWorry%monkeys[mi].testDiv == 0
				dest := uint64(0)
				if isDiv {
					if shouldLog {
						fmt.Printf("    Current worry level is divisible by %d.\n", monkeys[mi].testDiv)
						fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", maybeLowerWorry, monkeys[mi].trueDest)
					}
					dest = monkeys[mi].trueDest
				} else {
					if shouldLog {
						fmt.Printf("    Current worry level is not divisible by %d.\n", monkeys[mi].testDiv)
						fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", maybeLowerWorry, monkeys[mi].falseDest)
					}
					dest = monkeys[mi].falseDest
				}

				monkeys[dest].items = append(monkeys[dest].items, maybeLowerWorry)
			}

			monkeys[mi].items = make([]uint64, 0)
		}

		if shouldLog {
			fmt.Printf("\nAfter round %d, the monkeys are holding items with these worry levels:\n", round)
			for index, m := range monkeys {
				items := joinUint64s(m.items, ", ")
				fmt.Printf("Monkey %d: %s\n", index, items)
			}
		}
	}

	if shouldLog {
		fmt.Println()
		for index, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times.\n", index, m.inspectCount)
		}
		fmt.Println()
	}

	inspectCounts := make([]uint64, len(monkeys))
	for _, m := range monkeys {
		inspectCounts = append(inspectCounts, m.inspectCount)
	}

	return productOfTopTwo(inspectCounts)
}

func reverseItems(items *[]uint64) {
	for i, j := 0, len(*items)-1; i < j; i, j = i+1, j-1 {
		(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
	}
}

func resolveOperand(opa Operand) uint64 {
	switch opa.(type) {
	case Old:
		return 0
	case Number:
		return opa.(Number).value
	default:
		panic("Unknown Operand type")
	}
}

func calculateModulo(monkeys []Monkey) uint64 {
	modulo := uint64(1)
	for _, m := range monkeys {
		modulo *= m.testDiv
	}
	return modulo
}

func parseMonkey(input string) (Monkey, error) {
	re := regexp.MustCompile(`^Monkey (\d+):\n.+Starting items: ([\d, ]+)\n.+Operation: new = old ([*+]) (\d+|old)\n.+Test: divisible by (\d+)\n.+If true: throw to monkey (\d+)\n.+If false: throw to monkey (\d+)$`)
	match := re.FindStringSubmatch(input)
	if len(match) != 8 {
		return Monkey{}, fmt.Errorf("Failed to parse monkey: %s", input)
	}

	monkey := Monkey{}
	monkey.items = parseUint64s(match[2])
	monkey.op = parseOperator(match[3])
	monkey.opa = parseOperand(match[4])
	monkey.testDiv = parseUint64(match[5])
	monkey.trueDest = parseUint64(match[6])
	monkey.falseDest = parseUint64(match[7])
	monkey.inspectCount = 0

	return monkey, nil
}

func parseUint64s(input string) []uint64 {
	items := strings.Split(input, ", ")
	parsedItems := make([]uint64, 0, len(items))
	for _, item := range items {
		value, _ := strconv.ParseUint(item, 10, 64)
		parsedItems = append(parsedItems, value)
	}
	return parsedItems
}

func parseUint64(input string) uint64 {
	value, _ := strconv.ParseUint(input, 10, 64)
	return value
}

func parseOperator(input string) Operator {
	if input == "+" {
		return Add
	} else if input == "*" {
		return Times
	}
	panic("Unknown operator")
}

func parseOperand(input string) Operand {
	if input == "old" {
		return Old{}
	} else {
		value, _ := strconv.ParseUint(input, 10, 64)
		return Number{value: value}
	}
}

func joinUint64s(items []uint64, separator string) string {
	strItems := make([]string, 0, len(items))
	for _, item := range items {
		strItems = append(strItems, strconv.FormatUint(item, 10))
	}
	return strings.Join(strItems, separator)
}

func productOfTopTwo(items []uint64) uint64 {
	if len(items) < 2 {
		panic("Insufficient items")
	}

	sortDescending(items)
	return items[0] * items[1]
}

func sortDescending(items []uint64) {
	maxIdx := 0
	for i := 1; i < len(items); i++ {
		if items[i] > items[maxIdx] {
			maxIdx = i
		}
	}
	items[0], items[maxIdx] = items[maxIdx], items[0]
}

func readFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func main() {
	input, _ := readFile("input")
	fmt.Printf("Part 1: %d\n", solvePart1(input))
	fmt.Printf("Part 2: %d\n", solvePart2(input))
}
