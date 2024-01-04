package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davidn5013/aoc/cast"
	"github.com/davidn5013/aoc/scripts/aoc"
	"github.com/davidn5013/aoc/sol/d2307"
	"github.com/davidn5013/aoc/util"
)

// new solutios in sol catalog using modules instead of main files
// Format "Year,day Year,day" (day without begining zero)
var sols = solutions{
	sol{
		year:  "2023",
		day:   "7",
		part1: d2307.Part1,
		part2: d2307.Part2,
	},
}

type sol struct {
	year  string
	day   string
	part1 func(string) int
	part2 func(string) int
}

type solutions []sol

func RunAoc(year, day, filename, cookie string) {
	inputfile := inputfile(year, day, "input", filename)

	if !util.FileExists(inputfile) {
		if cookie != "" {
			aoc.GetInput(cast.ToInt(day), cast.ToInt(year), cookie)
		} else {
			log.Fatal("Need aoc cookie for download of missing inputfile")
		}
	}

	for _, s := range sols {
		if s.year == year && s.day == day {
			buf, _ := os.ReadFile(inputfile)
			input := string(buf)
			fmt.Printf("Part1: %d\nPart2: %d\n", s.part1(input), s.part2(input))
		}
	}

}

func inputfile(year, day, catalog, filename string) string {
	d := "0" + day
	f := year + "day" + d[len(d)-2:] + filename
	return filepath.Join(catalog, f)
}
