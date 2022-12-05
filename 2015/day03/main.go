package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// The sample input is not ment to inpu all 3 line
// but line by line and that give the correct value
// go:embed inp.txt
//go:embed input.txt
var input string

// INFO print all [INFO] lines
const INFO = true

func main() {
	// Split on DOS carrier return...
	lines := strings.Split(strings.TrimSpace(input), "\r\n")
	cntLines := len(lines)
	// show sample input file for format check
	if len(input) < 100 {
		Infof("[INFO] Mini input file:\n%s\n", input)
		for linenr, line := range lines {
			Infof("[INFO] Line %d\t: %s\n", linenr, line)
		}
	}

	// Check input number
	Infof("[INFO] Read number of lines: %d\n", cntLines)

	// Corr x,y cordinations in map
	type Corr struct {
		x int
		y int
	}
	corr := Corr{0, 0}
	path := make(map[Corr]int)
	// path2 := make(map[Corr]int)
	path[corr] = 1
	// path[corr] = 2

	for linenr, line := range lines {
		for pos, char := range line {
			Infof("[INFO] Line %0d\t Pos %0d\t Char %c\n", linenr, pos, char)
			correct := true
			switch char {
			case '^':
				corr.y--
			case '>':
				corr.x++
			case 'v':
				corr.y++
			case '<':
				corr.x--
			default:
				correct = false
			}
			if correct {
				path[corr] = path[corr] + 1
			}
		}
	}
	fmt.Printf("Solution 1 length of path %d:", len(path))
}

// Infof custom fmt.Printf for INFO lines
func Infof(format string, a ...any) {
	if INFO {
		fmt.Printf(format, a...)
	}
}
