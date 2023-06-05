package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := file("input")
	fmt.Printf("Part 1: %d\n", solvePart1(input))
	fmt.Printf("Part 2: %d\n", solvePart2(input))
}

func solvePart1(input string) uint32 {
	sections := strings.Split(input, "\n\n")
	maxSum := uint32(0)
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		sum := uint32(0)
		for _, line := range lines {
			num, _ := strconv.ParseUint(line, 10, 32)
			sum += uint32(num)
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func solvePart2(input string) uint32 {
	sections := strings.Split(input, "\n\n")
	sums := make([]uint32, len(sections))
	for i, section := range sections {
		lines := strings.Split(section, "\n")
		sum := uint32(0)
		for _, line := range lines {
			num, _ := strconv.ParseUint(line, 10, 32)
			sum += uint32(num)
		}
		sums[i] = sum
	}
	sortDescending(sums)
	totalSum := uint32(0)
	for _, sum := range sums[:3] {
		totalSum += sum
	}
	return totalSum
}

func sortDescending(slice []uint32) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] > slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func file(path string) string {
	data, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(data))
}
