package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readFile("input")

	fmt.Printf("Part 1: %d\n", solvePart1(input))
	fmt.Printf("Part 2: %d\n", solvePart2(input))
}

func solvePart1(input string) int {
	return solveWithWindowSize(input, 4)
}

func solvePart2(input string) int {
	return solveWithWindowSize(input, 14)
}

func solveWithWindowSize(input string, windowSize int) int {
	runes := []rune(input)
	for i := 0; i <= len(runes)-windowSize; i++ {
		window := runes[i : i+windowSize]
		if allUnique(window) {
			return i + windowSize
		}
	}
	panic("No solution found")
}

func allUnique(window []rune) bool {
	seen := make(map[rune]bool)
	for _, r := range window {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func readFile(path string) string {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(input))
}
