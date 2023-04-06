package day4

import (
	"strconv"
	"strings"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var lines []string = reader.ReadLines("day4/input.txt")
		var pairs [][][]int = parseLines(lines)
		var p1 string = partOne(pairs)
		var p2 string = partTwo(pairs)
		c <- NewSolution(4, 1, p1)
		c <- NewSolution(4, 2, p2)
	},
}

func partOne(pairs [][][]int) string {
	var count int
	for _, pair := range pairs {
		var fromA, toA int = pair[0][0], pair[0][1]
		var fromB, toB int = pair[1][0], pair[1][1]
		var aContainsB bool = fromB >= fromA && toB <= toA
		var bContainsA bool = fromA >= fromB && toA <= toB
		if aContainsB || bContainsA {
			count = count + 1
		}
	}
	return strconv.Itoa(count)
}

func partTwo(pairs [][][]int) string {
	var count int
	for _, pair := range pairs {
		var fromA, toA int = pair[0][0], pair[0][1]
		var fromB, toB int = pair[1][0], pair[1][1]
		var aContainsB bool = fromB >= fromA && fromB <= toA
		var bContainsA bool = fromA >= fromB && fromA <= toB
		if aContainsB || bContainsA {
			count = count + 1
		}
	}
	return strconv.Itoa(count)
}

func parseLines(lines []string) [][][]int {
	var pairs [][][]int
	for _, line := range lines {
		var ranges []string = strings.Split(line, ",")
		var rangeA []string = strings.Split(ranges[0], "-")
		var rangeB []string = strings.Split(ranges[1], "-")
		var fromA, _ = strconv.Atoi(rangeA[0])
		var toA, _ = strconv.Atoi(rangeA[1])
		var fromB, _ = strconv.Atoi(rangeB[0])
		var toB, _ = strconv.Atoi(rangeB[1])
		var pair [][]int = [][]int{{fromA, toA}, {fromB, toB}}
		pairs = append(pairs, pair)
	}
	return pairs
}
