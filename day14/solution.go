package day14

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var rePoints = regexp.MustCompile(`(\d+),(\d+)`)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var xMax = 0
		var yMax = 0

		var segments = [][]point{}
		var lines = reader.ReadLines("day14/input.txt")

		for _, line := range lines {
			var matches = rePoints.FindAllStringSubmatch(line, -1)
			var segment = []point{}
			for _, match := range matches {
				var x, _ = strconv.Atoi(match[1])
				var y, _ = strconv.Atoi(match[2])
				if x > xMax {
					xMax = x
				}
				if y > yMax {
					yMax = y
				}
				segment = append(segment, point{x: x, y: y})
			}
			segments = append(segments, segment)
		}

		var p1 = partOne(segments, xMax, yMax)
		var p2 = partTwo(segments, xMax, yMax)

		c <- NewSolution(14, 1, p1)
		c <- NewSolution(14, 2, p2)
	},
}

func makeCave(segments [][]point, xMax int, yMax int) [][]bool {
	var cave = make([][]bool, yMax+3)
	for i := 0; i < len(cave); i++ {
		cave[i] = make([]bool, xMax+200)
	}
	for _, segment := range segments {
		var curr = segment[0]
		cave[curr.y][curr.x] = true
		for i := 1; i < len(segment); i++ {
			var next = segment[i]
			var points = interpolate(curr, next)
			for _, point := range points {
				cave[point.y][point.x] = true
			}
			curr = next
			cave[curr.y][curr.x] = true
		}
	}
	return cave
}

func partOne(segments [][]point, xMax int, yMax int) string {
	var count int
	var tick func()
	var active = point{x: -1}
	var cave = makeCave(segments, xMax, yMax)
	tick = func() {
		if active.x == -1 {
			active = point{x: 500, y: 0}
		}
		if active.y+2 > len(cave) {
			return
		} else if cave[active.y+1][active.x] == false {
			active.y++
		} else if cave[active.y+1][active.x-1] == false {
			active.y++
			active.x--
		} else if cave[active.y+1][active.x+1] == false {
			active.y++
			active.x++
		} else {
			cave[active.y][active.x] = true
			active.x = -1
			count++
		}
		tick()
	}
	tick()
	return fmt.Sprint(count)
}

func partTwo(segments [][]point, xMax int, yMax int) string {
	var count int
	var tick func()
	var active = point{x: -1}
	var cave = makeCave(segments, xMax, yMax)
	for i := 0; i < len(cave[0]); i++ {
		cave[yMax+2][i] = true
	}
	tick = func() {
		if active.x == -1 {
			active = point{x: 500, y: 0}
		}
		if cave[0][500] == true {
			return
		} else if cave[active.y+1][active.x] == false {
			active.y++
		} else if cave[active.y+1][active.x-1] == false {
			active.y++
			active.x--
		} else if cave[active.y+1][active.x+1] == false {
			active.y++
			active.x++
		} else {
			cave[active.y][active.x] = true
			active.x = -1
			count++
		}
		tick()
	}
	tick()
	return fmt.Sprint(count)
}

func interpolate(low point, high point) []point {
	var points = []point{}
	if high.x < low.x || high.y < low.y {
		low, high = high, low
	}
	for i := low.x + 1; i < high.x; i++ {
		points = append(points, point{x: i, y: low.y})
	}
	for i := low.y + 1; i < high.y; i++ {
		points = append(points, point{x: low.x, y: i})
	}
	return points
}

type point struct {
	x int
	y int
}
