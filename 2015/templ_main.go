package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inp.txt
// go:embed input.txt
var input string

// INFO print all [INFO] lines
const INFO = true

func main() {
	// Split on DOS carrier return...
	lines := strings.Split(strings.TrimSpace(input), "\r\n")
	cntLines := len(lines)
	// show sample input file for format check
	if cntLines < 10 {
		Infof("[INFO] Input file:\n%s\n", input)
		for linenr, line := range lines {
			Infof("[INFO] Line %d\t: %s\n", linenr, line)
		}
	}

	// Check input number
	Infof("[INFO] Read number of lines: %d\n", cntLines)

	// - Solution part
	// lines , cntLines
	// sum
	var sum int
	for linenr, line := range lines {
		cells := strings.Split(line, "x")
		// Infof("[INFO] linenr [%1d]\t%s %s %s\n", linenr, cells[0])
		x, _ := strconv.Atoi(cells[0])
		// Infof("[INFO] linenr [%1d]\t%d %d %d\n", linenr, l, w, h)
		// minSide := int(math.Min(math.Min(float64(l*w), float64(w*h)), float64(h*l)))
		// sum +=
	}

	fmt.Printf("Total : %d\n", sum)
}

// Infof custom fmt.Printf for INFO lines
func Infof(format string, a ...any) {
	if INFO {
		fmt.Printf(format, a...)
	}
}
