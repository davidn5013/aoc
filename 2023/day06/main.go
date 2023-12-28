package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/davidn5013/aoc/cast"
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

func part1(input string) int {
	matches := parseInput(input)
	multiplier := 1
	for _, m := range matches {
		dist := m.dist
		time := m.time
		winners := 0
		// fmt.Println("Dist and time :", dist, time)
		for i := 0; i <= time; i++ {
			roundspeed := i
			roundtime := time - i
			rounddist := roundtime * roundspeed
			if rounddist > dist {
				winners++
				// fmt.Println("holdtime, rounddist", roundspeed, rounddist, winners)
			}

		}
		if winners > 0 {
			multiplier *= winners
		}
	}

	// fmt.Println(multiplier)
	return multiplier
}

func part2(input string) int {
	m := parseInput2(input)
	multiplier := 1
	dist := m[0].dist
	time := m[0].time
	winners := 0
	// fmt.Println("Dist and time :", dist, time)
	for i := 0; i <= time; i++ {
		roundspeed := i
		roundtime := time - i
		rounddist := roundtime * roundspeed
		if rounddist > dist {
			winners++
			// fmt.Println("holdtime, rounddist", roundspeed, rounddist, winners)
		}

	}
	if winners > 0 {
		multiplier *= winners
	}

	// fmt.Println(multiplier)
	return multiplier
}

type Matches struct {
	time int
	dist int
}

func parseInput(input string) (ans []Matches) {
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])
	dists := strings.Fields(lines[1])
	if len(times) != len(dists) {
		log.Fatal("ERR not same length on times and dists")
	}
	var m []Matches
	m = make([]Matches, len(times)-1)
	for i := 1; i < len(times); i++ {
		m[i-1] = Matches{time: cast.ToInt(times[i]), dist: cast.ToInt(dists[i])}
	}

	return m
}

func parseInput2(input string) (ans []Matches) {
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])
	dists := strings.Fields(lines[1])
	t2 := strings.Join(times[1:], "")
	d2 := strings.Join(dists[1:], "")

	var m []Matches
	m = make([]Matches, 1)
	m[0] = Matches{time: cast.ToInt(t2), dist: cast.ToInt(d2)}

	// fmt.Println(m)
	return m
}
