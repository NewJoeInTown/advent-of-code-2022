package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ','
		})

		a := parseUint8(parts[0])
		b := parseUint8(parts[1])
		c := parseUint8(parts[2])
		d := parseUint8(parts[3])

		if isSubset(a, b, c, d) || isSubset(c, d, a, b) {
			count++
		}
	}

	return count
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ','
		})

		a := parseUint8(parts[0])
		b := parseUint8(parts[1])
		c := parseUint8(parts[2])
		d := parseUint8(parts[3])

		if !isDisjoint(a, b, c, d) {
			count++
		}
	}

	return count
}

func isSubset(a, b, c, d uint8) bool {
	for i := a; i <= b; i++ {
		if i >= c && i <= d {
			return true
		}
	}
	return false
}

func isDisjoint(a, b, c, d uint8) bool {
	return b < c || d < a
}

func parseUint8(s string) uint8 {
	var val uint8
	fmt.Sscanf(s, "%d", &val)
	return val
}

func readFile(path string) string {
	content, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(content))
}
