package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type FileOrDir struct {
	Size uint32
	Name string
}

type Command int

const (
	Cd Command = iota
	CdUp
	CdRoot
	Ls
)

type Entry struct {
	Command Command
	Output []FileOrDir
}

func main() {
	fmt.Printf("Part 1: %d\n", solvePart1(readFile("input")))
	fmt.Printf("Part 2: %d\n", solvePart2(readFile("input")))
}

func solvePart1(input string) uint32 {
	entries := parseEntries(input)
	dirs := calculateDirectorySizes(entries)
	dirsIncl := aggregateDirectorySizes(dirs)
	var sum uint32
	for _, size := range dirsIncl {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func solvePart2(input string) uint32 {
	entries := parseEntries(input)
	dirs := calculateDirectorySizes(entries)
	dirsIncl := aggregateDirectorySizes(dirs)

	totalSpace := uint32(70000000)
	usedSpace := dirsIncl["/"]
	availableSpace := totalSpace - usedSpace
	neededAvailableSpace := uint32(30000000)
	neededAdditionalAvailableSpace := neededAvailableSpace - availableSpace

	var answer uint32
	for _, space := range dirsIncl {
		if space >= neededAdditionalAvailableSpace {
			answer = space
			break
		}
	}
	return answer
}

func parseEntries(input string) []Entry {
	lines := strings.Split(input, "\n")
	entries := make([]Entry, 0)
	lastFileOrDir := make([]FileOrDir, 0)

	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			if len(lastFileOrDir) > 0 {
				entries = append(entries, Entry{Output: lastFileOrDir})
				lastFileOrDir = make([]FileOrDir, 0)
			}

			cmd := line[2:]
			command := parseCommand(cmd)
			entries = append(entries, Entry{Command: command})
		} else {
			fileOrDir := parseFileOrDir(line)
			lastFileOrDir = append(lastFileOrDir, fileOrDir)
		}
	}

	if len(lastFileOrDir) > 0 {
		entries = append(entries, Entry{Output: lastFileOrDir})
	}

	return entries
}

func calculateDirectorySizes(entries []Entry) map[string]uint32 {
	dirSizes := make(map[string]uint32)
	currDir := make([]string, 0)

	for _, entry := range entries {
		switch entry.Command {
		case Cd:
			currDir = append(currDir, entry.Output[0].Name)
		case CdUp:
			currDir = currDir[:len(currDir)-1]
		case CdRoot:
			currDir = make([]string, 0)
		case Ls:
		}

		if len(entry.Output) > 0 {
			for _, output := range entry.Output {
				switch output.Name[0:3] {
				case "dir":
					dirPath := "/" + strings.Join(currDir, "/")
					dirSizes[dirPath] += 0
				default:
					filePath := "/" + strings.Join(currDir, "/")
					dirSizes[filePath] += output.Size
				}
			}
		}
	}

	return dirSizes
}

func aggregateDirectorySizes(sizes map[string]uint32) map[string]uint32 {
	aggrSizes := make(map[string]uint32)

	for path := range sizes {
		aggrSize := uint32(0)
		for key := range sizes {
			if strings.HasPrefix(key, path) {
				aggrSize += sizes[key]
			}
		}
		aggrSizes[path] = aggrSize
	}

	return aggrSizes
}

func parseCommand(cmd string) Command {
	switch cmd {
	case "ls":
		return Ls
	case "cd ..":
		return CdUp
	case "cd /":
		return CdRoot
	default:
		return Cd
	}
}

func parseFileOrDir(str string) FileOrDir {
	if strings.HasPrefix(str, "dir") {
		name := str[4:]
		return FileOrDir{Name: name}
	} else {
		parts := strings.Fields(str)
		size, _ := strconv.Atoi(parts[0])
		name := parts[1]
		return FileOrDir{Size: uint32(size), Name: name}
	}
}

func readFile(path string) string {
	content, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(content))
}
