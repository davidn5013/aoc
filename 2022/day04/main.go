// Package main https://adventofcode.com/2022/day/4
package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/davidn5013/aoc/utl"
)

// debug text 0-10 : 10 Max
const debuglvl = 0

//go:embed input.txt
var inputFile string

// main Run Solution for Advent of Code
func main() {
	// Split on line feed (Unix file save) after removing white
	input := strings.Split(strings.TrimSpace(inputFile), "\n")

	start := SetTimer()
	fmt.Printf("Solution 1 = (528) %d\n", mostSol(input, 1))
	log.Printf("took %s", start())
	start = SetTimer()
	fmt.Printf("Solution 2 = (881) %d\n", mostSol(input, 2))
	log.Printf("took %s", start())
}

// mostSol , solution part 1 & 2
func mostSol(lines []string, v int) (ret int) {
	utl.Debug(debuglvl >= 10, "Input for %s\n %v\n", utl.CurrFuncName(), lines)

	// extra pars of "2-4,6-8" to [0-3]int
	for idx, line := range lines {
		utl.Debug(debuglvl >= 10, "Line %d = %#v\n", idx, line)

		section := mostGetSection(line, idx)

		if v == 1 {
			interSec := sectOverLap(section[0:2], section[2:4])
			if len(interSec) > 0 {
				ret++
				utl.Debug(debuglvl >= 2, "Intersection array in line %d = %#v %d\n", idx+1, interSec, len(interSec)/2)
			}
		} else if v == 2 {
			interSec := sectIncludes(section[0:2], section[2:4])
			if len(interSec) > 0 {
				ret++
				utl.Debug(debuglvl >= 2, "Intersection array in line %d = %#v %d\n", idx+1, interSec, len(interSec))
			}

		}

	}

	return ret
}

// secOverLap is what solv part 1
//
//	a0------------a1
//	     b0-----------b1  or vice versa
func sectOverLap(a, b []int) (res []int) {
	if (a[0] <= b[0] && a[1] >= b[1]) || (b[0] <= a[0] && b[1] >= a[1]) {
		res = append(res, a[0], a[1], b[0], b[1])
	}
	return res
}

// secIntersect is what solv part 2
//
//	  a0 ----------------- a1
//		 b0 ------------ b1 or vice versa
func sectIncludes(a, b []int) (res []int) {
	if !(a[1] < b[0] || a[0] > b[1]) {
		res = append(res, a[0], a[1], b[0], b[1])
	}
	return res
}

// mostGetSection 11-73,29-73 -> []int{11,73,29,73}
// Just a line parser nothing see here
func mostGetSection(line string, linenr int) (section []int) {
	for _, blocks := range strings.Split(line, ",") {
		for _, cell := range strings.Split(blocks, "-") {
			i, err := strconv.Atoi(string(cell))
			utl.PanicIf(err != nil, "ERR: %s wrong input in line %d\n", utl.CurrFuncName(), linenr)
			section = append(section, i)
		}
	}
	return section
}

// SetTimer Set a timmer and return a func
// that returns time.Duration from the timer
// func main() {
// stopTimer := SetTimer()
// ...
// fmt.Printf("Elapsed time:%v\n".stopTimer())
func SetTimer() func() time.Duration {
	t := time.Now()
	return func() time.Duration {
		return time.Since(t)
	}
}
