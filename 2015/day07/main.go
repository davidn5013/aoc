package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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

func part1(input string) string {
	registery := make(map[string]int16, 100)
	for _, line := range strings.Split(input, "\n") {
		log.Println(line)
		var tval int16
		switch {
		case strings.Contains(line, "AND"):
			var reg1, reg2 string
			fmt.Sscanf(line, "%s AND %s", &reg1, &reg2)
			tval = registery[reg1] & registery[reg2]
		case strings.Contains(line, "OR"):
			var reg1, reg2 string
			fmt.Sscanf(line, "%s OR %s", &reg1, &reg2)
			tval = registery[reg1] | registery[reg2]
		case strings.Contains(line, "->"):
			var val int16
			var reg string
			fmt.Sscanf(line, "%d -> %s", &val, &reg)
			log.Println(":", val, reg)
			if val == 0 {
				registery[reg] = tval
			} else {
				registery[reg] = val
			}
		}
	}
	return ""
}

func part2(input string) int {
	return 0
}
