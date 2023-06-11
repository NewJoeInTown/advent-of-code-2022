package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(file("input")))
	fmt.Printf("Part 2:\n%s\n", strings.ReplaceAll(solvePart2(file("input")), ".", " "))
}

func solvePart1(input string) int {
	xValues := getXValues(strings.Split(input, "\n"))
	signalStrengths := make([]int, len(xValues))
	for c, x := range xValues {
		signalStrengths[c] = (c + 1) * x
	}
	interesting := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, c := range interesting {
		sum += signalStrengths[c-1]
	}
	return sum
}

func solvePart2(input string) string {
	xValues := getXValues(strings.Split(input, "\n"))
	pixels := getPixels(xValues)
	display := make([]byte, len(pixels))
	for i, p := range pixels {
		if p {
			display[i] = '#'
		} else {
			display[i] = '.'
		}
	}
	chunks := chunkString(string(display), 40)
	return strings.Join(chunks, "\n")
}

func getXValues(lines []string) []int {
	x := 1
	cycleVec := make([]int, 0, len(lines))

	for _, l := range lines {
		switch l {
		case "noop":
			cycleVec = append(cycleVec, x)
		default:
			split := strings.SplitN(l, " ", 2)
			oInt, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			if split[0] != "addx" {
				panic("Invalid instruction")
			}
			cycleVec = append(cycleVec, x, x)
			x += oInt
		}
	}

	return cycleVec
}

func getPixels(xValues []int) []bool {
	spriteVec := make([][]int, len(xValues))
	for i, x := range xValues {
		spriteVec[i] = []int{x - 1, x, x + 1}
	}

	cycles := make([]int, 240)
	for i := 0; i < 240; i++ {
		cycles[i] = i + 1
	}

	pixels := make([]bool, len(cycles))
	for i, cycle := range cycles {
		a, b, c := spriteVec[i][0], spriteVec[i][1], spriteVec[i][2]
		cycle = cycle - 1
		cycle = cycle % 40
		pixels[i] = cycle == a || cycle == b || cycle == c
	}
	return pixels
}

func file(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

func chunkString(s string, chunkSize int) []string {
	chunks := make([]string, 0, (len(s)+chunkSize-1)/chunkSize)
	for len(s) >= chunkSize {
		chunks = append(chunks, s[:chunkSize])
		s = s[chunkSize:]
	}
	if len(s) > 0 {
		chunks = append(chunks, s)
	}
	return chunks
}
