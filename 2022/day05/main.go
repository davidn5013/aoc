// Package main https://adventofcode.com/2022/day/5
package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/davidn5013/aoc/utl"
)

//go:embed input.txt
var inputFile string

// More debug text 0-10 : 10 Max
const debuglvl = 10

// main Run Solution for Advent of Code
func main() {
	var (
		// Split on line feed (Unix file save) after removing white
		input = strings.Split(strings.TrimSpace(inputFile), "\n")
	)

	fmt.Println("Advent of code day 5")

	start := utl.SetTimer()
	fmt.Printf("Solution 1 = %d\n", sol1(input))
	log.Printf("took %s", start())

	start = utl.SetTimer()
	fmt.Printf("Solution 2 = %d\n", sol2(input))
	log.Printf("took %s", start())

}

// sol1 Solution part 1
func sol1(lines []string) (ret int) {

	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	var ()

	return ret
}

// sol2 Solution part 2
func sol2(lines []string) (ret int) {

	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	var ()

	return ret
}
