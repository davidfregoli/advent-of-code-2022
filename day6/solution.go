package day6

import (
	"fmt"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var input = reader.ReadLine("day6/input.txt")
		var b1 buffer = buffer{size: 4}
		var b2 buffer = buffer{size: 14}
		var p1 string = findIndex(input, b1)
		var p2 string = findIndex(input, b2)
		return NewSolution(6, 1, p1), NewSolution(6, 2, p2)
	},
}

type buffer struct {
	count int
	data  string
	size  int
}

func findIndex(input string, b buffer) string {
	var index int
	for _, char := range input {
		b.push(char)
		if b.valid() {
			index = b.count
			break
		}
	}
	return fmt.Sprint(index)
}

func (b *buffer) push(char rune) {
	b.count++
	b.data = b.data + string(char)
	if len(b.data) > b.size {
		b.data = b.data[1:]
	}
}

func (b *buffer) valid() bool {
	if len(b.data) < b.size {
		return false
	}
	set := map[rune]bool{}
	for _, char := range b.data {
		if set[char] {
			return false
		}
		set[char] = true
	}
	return true
}
