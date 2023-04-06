package day1

import (
	"sort"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var lines []string = reader.ReadLines("day1/input.txt")

		var elfs [][]int = groupByElf(lines)
		var sums []int = sumByElf(elfs)
		var p1, p2 string = partOne(sums), partTwo(sums)
		c <- NewSolution(1, 1, p1)
		c <- NewSolution(1, 2, p2)
	},
}

func partOne(sums []int) string {
	return strconv.Itoa(sums[0])
}

func partTwo(sums []int) string {
	var sum int
	for i := 0; i < 3; i++ {
		sum += sums[i]
	}
	return strconv.Itoa(sum)
}

func groupByElf(lines []string) [][]int {
	var elfs [][]int
	var elf []int
	for _, line := range lines {

		if line == "" {
			elfs = append(elfs, elf)
			elf = []int{}
			continue
		}

		n, _ := strconv.Atoi(line)
		elf = append(elf, n)
	}

	if len(elf) > 0 {
		elfs = append(elfs, elf)
	}

	return elfs
}

func sumByElf(elfs [][]int) []int {
	var sums []int

	for _, elf := range elfs {
		var sum int
		for _, cals := range elf {
			sum = sum + cals
		}
		sums = append(sums, sum)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums
}
