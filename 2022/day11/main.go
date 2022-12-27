// Package main https://adventofcode.com/2022/day/
package main

import (
	_ "embed"

	"github.com/davidn5013/aoc/utl"
)

func main() {
	inputfiles := []string{
		"input_test.txt",
		"input.txt",
	}

	for idx, inputfile := range inputfiles {
		utl.Debug(debuglvl >= 1, "File \t\t: %d %s\n", idx+1, inputfile)
		a := newAoc()
		a.runSolutions(inputfile)
	}
}
