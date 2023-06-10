package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	x, y int
}

type Motion struct {
	dir   Direction
	steps int
}

type Direction int

const (
	Right Direction = iota
	Up
	Down
	Left
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(file("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(file("input")))
}

func solvePart1(input string) int {
	motions := parseMotions(input)
	start := Coord{0, 4}
	finalVisited := make(map[Coord]struct{})
	finalVisited[start] = struct{}{}

	head, tail := start, start
	for _, motion := range motions {
		head, tail = moveHeadAndTail(head, tail, motion, finalVisited)
	}

	return len(finalVisited)
}

func solvePart2(input string) int {
	motions := parseMotions(input)
	start := Coord{0, 4}
	finalVisited := make(map[Coord]struct{})
	finalVisited[start] = struct{}{}

	knots := []Coord{start, start}
	for _, motion := range motions {
		knots = moveKnots(knots, motion, finalVisited)
	}

	return len(finalVisited)
}

func moveHeadAndTail(head, tail Coord, motion Motion, finalVisited map[Coord]struct{}) (Coord, Coord) {
	for i := 0; i < motion.steps; i++ {
		newHead := adjustCoord(head, motion.dir)
		newTail := adjustKnot(newHead, tail)
		finalVisited[newTail] = struct{}{}
		head, tail = newHead, newTail
	}

	return head, tail
}

func moveKnots(knots []Coord, motion Motion, finalVisited map[Coord]struct{}) []Coord {
	for i := 0; i < motion.steps; i++ {
		newKnots := moveRope(knots, motion.dir)
		newTail := newKnots[len(newKnots)-1]
		finalVisited[newTail] = struct{}{}
		knots = newKnots
	}

	return knots
}

func moveRope(knots []Coord, dir Direction) []Coord {
	newHead := adjustCoord(knots[0], dir)

	newKnots := make([]Coord, 0, len(knots)*3)
	newKnots = append(newKnots, newHead, newHead)
	newKnots = append(newKnots, knots...)
	newKnots = append(newKnots, knots...)

	for i := 0; i < len(newKnots)-1; i++ {
		newKnots[i] = adjustKnot(newKnots[i], newKnots[i+1])
	}

	return newKnots[:len(newKnots)-1]
}

func adjustKnot(previousKnot, knot Coord) Coord {
	xhd, yhd := previousKnot.x-knot.x, previousKnot.y-knot.y
	mhd := mhDist(previousKnot, knot)

	switch mhd {
	case 0, 1:
		return knot
	case 2:
		switch {
		case xhd == 0 && yhd == 2:
			return adjustCoord(knot, Down)
		case xhd == 0 && yhd == -2:
			return adjustCoord(knot, Up)
		case xhd == 2 && yhd == 0:
			return adjustCoord(knot, Right)
		case xhd == -2 && yhd == 0:
			return adjustCoord(knot, Left)
		default:
			return knot
		}
	case 3:
		switch {
		case xhd == 1 && yhd == 2:
			return adjustCoord(adjustCoord(knot, Right), Down)
		case xhd == 1 && yhd == -2:
			return adjustCoord(adjustCoord(knot, Right), Up)
		case xhd == -1 && yhd == 2:
			return adjustCoord(adjustCoord(knot, Left), Down)
		case xhd == -1 && yhd == -2:
			return adjustCoord(adjustCoord(knot, Left), Up)
		case xhd == 2 && yhd == 1:
			return adjustCoord(adjustCoord(knot, Right), Down)
		case xhd == 2 && yhd == -1:
			return adjustCoord(adjustCoord(knot, Right), Up)
		case xhd == -2 && yhd == 1:
			return adjustCoord(adjustCoord(knot, Left), Down)
		case xhd == -2 && yhd == -1:
			return adjustCoord(adjustCoord(knot, Left), Up)
		default:
			panic("unexpected case")
		}
	default:
		panic("unreachable")
	}
}

func adjustCoord(coord Coord, dir Direction) Coord {
	switch dir {
	case Right:
		return Coord{coord.x + 1, coord.y}
	case Up:
		return Coord{coord.x, coord.y - 1}
	case Down:
		return Coord{coord.x, coord.y + 1}
	case Left:
		return Coord{coord.x - 1, coord.y}
	default:
		panic("unexpected direction")
	}
}

func mhDist(a, b Coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseMotions(input string) []Motion {
	lines := strings.Split(input, "\n")
	motions := make([]Motion, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		steps := parseInt(parts[1])
		dir := parseDirection(parts[0])
		motions[i] = Motion{dir, steps}
	}

	return motions
}

func parseInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		panic(err)
	}
	return n
}

func parseDirection(s string) Direction {
	switch s {
	case "R":
		return Right
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	default:
		panic("unexpected direction")
	}
}

func file(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
