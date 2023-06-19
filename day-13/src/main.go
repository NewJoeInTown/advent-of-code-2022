package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) int {
	packets := parsePackets(input)

	leftSmallerIndices := make([]int, 0)
	for i, packet := range packets {
		if packet[0] < packet[1] {
			leftSmallerIndices = append(leftSmallerIndices, i+1)
		}
	}

	sum := 0
	for _, index := range leftSmallerIndices {
		sum += index
	}

	return sum
}

func solvePart2(input string) int {
	dividerPacket1 := parsePacket("[[2]]")
	dividerPacket2 := parsePacket("[[6]]")

	packets := parsePackets(input)
	packets = append(packets, dividerPacket1, dividerPacket2)

	sort.Slice(packets, func(i, j int) bool {
		return comparePackets(packets[i], packets[j]) < 0
	})

	dividerPacketIndices := make([]int, 0)
	for i, packet := range packets {
		if packet == dividerPacket1 || packet == dividerPacket2 {
			dividerPacketIndices = append(dividerPacketIndices, i+1)
		}
	}

	product := 1
	for _, index := range dividerPacketIndices {
		product *= index
	}

	return product
}

func parsePackets(input string) [][]int {
	lines := strings.Split(input, "\n\n")
	packets := make([][]int, len(lines))

	for i, line := range lines {
		pair := strings.Split(line, "\n")
		a, _ := strconv.Atoi(pair[0])
		b, _ := strconv.Atoi(pair[1])
		packets[i] = []int{a, b}
	}

	return packets
}

func parsePacket(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	packet := make([]int, len(matches))

	for i, match := range matches {
		packet[i], _ = strconv.Atoi(match)
	}

	return packet
}

func comparePackets(packet1, packet2 []int) int {
	if packet1[0] < packet2[0] {
		return -1
	} else if packet1[0] > packet2[0] {
		return 1
	} else {
		return comparePacketsTail(packet1[1:], packet2[1:])
	}
}

func comparePacketsTail(packet1, packet2 []int) int {
	if len(packet1) == 0 && len(packet2) == 0 {
		return 0
	} else if len(packet1) == 0 {
		return -1
	} else if len(packet2) == 0 {
		return 1
	} else if packet1[0] < packet2[0] {
		return -1
	} else if packet1[0] > packet2[0] {
		return 1
	} else {
		return comparePacketsTail(packet1[1:], packet2[1:])
	}
}

func readFile(path string) string {
	data, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(data))
}
