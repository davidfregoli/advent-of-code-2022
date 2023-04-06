package day8

import (
	"fmt"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var lines = reader.ReadLines("day8/input.txt")
		var forest [][]int = [][]int{}

		for _, line := range lines {
			var row = []int{}
			for _, char := range line {
				tree, _ := strconv.Atoi(string(char))
				row = append(row, tree)
			}
			forest = append(forest, row)
		}

		var p1 = partOne(forest)
		var p2 = partTwo(forest)

		c <- NewSolution(8, 1, p1)
		c <- NewSolution(8, 2, p2)
	},
}

func partOne(forest [][]int) string {
	var count int

	for i, row := range forest {
		for j, tree := range row {

			var fromL bool = true
			for t := 0; t < j; t++ {
				if forest[i][t] >= tree {
					fromL = false
					break
				}
			}

			var fromR bool = true
			for t := j + 1; t < len(forest[i]); t++ {
				if forest[i][t] >= tree {
					fromR = false
					break
				}
			}

			var fromT bool = true
			for t := 0; t < i; t++ {
				if forest[t][j] >= tree {
					fromT = false
					break
				}
			}

			var fromB bool = true
			for t := i + 1; t < len(forest); t++ {
				if forest[t][j] >= tree {
					fromB = false
					break
				}
			}

			var visible bool = fromL || fromR || fromT || fromB
			if visible {
				count++
			}

		}
	}

	return fmt.Sprint(count)
}

func partTwo(forest [][]int) string {
	var max int

	for i, row := range forest {
		for j, tree := range row {

			var scoreL int
			for t := j - 1; t >= 0; t-- {
				scoreL++
				if forest[i][t] >= tree {
					break
				}
			}

			var scoreR int
			for t := j + 1; t < len(forest[i]); t++ {
				scoreR++
				if forest[i][t] >= tree {
					break
				}
			}

			var scoreT int
			for t := i - 1; t >= 0; t-- {
				scoreT++
				if forest[t][j] >= tree {
					break
				}
			}

			var scoreB int
			for t := i + 1; t < len(forest); t++ {
				scoreB++
				if forest[t][j] >= tree {
					break
				}
			}

			var score int = scoreL * scoreR * scoreT * scoreB
			if score > max {
				max = score
			}

		}
	}

	return fmt.Sprint(max)
}
