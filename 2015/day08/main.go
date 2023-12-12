package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var realInput string

//go:embed test.txt
var testInput string

func main() {

	var part int
	var selectFile bool

	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&selectFile, "td", false, "Test Data - Use test input instead for real")
	flag.Parse()

	var input *string
	if !selectFile {
		input = &realInput
	} else {
		input = &testInput
	}

	switch part {
	case 1:
		ans := part1(*input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output part 1:", ans)
	case 2:
		ans := part2(*input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output part 2:", ans)
	default:
		fmt.Println("Missing part")
		os.Exit(-1)
	}
}

func part1(input string) int {

	parsed := parseInput(input)
	lenTot := 0
	escapedTot := 0

	for _, line := range parsed {

		if len(line) <= 0 {
			continue
		}

		// fmt.Println(line, len(line), escapedString(line))

		lenTot += len(line)
		escapedTot += escapedString(line)

	}

	// fmt.Println(escapedTot, lenTot)
	return lenTot - escapedTot

}

func escapedString(line string) (escaped int) {

	for i := 1; i < len(line)-1; i++ {
		escaped++

		if line[i] == '\\' {
			if line[i+1] == 'x' {
				i += 3
			} else {
				i += 1
			}
		}

	}

	return escaped

}

func part2(input string) int {

	parsed := parseInput(input)

	total := 0
	encoded := 0

	nr := 0

	// fmt.Printf("%s\n", parsed)
	for _, v := range parsed {
		if len(v) <= 0 {
			break
		}

		t, e := encodeString(v)

		total += t
		encoded += e
		nr++

		// fmt.Println("  ", t, e)

	}

	// fmt.Println("Totals:", encoded, total)
	return encoded - total

}

func encodeString(line string) (total, encode int) {

	// outer quotes
	encodeString := "\""
	for i := 0; i < len(line); i++ {

		switch line[i] {
		// " = \"
		// \ = \\
		// \xYY= \\xYY (cover by \)
		case '\\':
			encodeString += "\\\\"
		case '"':
			encodeString += "\\\""
		default:
			encodeString += string(line[i])
		}
	}
	encodeString += "\""
	// fmt.Print(" > encoded string: ", encodeString)
	return len(line), len(encodeString)

}

func parseInput(input string) (ans []string) {

	return strings.Split(input, "\n")

}
