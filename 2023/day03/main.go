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
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	util.SetDebuglvl(10)

	if part == 1 {
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
	total := 0
	numstr := ""

	for line := 0; line < len(parsed); line++ {

		for col := 0; col < len(parsed[line]); col++ {
			c := parsed[line][col]

			if c >= '0' && c <= '9' {
				numstr += string(c)
			} else {
				if numstr != "" {
					mutAdj(parsed, &numstr, line, col, &total)
				}
			}

		}

		if numstr != "" {
			mutAdj(parsed, &numstr, line, len(parsed[line])-1, &total)
		}

	}

	return total
}

func mutAdj(input []string, numstr *string, line, col int, total *int) {
	if adjacent(input, *numstr, line, col) {

		n, err := strconv.Atoi(*numstr)
		if err != nil {
			log.Fatalf("ERR converting string %s to number", *numstr)
		}

		// fmt.Printf("%s ", *numstr)
		*total += n

	}

	*numstr = ""

}

func adjacent(input []string, numstr string, line, col int) bool {
	for adjline := line - 1; adjline <= line+1; adjline++ {
		if adjline < 0 || adjline >= len(input) {
			continue
		}

		for adjcol := col; adjcol >= col-len(numstr)-1; adjcol-- {

			if adjcol < 0 || adjcol >= len(input[adjline]) {
				continue
			}

			c := input[adjline][adjcol]

			isNum := c >= '0' && c <= '9'
			if !isNum && c != '.' {
				return true
			}

		}
	}
	return false
}

/*

För att match spara numret, kol,rad,längd där det unik är kol,rad

*/

type GearPos struct {
	row      int
	colStart int
	colEnd   int
}
type GearValue int
type Gears map[GearPos]GearValue

func (g Gears) setGear(row, colStart, colEnd, value int) {
	g[GearPos{
		row:      row,
		colStart: colStart,
		colEnd:   colEnd,
	}] = GearValue(value)
}

func (g Gears) sumgears(input []string) (sum int) {
	used := make(map[GearPos]GearValue)

	for line := 0; line < len(input); line++ {
		for col := 0; col < len(input[line]); col++ {

			c := input[line][col]
			if c != '*' {
				continue
			}

			var fgear, sgear int
			for adjline := line - 1; adjline <= line+1; adjline++ {
				for adjcol := col; adjcol <= col+1; adjcol++ {

					// fmt.Print(" check ", adjline+1, " ", adjcol+1)
					// bounce check
					if adjline < 0 || adjline > len(input) ||
						adjcol < 0 || adjcol > len(input[line]) {
						continue
					}

					for gpos, gearval := range g {

						if fgear != 0 && sgear != 0 {
							break
						}

						if gpos.row != adjline ||
							adjcol < gpos.colStart || adjcol > gpos.colEnd {
							continue
						}

						_, ok := used[gpos]

						if ok {
							continue
						}

						if fgear == 0 {
							fgear = int(gearval)
							used[gpos] = gearval
							continue
						}

						if sgear == 0 {
							sgear = int(gearval)
							used[gpos] = gearval
						}

					}

				}

			}

			if fgear != 0 && sgear != 0 {
				// fmt.Println("Sum of gears ", fgear, sgear, fgear*sgear)
				sum += fgear * sgear
			}

		}

	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	numstr := ""

	var g Gears = make(map[GearPos]GearValue)

	for line := 0; line < len(parsed); line++ {

		for col := 0; col < len(parsed[line]); col++ {
			c := parsed[line][col]

			if c >= '0' && c <= '9' {
				numstr += string(c)
			} else {
				if numstr != "" {
					n, err := strconv.Atoi(numstr)
					if err != nil {
						log.Fatalf("ERR converting string %s to number", numstr)
					}
					g.setGear(line, col-len(numstr), col, n)
				}
				numstr = ""
			}

		}

		if numstr != "" {
			n, err := strconv.Atoi(numstr)
			if err != nil {
				log.Fatalf("ERR converting string %s to number", numstr)
			}
			g.setGear(line, len(parsed[line])-len(numstr), len(parsed[line]), n)
			numstr = ""
		}

	}

	// for i, v := range g {
	// 	fmt.Println(i.row+1, i.colStart+1, i.colEnd+1, v)
	// }

	return g.sumgears(parsed)
}

func parseInput(input string) (ans []string) {
	ans = strings.Split(input, "\n")
	return ans
}
