package day15

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/davidfregoli/advent-of-code-2022/reader"
	. "github.com/davidfregoli/advent-of-code-2022/setup"
)

const M2 float64 = 2000000

var rePoints = regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)

type segment struct {
	Y     int
	Start int
	End   int
}

var Problem DayProblem = DayProblem{
	Solve: func(c chan *Solution) {
		var scope = 4000000
		var scanned = map[int]bool{}
		var rows = make([][]segment, scope+1)
		var lines = reader.ReadLines("day15/input.txt")
		var beaconsInLine = []int{}

		for _, line := range lines {
			var matches = rePoints.FindAllStringSubmatch(line, -1)
			var sensorX, _ = strconv.ParseFloat(matches[0][1], 64)
			var sensorY, _ = strconv.ParseFloat(matches[0][2], 64)
			var beaconX, _ = strconv.ParseFloat(matches[1][1], 64)
			var beaconY, _ = strconv.ParseFloat(matches[1][2], 64)
			var distance = math.Abs(sensorX-beaconX) + math.Abs(sensorY-beaconY)
			var M2width = distance - math.Abs(M2-sensorY)
			if M2width >= 0 {
				for i := sensorX - M2width; i < sensorX+M2width+1; i++ {
					scanned[int(i)] = true
				}
			}
			if beaconY == M2 {
				beaconsInLine = append(beaconsInLine, int(beaconX))
			}
			for i := sensorY - distance; i <= sensorY+distance; i++ {
				if i >= 0 && i <= float64(scope) {
					var width = distance - math.Abs(i-sensorY)
					var smtNew = segment{
						Y:     int(i),
						Start: int(sensorX - width),
						End:   int(sensorX + width),
					}
					var merged = false
					for j := 0; j < len(rows[int(i)]); j++ {
						var smtOld = &rows[int(i)][j]
						if smtOld.Start <= smtNew.Start && smtOld.End >= smtNew.Start {
							if smtNew.End > smtOld.End {
								smtOld.End = smtNew.End
							}
							merged = true
						} else if smtNew.Start <= smtOld.Start && smtNew.End >= smtOld.Start {
							if smtNew.Start < smtOld.Start {
								smtOld.Start = smtNew.Start
							}
							if smtNew.End > smtOld.End {
								smtOld.End = smtNew.End
							}
							merged = true
						}
					}
					if !merged {
						rows[int(i)] = append(rows[int(i)], smtNew)
					}
				}
			}
		}

		var p1, p2 string = partOne(scanned, beaconsInLine), partTwo(rows)
		c <- NewSolution(15, 1, p1)
		c <- NewSolution(15, 2, p2)
	},
}

func partOne(scanned map[int]bool, beacons []int) string {
	for _, x := range beacons {
		delete(scanned, x)
	}

	var count = len(scanned)

	return fmt.Sprint(count)

}

func partTwo(rows [][]segment) string {
	for _, row := range rows {
		sort.Slice(row, func(i, j int) bool {
			return row[i].Start < row[j].Start || (row[i].Start == row[j].Start && row[i].End < row[j].End)
		})
		var curr = row[0]
		for i := 1; i < len(row); i++ {
			var next = row[i]
			if next.Start > curr.End+1 {
				var frequency = (curr.End+1)*4000000 + curr.Y
				return fmt.Sprint(frequency)
			}
			curr = next
		}
	}
	panic("unreachable")
}
