package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	X, Y int
}

type Map [][]byte

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) int {
	mapData := parseMap(input)
	start := findCoordinate(mapData, 'S')
	end := findCoordinate(mapData, 'E')

	successors := func(c Coord) []Coord {
		return getNeighbors(mapData, c)
	}

	path := bfs(start, successors, func(c Coord) bool {
		return c == end
	})

	numSteps := len(path) - 1
	return numSteps
}

func solvePart2(input string) int {
	mapData := parseMap(input)
	end := findCoordinate(mapData, 'E')

	successors := func(c Coord) []Coord {
		return getNeighbors(mapData, c)
	}

	paths := make([][]Coord, 0)
	xLen := len(mapData[0])
	yLen := len(mapData)

	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			if mapData[y][x] == 'a' {
				p := bfs(Coord{x, y}, successors, func(c Coord) bool {
					return c == end
				})
				paths = append(paths, p)
			}
		}
	}

	shortestPath := paths[0]
	for _, path := range paths {
		if len(path) < len(shortestPath) {
			shortestPath = path
		}
	}

	numSteps := len(shortestPath) - 1
	return numSteps
}

func parseMap(input string) Map {
	lines := strings.Split(input, "\n")
	mapData := make(Map, len(lines))

	for i, line := range lines {
		mapData[i] = []byte(line)
	}

	return mapData
}

func findCoordinate(mapData Map, target byte) Coord {
	for y, row := range mapData {
		for x, value := range row {
			if value == target {
				return Coord{x, y}
			}
		}
	}

	return Coord{-1, -1}
}

func getNeighbors(mapData Map, coord Coord) []Coord {
	x, y := coord.X, coord.Y
	xLen := len(mapData[0])
	yLen := len(mapData)

	neighbors := make([]Coord, 0)

	if isValid(x-1, y, xLen, yLen) {
		neighbors = append(neighbors, Coord{x - 1, y})
	}
	if isValid(x+1, y, xLen, yLen) {
		neighbors = append(neighbors, Coord{x + 1, y})
	}
	if isValid(x, y-1, xLen, yLen) {
		neighbors = append(neighbors, Coord{x, y - 1})
	}
	if isValid(x, y+1, xLen, yLen) {
		neighbors = append(neighbors, Coord{x, y + 1})
	}

	return neighbors
}

func isValid(x, y, xLen, yLen int) bool {
	return x >= 0 && x < xLen && y >= 0 && y < yLen
}

func readFile(path string) string {
	data, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(data))
}

func bfs(start Coord, successors func(Coord) []Coord, isGoal func(Coord) bool) []Coord {
	visited := make(map[Coord]bool)
	queue := [][]Coord{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1]
		if isGoal(current) {
			return path
		}

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, successor := range successors(current) {
			newPath := append(path, successor)
			queue = append(queue, newPath)
		}
	}

	return nil
}
