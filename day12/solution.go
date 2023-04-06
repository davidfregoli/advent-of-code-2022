package day12

import (
	"fmt"
	"sort"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var (
			start []int
			end   []int
			cells = [][]int{}
			lines = reader.ReadLines("day12/input.txt")
		)

		for y, line := range lines {
			var row = []int{}
			for x, cell := range line {
				if cell == 'S' {
					start = []int{y, x}
					row = append(row, 1)
					continue
				}
				if cell == 'E' {
					end = []int{y, x}
					row = append(row, 27)
					continue
				}
				row = append(row, int(cell)-96)
			}
			cells = append(cells, row)
		}

		var max = len(cells) * len(cells[0])
		var history = [][]int{}
		for _, row := range cells {
			hRow := make([]int, len(row))
			for i := range hRow {
				hRow[i] = max
			}
			history = append(history, hRow)
		}
		var p1 = fmt.Sprint(explore(start[0], start[1], cells, 0, history))

		for i, row := range cells {
			for j := range row {
				history[i][j] = max
			}
		}
		var p2 = fmt.Sprint(exploreDown(end[0], end[1], cells, 0, history))

		c <- NewSolution(12, 1, p1)
		c <- NewSolution(12, 2, p2)
	},
}

func explore(y int, x int, cells [][]int, steps int, history [][]int) int {
	history[y][x] = steps
	if cells[y][x] == 27 {
		return steps
	}
	steps++
	var height = len(cells)
	var width = len(cells[0])
	var directions = []int{-1, -1, -1, -1}
	if x > 0 {
		if cells[y][x-1] <= cells[y][x]+1 && history[y][x-1] > steps {
			directions[0] = explore(y, x-1, cells, steps, history)
		}
	}
	if y > 0 {
		if cells[y-1][x] <= cells[y][x]+1 && history[y-1][x] > steps {
			directions[1] = explore(y-1, x, cells, steps, history)
		}
	}
	if x+1 < width {
		if cells[y][x+1] <= cells[y][x]+1 && history[y][x+1] > steps {
			directions[2] = explore(y, x+1, cells, steps, history)
		}
	}
	if y+1 < height {
		if cells[y+1][x] <= cells[y][x]+1 && history[y+1][x] > steps {
			directions[3] = explore(y+1, x, cells, steps, history)
		}
	}
	sort.Slice(directions, func(i, j int) bool {
		if directions[i] == -1 {
			return false
		}
		if directions[j] == -1 {
			return true
		}
		return directions[i] < directions[j]
	})
	return directions[0]
}

func exploreDown(y int, x int, cells [][]int, steps int, history [][]int) int {
	history[y][x] = steps
	if cells[y][x] == 1 {
		return steps
	}
	steps++
	var height = len(cells)
	var width = len(cells[0])
	var directions = []int{-1, -1, -1, -1}
	if x > 0 {
		if cells[y][x-1] >= cells[y][x]-1 && history[y][x-1] > steps {
			directions[0] = exploreDown(y, x-1, cells, steps, history)
		}
	}
	if y > 0 {
		if cells[y-1][x] >= cells[y][x]-1 && history[y-1][x] > steps {
			directions[1] = exploreDown(y-1, x, cells, steps, history)
		}
	}
	if x+1 < width {
		if cells[y][x+1] >= cells[y][x]-1 && history[y][x+1] > steps {
			directions[2] = exploreDown(y, x+1, cells, steps, history)
		}
	}
	if y+1 < height {
		if cells[y+1][x] >= cells[y][x]-1 && history[y+1][x] > steps {
			directions[3] = exploreDown(y+1, x, cells, steps, history)
		}
	}
	sort.Slice(directions, func(i, j int) bool {
		if directions[i] == -1 {
			return false
		}
		if directions[j] == -1 {
			return true
		}
		return directions[i] < directions[j]
	})
	return directions[0]
}
