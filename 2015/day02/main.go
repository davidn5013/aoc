package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// go:embed inp.txt
//go:embed input.txt
var input string

// INFO print all [INFO] lines
const INFO = false

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

	// sumSquarefeet is total some of paper
	var sumSquarefeet int

	// sumWrapfeet
	var sumWrapfeet int

	for linenr, line := range lines {
		cells := strings.Split(line, "x")

		// Infof("[INFO] linenr [%1d]\t%s %s %s\n", linenr, cells[0], cells[1], cells[2])

		l, _ := strconv.Atoi(cells[0])
		w, _ := strconv.Atoi(cells[1])
		h, _ := strconv.Atoi(cells[2])

		// Infof("[INFO] linenr [%1d]\t%d %d %d\n", linenr, l, w, h)

		paper := 2*l*w + 2*w*h + 2*h*l
		minSide := int(math.Min(math.Min(float64(l*w), float64(w*h)), float64(h*l)))

		minSideWrap := int(math.Min(math.Min(float64(l+l+w+w), float64(w+w+h+h)), float64(h+h+l+l)))

		Infof("[INFO] linenr [%1d]\t%d,%d,%d=%d\n", linenr, l, w, h, paper+minSide)

		sumSquarefeet += paper + minSide
		sumWrapfeet += minSideWrap + (l * w * h)
	}

	fmt.Printf("Total feet of paper: %d\n", sumSquarefeet)
	fmt.Printf("Total feet of warp : %d\n", sumWrapfeet)

}

// Infof custom fmt.Printf for INFO lines
func Infof(format string, a ...any) {
	if INFO {
		fmt.Printf(format, a...)
	}
}
