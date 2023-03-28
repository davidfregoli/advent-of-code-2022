package day2

import (
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var lines []string = reader.ReadLines("day2/input.txt")
		var p1, p2 string = play(lines, scores1), play(lines, scores2)
		return NewSolution(2, 1, p1), NewSolution(2, 2, p2)
	},
}

func play(lines []string, scores map[string]int) string {
	var points int
	for _, line := range lines {
		points += scores[line]
	}
	return strconv.Itoa(points)
}

var scores1 map[string]int = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var scores2 map[string]int = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}
