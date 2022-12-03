// Package main https://adventofcode.com/2022/day/2
// --- Day 2: Rock Paper Scissors ---
package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/davidn5013/aoc/utl"
)

const debugmode = false

//go:embed input.txt
var inputFile string

func main() {
	var (
		input = strings.Split(strings.TrimSpace(inputFile), "\n")
	)

	fmt.Printf("Solution 1 = %d\n", sol1(input))
	fmt.Printf("Solution 2 = %d\n", sol2(input))

}

// Solution part 1
func sol1(lines []string) int {
	var (
		score int = 0
	)
	for _, line := range lines {
		p1, p2 := line[:len(line)/2], line[len(line)/2:]
		utl.Debug(debugmode, "The Two parts : %s %s\n", p1, p2)
		for _, oneChar := range p1 {
			if strings.Contains(p2, string(oneChar)) {
				// a-z 1-26 A-Z 27-52
				if strings.ToUpper(string(oneChar)) == string(oneChar) {
					score += int(oneChar) - 38
					utl.Debug(debugmode, "Dub score=%d %s score %d \n", oneChar-38, string(oneChar), score)
					break
				} else {
					score += int(oneChar) - 96
					utl.Debug(debugmode, "Dub score=%d %s score %d\n", oneChar-96, string(oneChar), score)
					break
				}

			}
		}
	}
	return score
}

// Solution part 2
func sol2(lines []string) int {
	var (
		score       int = 0
		comperline1 string
	)

	for i := 0; i < len(lines); i += 3 {
		comperline1 = lines[i+0] + lines[i+1] + lines[i+2]
		utl.Debug(debugmode, "lines to test \n%s\n%s\n%s\n", lines[i+0], lines[i+1], lines[i+2])
		unikchar := containsInThree(comperline1, lines[i+0], lines[i+1], lines[i+2])
		// a-z 1-26 A-Z 27-52
		if strings.ToUpper(string(unikchar)) == string(unikchar) {
			score += int(unikchar) - 38
			utl.Debug(debugmode, "Dub score=%d %s score %d \n", unikchar-38, string(unikchar), score)
			continue
		} else {
			score += int(unikchar) - 96
			utl.Debug(debugmode, "Dub score=%d %s score %d\n", unikchar-96, string(unikchar), score)
			continue
		}
	}

	return score

}

func containsInThree(comperline1, lines0, lines1, lines2 string) rune {
	for _, r := range comperline1 {
		b1 := strings.Contains(lines0, string(r))
		b2 := strings.Contains(lines1, string(r))
		b3 := strings.Contains(lines2, string(r))
		utl.Debug(debugmode, "Char test %s %v %v %v\n", string(r), b1, b2, b3)
		if b1 && b2 && b3 {
			return r
		}
	}
	return ' '
}
