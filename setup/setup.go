package setup

import (
	"encoding/json"
	"fmt"
	"sort"
)

type AOC struct {
	Year      int
	Problems  []DayProblem
	Solutions []*Solution
	RunOnly   bool
}

func (aoc *AOC) Run() {
	c := make(chan *Solution)
	for _, problem := range aoc.Problems {
		go problem.Solve(c)
	}
	for range aoc.Problems {
		aoc.Solutions = append(aoc.Solutions, <-c, <-c)
	}
	sort.Slice(aoc.Solutions, func(i, j int) bool {
		return aoc.Solutions[i].Day < aoc.Solutions[j].Day ||
			(aoc.Solutions[i].Day == aoc.Solutions[j].Day && aoc.Solutions[i].Part < aoc.Solutions[j].Part)
	})
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
	Solve func(chan *Solution)
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
