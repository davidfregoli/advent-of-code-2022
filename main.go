package main

import (
	"github.com/davidfregoli/advent-of-code-2022/day1"
	"github.com/davidfregoli/advent-of-code-2022/day2"
	"github.com/davidfregoli/advent-of-code-2022/day3"
	"github.com/davidfregoli/advent-of-code-2022/day4"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

func main() {
	var aoc AOC = AOC{Year: 2023}
	aoc.Solve(day1.Problem)
	aoc.Solve(day2.Problem)
	aoc.Solve(day3.Problem)
	aoc.Solve(day4.Problem)
	aoc.Run()
	aoc.Print()
}

