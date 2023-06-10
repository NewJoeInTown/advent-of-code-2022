package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Trees [][]uint32

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) int {
	trees := parseTrees(input)

	assertEqual(len(trees), len(trees[0]))
	length := len(trees)

	visibleTrees := make([][2]int, 0)
	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			treeHeight := trees[y][x]

			visibleFromLeft := true
			for xx := 0; xx < x; xx++ {
				if trees[y][xx] >= treeHeight {
					visibleFromLeft = false
					break
				}
			}

			visibleFromRight := true
			for xx := x + 1; xx < length; xx++ {
				if trees[y][xx] >= treeHeight {
					visibleFromRight = false
					break
				}
			}

			visibleFromTop := true
			for yy := 0; yy < y; yy++ {
				if trees[yy][x] >= treeHeight {
					visibleFromTop = false
					break
				}
			}

			visibleFromBot := true
			for yy := y + 1; yy < length; yy++ {
				if trees[yy][x] >= treeHeight {
					visibleFromBot = false
					break
				}
			}

			if visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBot {
				visibleTrees = append(visibleTrees, [2]int{y, x})
			}
		}
	}

	return len(visibleTrees)
}

func solvePart2(input string) uint32 {
	trees := parseTrees(input)

	assertEqual(len(trees), len(trees[0]))
	length := len(trees)

	visibleTrees := make([]int, 0)
	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			checkView := func(rangeVals []int, f func(int) bool) int {
				count := 0
				for _, val := range rangeVals {
					count++
					if f(val) {
						break
					}
				}
				return count
			}

			viewUp := checkView(reverseRange(0, y), func(yy int) bool {
				return trees[yy][x] >= trees[y][x]
			})

			viewLeft := checkView(reverseRange(0, x), func(xx int) bool {
				return trees[y][xx] >= trees[y][x]
			})

			viewRight := checkView(rangeVals(x+1, length), func(xx int) bool {
				return trees[y][xx] >= trees[y][x]
			})

			viewDown := checkView(rangeVals(y+1, length), func(yy int) bool {
				return trees[yy][x] >= trees[y][x]
			})

			visibleTrees = append(visibleTrees, viewUp*viewLeft*viewRight*viewDown)
		}
	}

	maxValue := uint32(0)
	for _, value := range visibleTrees {
		if uint32(value) > maxValue {
			maxValue = uint32(value)
		}
	}

	return maxValue
}

func parseTrees(input string) Trees {
	lines := strings.Split(input, "\n")
	trees := make(Trees, len(lines))

	for i, line := range lines {
		numbers := strings.Split(line, "")
		trees[i] = make([]uint32, len(numbers))
		for j, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			trees[i][j] = uint32(num)
		}
	}

	return trees
}

func readFile(path string) string {
	content, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(content))
}

func assertEqual(a, b int) {
	if a != b {
		panic(fmt.Sprintf("Assertion failed: %d != %d", a, b))
	}
}

func reverseRange(start, end int) []int {
	result := make([]int, end-start)
	for i := 0; i < len(result); i++ {
		result[i] = end - i - 1
	}
	return result
}

func rangeVals(start, end int) []int {
	result := make([]int, end-start)
	for i := 0; i < len(result); i++ {
		result[i] = start + i
	}
	return result
}
