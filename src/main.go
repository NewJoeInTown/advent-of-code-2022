package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(file("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(file("input")))
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		values := parseIntTuple(line)
		a, b, c, d := values[0], values[1], values[2], values[3]
		h1 := createIntSet(a, b)
		h2 := createIntSet(c, d)
		if isSubsetOrEqual(h1, h2) || isSubsetOrEqual(h2, h1) {
			count++
		}
	}
	return count
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		values := parseIntTuple(line)
		a, b, c, d := values[0], values[1], values[2], values[3]
		h1 := createIntSet(a, b)
		h2 := createIntSet(c, d)
		if !isDisjoint(h1, h2) {
			count++
		}
	}
	return count
}

func parseIntTuple(s string) [4]int {
	parts := strings.Split(s, "-")
	values := [4]int{}
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		values[i] = num
	}
	return values
}

func createIntSet(start, end int) map[int]struct{} {
	set := make(map[int]struct{})
	for i := start; i <= end; i++ {
		set[i] = struct{}{}
	}
	return set
}

func isSubsetOrEqual(h1, h2 map[int]struct{}) bool {
	for k := range h1 {
		if _, ok := h2[k]; !ok {
			return false
		}
	}
	return true
}

func isDisjoint(h1, h2 map[int]struct{}) bool {
	for k := range h1 {
		if _, ok := h2[k]; ok {
			return false
		}
	}
	return true
}

func file(path string) string {
	data, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(data))
}
