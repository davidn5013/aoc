package main

import (
	"fmt"

	"github.com/davidn5013/aoc/utl"
)

// ---go:embed input.txt
// var inputFile string

// More debug text 0-10 : 10 Max
var (
	debuglvl = 10
)

func startmsg(funcname string) {
	utl.Debug(debuglvl >= 1, "Starting \t: %s\n", funcname)
}

func (a Aoc) runSolutions(filename string) {
	fmt.Printf("Solution 1 %s \t: %d\n", filename, a.sol1())
	fmt.Printf("Solution 2 %s \t: %d\n", filename, a.sol2())
}
