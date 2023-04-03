package day9

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var reMotion = regexp.MustCompile(`^([LURD]) (\d+)$`)
var dirMap = map[string][]int{
	"L": {-1, 0},
	"U": {0, 1},
	"R": {1, 0},
	"D": {0, -1},
}

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var motions = []motion{}
		var lines []string = reader.ReadLines("day9/input.txt")
		for _, line := range lines {
			match := reMotion.FindStringSubmatch(line)
			dir := match[1]
			amt, _ := strconv.Atoi(match[2])
			motions = append(motions, motion{dir, amt})
		}

		var p1 string = partOne(motions)
		var p2 string = partTwo(motions)

		return NewSolution(9, 1, p1), NewSolution(9, 2, p2)
	},
}

func partOne(motions []motion) string {
	var H = knot{x: 0, y: 0}
	var T = knot{x: 0, y: 0}

	var visited map[string]bool = map[string]bool{}
	visit := func(x, y int) {
		key := fmt.Sprint(x, ",", y)
		visited[key] = true
	}
	visit(T.x, T.y)

	for _, motion := range motions {
		for i := 0; i < motion.amt; i++ {
			H.move(motion.dir)
			T.follow(H)
			visit(T.x, T.y)
		}
	}

	return fmt.Sprint(len(visited))
}

func partTwo(motions []motion) string {
	var knots = make([]knot, 10)
	for i := range knots {
		knots[i] = knot{0, 0}
	}

	var visited map[string]bool = map[string]bool{}
	visit := func(x, y int) {
		key := fmt.Sprint(x, ",", y)
		visited[key] = true
	}
	visit(0, 0)

	for _, motion := range motions {
		for i := 0; i < motion.amt; i++ {
			knots[0].move(motion.dir)
			prev := knots[0]
			for i := 1; i < 10; i++ {
				curr := &knots[i]
				curr.follow(prev)
				prev = *curr
			}
			visit(knots[9].x, knots[9].y)
		}
	}

	return fmt.Sprint(len(visited))
}

type knot struct {
	x int
	y int
}

func (k *knot) move(dir string) {
	motion := dirMap[dir]
	x := motion[0]
	y := motion[1]
	k.x += x
	k.y += y
}

func (t *knot) follow(h knot) {
	dx := math.Abs(float64(h.x - t.x))
	dy := math.Abs(float64(h.y - t.y))

	if dx < 2 && dy < 2 {
		return
	}

	if dy > 1 {
		if h.y > t.y {
			t.y++
		} else {
			t.y--
		}
		if dx >= 1 {
			if h.x > t.x {
				t.x++
			} else {
				t.x--
			}
		}
		return
	}

	if dx > 1 {
		if h.x > t.x {
			t.x++
		} else {
			t.x--
		}
		if dy >= 1 {
			if h.y > t.y {
				t.y++
			} else {
				t.y--
			}
		}
	}
}

type motion struct {
	dir string
	amt int
}
