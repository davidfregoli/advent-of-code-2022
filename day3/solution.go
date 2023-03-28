package day3

import (
	"strconv"

	"golang.org/x/exp/maps"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var lines []string = reader.ReadLines("day3/input.txt")
		var p1, p2 string = partOne(lines), partTwo(lines)
		return NewSolution(3, 1, p1), NewSolution(3, 2, p2)
	},
}

func partOne(lines []string) string {
	var sum int
	for _, line := range lines {
		left, right := split(line)
		var common rune = findCommon(left, right)
		v := priority(common)
		sum = sum + v
	}
	return strconv.Itoa(sum)
}

func partTwo(lines []string) string {
	var sum int
	for i := 0; i < len(lines); i = i + 3 {
		var common rune = findCommon(lines[i], lines[i+1], lines[i+2])
		v := priority(common)
		sum = sum + v
	}
	return strconv.Itoa(sum)
}

func findCommon(groups ...string) rune {
	set := make(map[rune]bool)
	for i, group := range groups {
		innerSet := make(map[rune]bool)
		for _, letter := range group {
			if i == 0 || set[letter] {
				innerSet[letter] = true
			}
		}
		set = innerSet
	}
	return maps.Keys(set)[0]
}

func priority(i rune) int {
	p := int(i)
	p = p - 96
	if p <= 0 {
		p = p + 58
	}
	return p
}

func split(sack string) (string, string) {
	var length int = len(sack)
	return sack[0 : length/2], sack[length/2 : length]
}
