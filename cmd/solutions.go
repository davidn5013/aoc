package cmd

import (
	"github.com/davidn5013/aoc/sol/d2307"
	"github.com/davidn5013/aoc/sol/d2308"
	"github.com/davidn5013/aoc/sol/d2309"
	"github.com/davidn5013/aoc/sol/d2310"
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
	sol{
		year:  "2023",
		day:   "8",
		part1: d2308.Part1,
		part2: d2308.Part2,
	},
	sol{
		year:  "2023",
		day:   "9",
		part1: d2309.Part1,
		part2: d2309.Part2,
	},
	sol{
		year:  "2023",
		day:   "10",
		part1: d2310.Part1,
		part2: d2310.Part2,
	},
}

type sol struct {
	year  string
	day   string
	part1 func(string) int
	part2 func(string) int
}

type solutions []sol
