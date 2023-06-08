package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %s\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %s\n", solvePart2(readFile("input")))
}

func solvePart1(input string) string {
	return solveWithStrategy(input, CrateMover9000{})
}

func solvePart2(input string) string {
	return solveWithStrategy(input, CrateMover9001{})
}

func solveWithStrategy(input string, crateMover CrateMover) string {
	parts := strings.SplitN(input, "\n\n", 2)
	firstHalf := parts[0]
	secondHalf := parts[1]

	stackHeader := strings.Split(firstHalf, "\n")[len(strings.Split(firstHalf, "\n"))-1]
	numStacks := getNumStacks(stackHeader)
	stackRows := strings.Split(strings.TrimSpace(firstHalf), "\n")
	stackRows = stackRows[:len(stackRows)-1]
	stacks := parseStacks(numStacks, stackRows)

	stepRows := strings.Split(strings.TrimSpace(secondHalf), "\n")
	steps := parseSteps(stepRows)

	stacks = crateMover.Operate(steps, stacks)

	topCrates := ""
	for _, stack := range stacks {
		topCrates += string(stack[len(stack)-1])
	}
	return topCrates
}

func getNumStacks(stackHeader string) int {
	words := strings.Fields(stackHeader)
	lastWord := words[len(words)-1]
	numStacks, _ := strconv.Atoi(lastWord)
	return numStacks
}

func parseStacks(numStacks int, stackRows []string) [][]rune {
	stacks := make([][]rune, numStacks)

	for stackIndex := 0; stackIndex < numStacks; stackIndex++ {
		for _, stackRow := range stackRows {
			pos := 1 + stackIndex*4
			if pos < len(stackRow) {
				c := rune(stackRow[pos])
				if c != ' ' {
					stacks[stackIndex] = append(stacks[stackIndex], c)
				}
			}
		}
	}

	return stacks
}

func parseSteps(stepRows []string) []Step {
	steps := make([]Step, len(stepRows))

	for i, row := range stepRows {
		step := parseStep(row)
		steps[i] = step
	}

	return steps
}

func parseStep(row string) Step {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindStringSubmatch(row)

	count, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return Step{Count: count, From: from, To: to}
}

type Step struct {
	Count int
	From  int
	To    int
}

type CrateMover interface {
	Operate(steps []Step, stacks [][]rune) [][]rune
}

type CrateMover9000 struct{}

func (cm CrateMover9000) Operate(steps []Step, stacks [][]rune) [][]rune {
	stacksCopy := copyStacks(stacks)

	for _, step := range steps {
		for i := 0; i < step.Count; i++ {
			t := stacksCopy[step.From-1][len(stacksCopy[step.From-1])-1]
			stacksCopy[step.From-1] = stacksCopy[step.From-1][:len(stacksCopy[step.From-1])-1]
			stacksCopy[step.To-1] = append(stacksCopy[step.To-1], t)
		}
	}

	return stacksCopy
}

type CrateMover9001 struct{}

func (cm CrateMover9001) Operate(steps []Step, stacks [][]rune) [][]rune {
	stacksCopy := copyStacks(stacks)

	for _, step := range steps {
		temp := []rune{}
		for i := 0; i < step.Count; i++ {
			t := stacksCopy[step.From-1][len(stacksCopy[step.From-1])-1]
			stacksCopy[step.From-1] = stacksCopy[step.From-1][:len(stacksCopy[step.From-1])-1]
			temp = append(temp, t)
		}
		for i := 0; i < step.Count; i++ {
			t := temp[len(temp)-1]
			temp = temp[:len(temp)-1]
			stacksCopy[step.To-1] = append(stacksCopy[step.To-1], t)
		}
	}

	return stacksCopy
}

func copyStacks(stacks [][]rune) [][]rune {
	stacksCopy := make([][]rune, len(stacks))
	for i, stack := range stacks {
		stacksCopy[i] = make([]rune, len(stack))
		copy(stacksCopy[i], stack)
	}
	return stacksCopy
}

func readFile(path string) string {
	content, _ := ioutil.ReadFile(path)
	return string(content)
}
