package day5

import (
	"regexp"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var lines []string = reader.ReadLines("day5/input.txt")
		stacks, commands := parseLines(lines)
		var p1 string = partOne(stacks, commands)
		stacks, commands = parseLines(lines)
		var p2 string = partTwo(stacks, commands)
		return NewSolution(5, 1, p1), NewSolution(5, 2, p2)
	},
}

func partOne(stacks map[int][]byte, commands [][]int) string {
	for _, cmd := range commands {
		qty, from, to := cmd[0], cmd[1], cmd[2]
		for i := 1; i <= qty; i++ {
			last := len(stacks[from]) - 1
			crate := stacks[from][last]
			stacks[from] = stacks[from][0:last]
			stacks[to] = append(stacks[to], crate)
		}
	}
	return getMessage(stacks)
}

func partTwo(stacks map[int][]byte, commands [][]int) string {
	for _, cmd := range commands {
		qty, from, to := cmd[0], cmd[1], cmd[2]
		last := len(stacks[from]) - qty
    crates := stacks[from][last:]
		stacks[from] = stacks[from][0:last]
		stacks[to] = append(stacks[to], crates...)
	}
	return getMessage(stacks)
}

func getMessage(stacks map[int][]byte) string {
	var length int = len(stacks)
	msg := make([]byte, length)
	for k, stack := range stacks {
		last := len(stack) - 1
		msg[k-1] = stack[last]
	}
	return string(msg)
}

func parseLines(lines []string) (map[int][]byte, [][]int) {
	var stacks map[int][]byte = map[int][]byte{}
	var l int

	for lines[l][1] != '1' {
		var line string = lines[l]
		for i, j := 1, 1; j < len(line); i++ {
			var crate byte = line[j]
			j = i*4 + 1
			if crate == ' ' {
				continue
			}
			stacks[i] = append([]byte{crate}, stacks[i]...)
		}
		l++
	}

	pattern := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	var commands [][]int
	for l += 2; l < len(lines); l++ {
		line := lines[l]
		matches := pattern.FindStringSubmatch(line)
		qty, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		commands = append(commands, []int{qty, from, to})
	}

	return stacks, commands
}
