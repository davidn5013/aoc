package main

import (
	"fmt"

	"github.com/davidn5013/aoc/utl"
)

// ---go:embed input.txt
// var inputFile string

// More debug text 0-10 : 10 Max
var (
	debuglvl = 4
)

func startmsg(funcname string) {
	utl.Debug(debuglvl >= 1, "Starting \t: %s\n", funcname)
}

func (a aoc) runSolutions(filename string) {
	fmt.Printf("File\t\t: %s\nSolution 1\t: %d \n", filename, a.sol1(filename))
	fmt.Printf("File\t\t: %s\nSolution 2\t: %d \n", filename, a.sol2(filename))
	fmt.Println()
}
