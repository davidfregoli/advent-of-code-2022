package main

import (
	"github.com/davidfregoli/advent-of-code-2022/day1"
	"github.com/davidfregoli/advent-of-code-2022/day10"
	"github.com/davidfregoli/advent-of-code-2022/day11"
	"github.com/davidfregoli/advent-of-code-2022/day12"
	"github.com/davidfregoli/advent-of-code-2022/day13"
	"github.com/davidfregoli/advent-of-code-2022/day14"
	"github.com/davidfregoli/advent-of-code-2022/day15"
	"github.com/davidfregoli/advent-of-code-2022/day2"
	"github.com/davidfregoli/advent-of-code-2022/day3"
	"github.com/davidfregoli/advent-of-code-2022/day4"
	"github.com/davidfregoli/advent-of-code-2022/day5"
	"github.com/davidfregoli/advent-of-code-2022/day6"
	"github.com/davidfregoli/advent-of-code-2022/day7"
	"github.com/davidfregoli/advent-of-code-2022/day8"
	"github.com/davidfregoli/advent-of-code-2022/day9"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

func main() {
	var aoc AOC = AOC{Year: 2023}
	aoc.Solve(day1.Problem)
	aoc.Solve(day2.Problem)
	aoc.Solve(day3.Problem)
	aoc.Solve(day4.Problem)
	aoc.Solve(day5.Problem)
	aoc.Solve(day6.Problem)
	aoc.Solve(day7.Problem)
	aoc.Solve(day8.Problem)
	aoc.Solve(day9.Problem)
	aoc.Solve(day10.Problem)
	aoc.Solve(day11.Problem)
	aoc.Solve(day12.Problem)
	aoc.Solve(day13.Problem)
	aoc.Solve(day14.Problem)
	aoc.Solve(day15.Problem)
	aoc.Run()
	aoc.Print()
}
