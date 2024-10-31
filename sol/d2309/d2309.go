// Package d2309 - Solution of Advent of code 2023 day 9 part 1 & part 2
package d2309

import (
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/util"
)

func sensorPrediction(input []int) (int, int) {
	temp := input
	prev := input[0]
	next := input[len(input)-1]

	for {
		dif := diff(temp)
		// difs = append(difs, dif)

		prev -= dif[0]
		next += dif[len(dif)-1]
		temp = dif

		util.Debugf("%d%v%d ", prev, dif, next)

		if isSliceZero(dif) {
			break
		}
	}

	return prev, next
}

func diff(input []int) (ans []int) {
	ans = make([]int, 0)

	for i := 1; i < len(input); i++ {
		ans = append(ans, input[i]-input[i-1])
	}

	return ans
}

func isSliceZero(list []int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}

	return true
}

// Part1 Mirage Maintenance
func Part1(input string) int {
	// previewInput(input)
	lines := strings.Split(input, "\n")
	totalPredicition := 0
	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		sensor := make([]int, 0)
		util.Infof("input %v\t", line)

		for _, s := range strings.Split(line, " ") {
			x, _ := strconv.Atoi(string(s))
			sensor = append(sensor, x)
		}

		util.Infof(" %v ", sensor)
		_, a := sensorPrediction(sensor)
		util.Infof(" %d\n", a)
		totalPredicition += a
	}

	return totalPredicition
}

// Part2 Mirage Maintenance
func Part2(input string) int {
	// BUG
	// input 1 3 6 10 15 21
	// -1[2 3 4 5 6]27 -2[1 1 1 1]28 -2[0 0 0]28  [1 3 6 10 15 21]  -2
	// Result -2 but correct is 0

	lines := strings.Split(input, "\n")
	totPred := 0

	for _, line := range lines {

		util.Infof("input %v\n", line)

		if len(line) == 0 {
			continue
		}

		sensor := make([]int, 0)

		for _, s := range strings.Split(line, " ") {
			x, _ := strconv.Atoi(string(s))
			sensor = append(sensor, x)
		}

		a, _ := sensorPrediction(sensor)
		totPred += a

		util.Debugf(" %v ", sensor)
		util.Debugf(" %d\n", a)

	}

	return totPred
}
