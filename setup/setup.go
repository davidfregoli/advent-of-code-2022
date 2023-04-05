package setup

import (
	"encoding/json"
	"fmt"
)

type AOC struct {
	Year      int
	Problems  []DayProblem
	Solutions []*Solution
	RunOnly   bool
}

func (aoc *AOC) Run() {
	for _, problem := range aoc.Problems {
		solution1, solution2 := problem.Solve()
		aoc.Solutions = append(aoc.Solutions, solution1, solution2)
	}
}

func (aoc *AOC) Solve(p DayProblem) {
	if !aoc.RunOnly {
		aoc.Problems = append(aoc.Problems, p)
	}
}

func (aoc *AOC) Only(p DayProblem) {
	if !aoc.RunOnly {
		aoc.RunOnly = true
		aoc.Problems = []DayProblem{}
	}
	aoc.Problems = append(aoc.Problems, p)
}

func (aoc *AOC) Print() {
	data, _ := json.MarshalIndent(aoc.Solutions, "", "  ")
	fmt.Printf("%v", string(data))
}

type DayProblem struct {
	Solve func() (*Solution, *Solution)
}

type Solution struct {
	Day   int
	Part  int
	Value any
}

func NewSolution(day int, part int, value any) *Solution {
	return &Solution{
		Day:   day,
		Part:  part,
		Value: value,
	}
}

func Debug(d any) {
	data, _ := json.MarshalIndent(d, "", "  ")
	fmt.Printf("%v", string(data))
}
