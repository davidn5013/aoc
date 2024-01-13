// Package d2309 - Solution of Advent of code 2023 day 9 part 1 & part 2
package d2309

import (
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/util"
)

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
		_, a := predict(sensor)
		util.Infof(" %d\n", a)
		totalPredicition += a
	}
	return totalPredicition
}

func predict(input []int) (int, int) {
	// var difs [][]int
	slice1 := input

	hist := input[0]
	pred := input[len(input)-1]
	for {
		dif := sliceDiff(slice1)
		// difs = append(difs, dif)

		util.Debugf("%v ", dif)
		hist -= dif[0]
		pred += dif[len(dif)-1]
		slice1 = dif

		if IsSliceZero(dif) {
			break
		}
	}

	return hist, pred

}

func sliceDiff(input []int) (ans []int) {
	for i := 1; i < len(input); i++ {
		ans = append(ans, input[i]-input[i-1])
	}
	return ans
}

func IsSliceZero(list []int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}

func Part2(input string) int {
	// previewInput(input)
	lines := strings.Split(input, "\n")
	totalPredicition := 0
	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		sensor := make([]int, 0)
		util.Infof("input %v\n", line)

		for _, s := range strings.Split(line, " ") {
			x, _ := strconv.Atoi(string(s))
			sensor = append(sensor, x)
		}

		util.Debugf(" %v ", sensor)
		a, _ := predict(sensor)
		util.Debugf(" %d\n", a)
		totalPredicition += a
	}
	return totalPredicition
}
