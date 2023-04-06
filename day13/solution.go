package day13

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

var reTokens = regexp.MustCompile(`(\d+|[\[\],])`)

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var pairs = []pair{}
		var list = []any{}
		var lines = reader.ReadLines("day13/input.txt")

		for i := 0; i < len(lines); i += 3 {
			var left = parseTokens(lines[i])
			var right = parseTokens(lines[i+1])
			var pair = pair{
				Left:  left,
				Right: right,
			}
			pairs = append(pairs, pair)
			list = append(list, left, right)
		}

		var p1 = partOne(pairs)
		var p2 = partTwo(list)
		c <- NewSolution(13, 1, p1)
		c <- NewSolution(13, 2, p2)
	},
}

func partOne(pairs []pair) string {
	var sum int
	for i, pair := range pairs {
		var index = i + 1
		if compare(pair.Left, pair.Right) < 0 {
			sum += index
		}
	}
	return fmt.Sprint(sum)
}

func partTwo(list []any) string {
	var prod = 1
	var divider1 = []any{[]any{2}}
	var divider2 = []any{[]any{6}}
	list = append(list, divider1, divider2)

	sort.Slice(list, func(i, j int) bool {
		return compare(list[i].([]any), list[j].([]any)) < 0
	})

	for i, el := range list {
		if compare(el.([]any), divider1) == 0 || compare(el.([]any), divider2) == 0 {
			prod *= (i + 1)
		}
	}

	return fmt.Sprint(prod)
}

func compare(left []any, right []any) int {
	for i := range left {
		if len(right) < i+1 {
			return 1
		}
		var kindL = reflect.ValueOf(left[i]).Kind().String()
		var kindR = reflect.ValueOf(right[i]).Kind().String()
		if kindL == "int" && kindR == "int" {
			var diff = left[i].(int) - right[i].(int)
			if diff == 0 {
				continue
			}
			return diff
		}
		var diff int
		if kindL == "int" {
			diff = compare([]any{left[i]}, right[i].([]any))
		} else if kindR == "int" {
			diff = compare(left[i].([]any), []any{right[i]})
		} else {
			diff = compare(left[i].([]any), right[i].([]any))
		}
		if diff == 0 {
			continue
		}
		return diff
	}
	if len(right) > len(left) {
		return -1
	}
	return 0
}

func parseTokens(line string) []any {
	var matches = reTokens.FindAllStringSubmatch(line, -1)
	var tokens = []token{}
	for _, match := range matches {
		var val = match[1]
		if val == "," {
			continue
		}
		var token = token{}
		if val == "[" {
			token.Operation = "OPEN"
		} else if val == "]" {
			token.Operation = "CLOSE"
		} else {
			token.Operation = "NUMBER"
			token.Value, _ = strconv.Atoi(val)
		}
		tokens = append(tokens, token)
	}
	var structure, _ = build(tokens[1:])
	return structure
}

func build(tokens []token) ([]any, int) {
	var list = []any{}
	for i := 0; i < len(tokens); i++ {
		var token = tokens[i]
		if token.Operation == "OPEN" {
			var child, skip = build(tokens[i+1:])
			list = append(list, child)
			i += skip
		}
		if token.Operation == "CLOSE" {
			return list, i + 1
		}
		if token.Operation == "NUMBER" {
			list = append(list, token.Value)
		}
	}
	panic("unreachable")
}

type pair struct {
	Left  []any
	Right []any
}

type token struct {
	Operation string
	Value     int
}
