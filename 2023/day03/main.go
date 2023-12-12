package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

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

	var flgPart, flgDebuglvl int

	flag.IntVar(&flgPart, "part", 1, "part 1 or 2")
	flag.IntVar(&flgDebuglvl, "d", 0, "Set debug text level 1-10")
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

	util.SetDebuglvl(0)

	parsed := parseInput(input)
	sum := 0

	util.Debugf("lenght of parsed %d", len(parsed))
	for idx, line := range parsed {
		// util.Debugf("%d , %#v\n", idx, string(line))
		num := parseLine(line)

		for _, num := range num {
			util.Debugf("Check number %d %d %d\n", num.nr, num.start, num.end)
			if hasSymbol(parsed, idx, num) {
				sum += num.nr
			}
		}
	}
	return sum

}

func hasSymbol(parsed [][]byte, row int, n number) bool {

	symbol := ".0123456789"

	checkstring := ""
	for r := -1; r <= 1; r++ {

		for c := n.start - 1; c <= n.end+1; c++ {

			// out of bounce
			maxcols := len(parsed[row]) - 1
			maxrows := len(parsed) - 1
			testrow := row - r

			if c < 0 || c > maxcols ||
				testrow < 0 || testrow > maxrows {
				continue
			}

			if testrow == row &&
				c >= n.start && c < n.end {
				continue
			}

			check := parsed[testrow][c]
			checkstring += string(check)

			if !strings.Contains(symbol, string(check)) {
				util.Debugf("hasSymbol found symbol in string %s\n", checkstring)
				return true
			}
		}

	}

	util.Debugf("hasSymbol check string %s\n", checkstring)
	return false

}

type number struct {
	nr    int
	start int
	end   int
}

// parseLine
// take line ..667..878 return as int and start position and end position
func parseLine(line []byte) (nr []number) {
	s := ""
	for idx, b := range line {
		isDigit := unicode.IsDigit(rune(b))
		if isDigit {
			s += string(b)
		}

		lastbyte := idx == len(line)-1
		if (!isDigit || lastbyte) && s != "" {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("ERR Failed to convert %s %d, %s", s, n, err)
			}
			var temp number
			temp.nr = n
			temp.start = idx - len(s)
			temp.end = idx
			nr = append(nr, temp)
			s = ""
		}
	}
	return nr
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (parsed [][]byte) {
	for _, line := range strings.Split(input, "\n") {
		temp := []byte{}
		for _, b := range line {
			temp = append(temp, byte(b))
		}
		parsed = append(parsed, temp)
	}
	return parsed
}

func _parseInput(input string) (ans []string) {
	lines := strings.Split(input, "\n")

	for row, line := range lines {
		var numstr string
		var foundnumber bool

		for col := 0; col < len(line); col++ {
			r := line[col]
		nextcol:
			if !isNumber(string(r)) || col >= len(line)-1 {

				if !foundnumber {
					continue
				}

				// adjacent check for symbols from ever digit in number
				for i := len(numstr); i > 0; i-- {
					l := lines[row]
					c := l[col-i]
					util.Debugf(" >> adj test on %s ", string(c))
					adj := adjacent(lines, row, col-i)
					for _, a := range adj {
						if !isSymb(string(a)) {
							continue
						}

						util.Infof(" number:%s has symbol\n", numstr)
						ans = append(ans, numstr)

						numstr = ""
						foundnumber = false
						goto nextcol

					}
				}

				numstr = ""
				foundnumber = false

			} else {
				numstr += string(r)
				foundnumber = true
			}

		}

	}

	return ans
}

func adjacent(lines []string, row, col int) (adj string) {
	positions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},

		{0, 1},

		{1, 1},
		{0, -1},
		{1, -1},

		{1, 0},
	}

	// Presumption that all lines are the same length...
	maxc := len(lines[0])
	maxr := len(lines)

	for _, pos := range positions {
		r := row + pos[0]
		c := col + pos[1]

		// skip adjacent check of out of bounce
		if r < 0 || c < 0 || r >= maxr || c >= maxc {
			continue
		}

		l := lines[r]
		adj += string(l[c])
	}

	l := lines[row]
	util.Debugf(" > adjcent %s for row %d col %d are %s\n", string(l[col]), row, col, adj)

	return adj

}

func isNumber(s string) bool {
	numbers := "0123456789"
	if strings.ContainsAny(numbers, s) {
		return true
	}
	return false
}

func isSymb(s string) bool {
	notContains := ".0123456789"
	if !strings.ContainsAny(notContains, s) {
		return true
	}
	return false
}
