package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) uint32 {
	lines := strings.Split(input, "\n")
	var totalScore uint32

	for _, line := range lines {
		parts := strings.Split(line, " ")
		elfShape := parseShape(parts[0])
		meShape := parseShape(parts[1])

		score := meShape.score() + forMe(elfShape, meShape).score()
		totalScore += score
	}

	return totalScore
}

func solvePart2(input string) uint32 {
	lines := strings.Split(input, "\n")
	var totalScore uint32

	for _, line := range lines {
		parts := strings.Split(line, " ")
		elfShape := parseShape(parts[0])
		outcomeType := parseOutcome(parts[1])

		score := outcomeType.withOther(elfShape).score() + outcomeType.score()
		totalScore += score
	}

	return totalScore
}

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

func parseShape(s string) Shape {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		log.Fatalf("Invalid shape: %s", s)
		return -1 // This line will never be reached, added to satisfy Go compiler
	}
}

func (s Shape) score() uint32 {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

func parseOutcome(s string) Outcome {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		log.Fatalf("Invalid outcome: %s", s)
		return -1 // This line will never be reached, added to satisfy Go compiler
	}
}

func forMe(other Shape, me Shape) Outcome {
	switch {
	case other == Rock && me == Rock:
		return Draw
	case other == Rock && me == Paper:
		return Win
	case other == Rock && me == Scissors:
		return Lose
	case other == Paper && me == Rock:
		return Lose
	case other == Paper && me == Paper:
		return Draw
	case other == Paper && me == Scissors:
		return Win
	case other == Scissors && me == Rock:
		return Win
	case other == Scissors && me == Paper:
		return Lose
	case other == Scissors && me == Scissors:
		return Draw
	default:
		return -1 // This line will never be reached, added to satisfy Go compiler
	}
}

func (o Outcome) withOther(other Shape) Shape {
	switch o {
	case Lose:
		switch other {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		default:
			return -1 // This line will never be reached, added to satisfy Go compiler
		}
	case Draw:
		return other
	case Win:
		switch other {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors
