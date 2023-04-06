package day7

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var patternCD = regexp.MustCompile(`^\$ cd (.+)`)
var patternFolder = regexp.MustCompile(`^dir (.+)`)
var patternFile = regexp.MustCompile(`^([0-9]+) ([a-zA-Z.]+)`)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var cwd *file
		var root *file = newFolder("/")
		var lines []string = reader.ReadLines("day7/input.txt")

		for _, line := range lines {
			var log logLine = parseLine(line)

			if log.isCd {
				name := log.arguments[0]
				if name == "/" {
					cwd = root
				} else {
					cwd = cwd.children[name]
				}
			} else if log.isFolder {
				name := log.arguments[0]
				fldr := newFolder(cwd.path + name + "/")
				fldr.children[".."] = cwd
				cwd.children[name] = fldr
			} else if log.isFile {
				size, _ := strconv.Atoi(log.arguments[0])
				name := log.arguments[1]
				file := file{path: cwd.path + name, size: size}
				cwd.children[name] = &file
			}
		}

		var folderTable map[string]int = map[string]int{}
		buildFolderTable(root, folderTable)

		var p1 string = partOne(folderTable)
		var p2 string = partTwo(folderTable)

		c <- NewSolution(7, 1, p1)
		c <- NewSolution(7, 2, p2)
	},
}

func partOne(table map[string]int) string {
	var sum int
	for _, size := range table {
		if size <= 100000 {
			sum += size
		}
	}
	return fmt.Sprint(sum)
}

func partTwo(table map[string]int) string {
	var diskSize int = 70000000
	var required int = 30000000
	var used int = table["/"]
	var free int = diskSize - used
	var toDelete int = required - free
	var min int = diskSize
	for _, size := range table {
		if size < min && size >= toDelete {
			min = size
		}
	}
	return fmt.Sprint(min)
}

func buildFolderTable(folder *file, table map[string]int) {
	var sum int

	for name, child := range folder.children {
		if name == ".." {
			continue
		}
		if child.isFolder {
			buildFolderTable(child, table)
		}
		sum += child.size
	}

	folder.size = sum
	table[folder.path] = folder.size
}

func newFolder(path string) *file {
	return &file{
		path:     path,
		isFolder: true,
		children: map[string]*file{},
	}
}

func parseLine(line string) logLine {
	log := logLine{}
	log.isCd = patternCD.MatchString(line)
	if log.isCd {
		name := patternCD.FindStringSubmatch(line)[1]
		log.arguments = append(log.arguments, name)
		return log
	}
	log.isFolder = patternFolder.MatchString(line)
	if log.isFolder {
		name := patternFolder.FindStringSubmatch(line)[1]
		log.arguments = append(log.arguments, name)
		return log
	}
	log.isFile = patternFile.MatchString(line)
	if log.isFile {
		matches := patternFile.FindStringSubmatch(line)
		size, name := matches[1], matches[2]
		log.arguments = append(log.arguments, size, name)
	}
	return log
}

type logLine struct {
	isFolder  bool
	isFile    bool
	isCd      bool
	arguments []string
}

type file struct {
	isFolder bool
	path     string
	children map[string]*file
	size     int
}
