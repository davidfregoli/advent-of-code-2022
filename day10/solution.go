package day10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var reNoop = regexp.MustCompile(`^noop$`)
var reInstruction = regexp.MustCompile(`^addx (-?\d+)$`)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var lines = reader.ReadLines("day10/input.txt")
		var register int = 1
		var history = []int{0}

		for _, line := range lines {
			history = append(history, register)

			isNoop := reNoop.MatchString(line)
			if isNoop {
				continue
			}

			history = append(history, register)

			match := reInstruction.FindStringSubmatch(line)
			val, _ := strconv.Atoi(match[1])
			register += val
		}

		var p1 string = partOne(history)
		var p2 = partTwo(history)

		return NewSolution(10, 1, p1), NewSolution(10, 2, p2)
	},
}

func partOne(history []int) string {
	var sum = 0
	for i := 20; i <= 220; i += 40 {
		sum += i * history[i]
	}
	return fmt.Sprint(sum)
}

func partTwo(history []int) any {
	var screen = make([][]string, 6)
	for cycle := 1; cycle <= 240; cycle++ {
		row := (cycle - 1) / 40
		pixel := ((cycle - 1) % 40) + 1
		if history[cycle] == pixel || history[cycle] == pixel-1 || history[cycle] == pixel-2 {
			screen[row] = append(screen[row], "#")
		} else {
			screen[row] = append(screen[row], ".")
		}
	}
	return map[string]string{
		"row 1": strings.Join(screen[0], ""),
		"row 2": strings.Join(screen[1], ""),
		"row 3": strings.Join(screen[2], ""),
		"row 4": strings.Join(screen[3], ""),
		"row 5": strings.Join(screen[4], ""),
		"row 6": strings.Join(screen[5], ""),
	}
}
