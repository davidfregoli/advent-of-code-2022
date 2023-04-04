package day11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var reNumber = regexp.MustCompile(`\d+`)
var reOperation = regexp.MustCompile(`Operation: new = old ([+*]) (\d+|old)`)

var Problem DayProblem = DayProblem{
	Solve: func() (*Solution, *Solution) {
		var lines = reader.ReadLines("day11/input.txt")
		var p1 string = partOne(parseLines(lines))
		var p2 string = partTwo(parseLines(lines))
		return NewSolution(11, 1, p1), NewSolution(11, 2, p2)
	},
}

func partOne(monkeys []monkey) string {
	var activity = make([]int, len(monkeys))
	for round := 1; round <= 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for _, item := range monkey.Items {
				activity[i]++
				worry := monkey.Operation(item)
				worry = worry / 3
				if worry%monkey.Divisor == 0 {
					monkeys[monkey.Pass].Items = append(monkeys[monkey.Pass].Items, worry)
				} else {
					monkeys[monkey.Fail].Items = append(monkeys[monkey.Fail].Items, worry)
				}
			}
			monkey.Items = []int{}
		}
	}

	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})
	var product = activity[0] * activity[1]

	return fmt.Sprint(product)
}

func partTwo(monkeys []monkey) string {
	var lcm = 1
	for _, monkey := range monkeys {
		lcm *= monkey.Divisor
	}

	var activity = make([]int, len(monkeys))
	for round := 1; round <= 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for _, item := range monkey.Items {
				activity[i]++
				worry := monkey.Operation(item)
				worry = worry % lcm
				if worry%monkey.Divisor == 0 {
					monkeys[monkey.Pass].Items = append(monkeys[monkey.Pass].Items, worry)
				} else {
					monkeys[monkey.Fail].Items = append(monkeys[monkey.Fail].Items, worry)
				}
			}
			monkey.Items = []int{}
		}
	}

	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})
	var product = activity[0] * activity[1]

	return fmt.Sprint(product)
}

func parseLines(lines []string) []monkey {
	var monkeys = []monkey{}
	for i := 0; i < len(lines); {
		var mky = monkey{}

		i++
		var siMatch = reNumber.FindAllStringSubmatch(lines[i], -1)
		for _, match := range siMatch {
			var item, _ = strconv.Atoi(match[0])
			mky.Items = append(mky.Items, item)
		}

		i++
		var opMatch = reOperation.FindStringSubmatch(lines[i])
		var operator = opMatch[1]

		if opMatch[2] == "old" {
			if operator == "+" {
				mky.Operation = func(old int) int {
					return old + old
				}
			} else if operator == "*" {
				mky.Operation = func(old int) int {
					return old * old
				}
			}
		} else {
			var amt, _ = strconv.Atoi(opMatch[2])
			if operator == "+" {
				mky.Operation = func(old int) int {
					return old + amt
				}
			} else if operator == "*" {
				mky.Operation = func(old int) int {
					return old * amt
				}
			}
		}

		i++
		var divMatch = reNumber.FindStringSubmatch(lines[i])
		var divisor, _ = strconv.Atoi(divMatch[0])
		mky.Divisor = divisor

		i++
		var passMatch = reNumber.FindStringSubmatch(lines[i])
		var pass, _ = strconv.Atoi(passMatch[0])
		mky.Pass = pass

		i++
		var divFail = reNumber.FindStringSubmatch(lines[i])
		var fail, _ = strconv.Atoi(divFail[0])
		mky.Fail = fail

		monkeys = append(monkeys, mky)
		i++
		i++
	}
	return monkeys
}

type monkey struct {
	Items     []int
	Operation func(int) int
	Divisor   int
	Pass      int
	Fail      int
}
