package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var flgPart int
	var flgDebuglvl int

	flag.IntVar(&flgPart, "part", 1, "part 1 or 2")
	flag.IntVar(&flgDebuglvl, "d", 0, "debug information 1-10")
	flag.Parse()

	util.SetDebuglvl(flgDebuglvl)

	fmt.Println("Running part", flgPart)

	if flgPart == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	var sum int
	for _, line := range parsed {
		var first, secound byte
		for i := 0; i < len(line); i++ {
			if isAsciiNum(line[i]) && first == 0 {
				first = line[i]
			}

			j := len(line) - 1 - i
			if isAsciiNum(line[j]) && secound == 0 {
				secound = line[j]
			}

			if first != 00 && secound != 0 {
				break
			}
		}

		n, err := strconv.Atoi(string(first) + string(secound))
		if err != nil {
			log.Fatalln("Failed to convert ", line, err)
		}

		sum += n

	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	var sum int

	for _, line := range parsed {

		var first, secound byte
		for i := 0; i < len(line); i++ {

			s1 := line[:i]
			s2 := line[len(line)-i:]

			v1 := wordToNum(s1)
			v2 := wordToNum(s2)

			j := len(line) - 1 - i

			switch {

			case first == 0 && v1 != 0:
				first = v1

			case first == 0 && isAsciiNum(line[i]):
				first = line[i]

			}

			switch {

			case secound == 0 && v2 != 0:
				secound = v2

			case secound == 0 && isAsciiNum(line[j]):
				secound = line[j]

			}

			if first != 0 && secound != 0 {
				break
			}

		}

		if first == 0 || secound == 0 {
			log.Fatalln("no values reading string ", line)
		}

		n, err := strconv.Atoi(string(first) + string(secound))
		if err != nil {
			log.Fatalln("Failed to convert ", line, err)
		}

		sum += n

	}
	return sum
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func wordToNum(s string) byte {
	words := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for idx, value := range words {
		if strings.Contains(s, idx) {
			return value
		}
	}
	return 0
}

func isAsciiNum(r byte) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
