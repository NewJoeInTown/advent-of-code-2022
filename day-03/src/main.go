package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(file("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(file("input")))
}

func solvePart1(input string) uint32 {
	lines := strings.Split(input, "\n")
	sum := uint32(0)
	for _, line := range lines {
		s1, s2 := line[:len(line)/2], line[len(line)/2:]
		commonChar := findCommonCharacter(s1, s2)
		priority := getPriority(commonChar)
		sum += priority
	}
	return sum
}

func solvePart2(input string) uint32 {
	lines := strings.Split(input, "\n")
	sum := uint32(0)
	for i := 0; i < len(lines)-2; i += 3 {
		s1, s2, s3 := lines[i], lines[i+1], lines[i+2]
		commonChar := findCommonCharacterThreeStrings(s1, s2, s3)
		priority := getPriority(commonChar)
		sum += priority
	}
	return sum
}

func findCommonCharacter(s1, s2 string) byte {
	for _, c1 := range s1 {
		if strings.ContainsRune(s2, c1) {
			return byte(c1)
		}
	}
	return 0
}

func findCommonCharacterThreeStrings(s1, s2, s3 string) byte {
	for _, c1 := range s1 {
		if strings.ContainsRune(s2, c1) && strings.ContainsRune(s3, c1) {
			return byte(c1)
		}
	}
	return 0
}

func getPriority(c byte) uint32 {
	if c >= 'a' && c <= 'z' {
		return uint32(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return uint32(c-'A') + 27
	}
	return 0
}

func file(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
