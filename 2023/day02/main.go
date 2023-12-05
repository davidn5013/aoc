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

func part1(input string) (total int) {
	games := parseInput(input)
	for _, game := range games {
		possible := true
		for _, subgame := range game.subs {
			util.Debugf("%#v\n", subgame)
			if subgame.red > 12 || subgame.green > 13 || subgame.blue > 14 {
				possible = false
			}
		}
		if possible {
			total += game.nr
			util.Infof("\t add game nr %d\n", game.nr)
		}

	}

	return total
}

func part2(input string) (total int) {
	games := parseInput(input)
	for _, game := range games {
		var minRed, minGreen, minBlue int
		for _, subgame := range game.subs {
			// dbprint("%#v\n", subgame)
			if subgame.red > minRed {
				minRed = subgame.red
			}
			if subgame.green > minGreen {
				minGreen = subgame.green
			}
			if subgame.blue > minBlue {
				minBlue = subgame.blue
			}

		}

		total += minRed * minGreen * minBlue
		util.Infof(" > Gamenr %d: %d*%d*%d \t=%d\n", game.nr, minRed, minGreen, minBlue, minRed*minGreen*minBlue)

	}

	return total
}

type sub struct {
	red   int
	green int
	blue  int
}

type game struct {
	nr   int
	subs []sub
}

type games []game

func parseInput(input string) (newgames games) {
	for _, line := range strings.Split(input, "\n") {
		gameparts := strings.Split(line, ":")
		gamenr, err := strconv.Atoi(strings.Split(gameparts[0], "Game ")[1])
		if err != nil {
			log.Fatalf("Can't convert gamenr on line %s\n", line)
		}
		subsets := strings.Split(gameparts[1], ";")

		util.Debugf(" %#v\n", gamenr)
		util.Debugf(" %#v\n", subsets)

		var s []sub
		for _, line := range subsets {

			parts := strings.Split(line, ",")
			var r, g, b int

			for _, part := range parts {
				var cubes int
				var color string
				fmt.Sscanf(part, "%d %s", &cubes, &color)

				switch color {
				case "red":
					r = cubes
				case "green":
					g = cubes
				case "blue":
					b = cubes
				default:
					log.Fatalf("Unknow color in line %s\n", line)
				}

				util.Debugf(" \t r=%d g=%d b=%d\n", r, g, b)

			}
			s = append(s, sub{r, g, b})
		}

		newgames = append(newgames, game{nr: gamenr, subs: s})
	}
	return newgames

}
