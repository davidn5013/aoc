// Package main https://adventofcode.com/2022/day/4
// --- Day 4:
package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const debugmode = false

//go:embed input.txt
var inputFile string

func main() {
	var (
		input = strings.Split(strings.TrimSpace(inputFile), "\n")
	)

	fmt.Printf("Solution 1 = %d\n", sol1(input))
	// fmt.Printf("Solution 2 = %d\n", sol2(input))

}

// Solution part 1
func sol1(lines []string) int {
	var (
		score int = 0
	)
	return score
}

// Solution part 2
func sol2(lines []string) int {
	var (
		score int = 0
	)

	return score
}
