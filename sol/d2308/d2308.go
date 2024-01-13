// Package d2308 - Solution of Advent of code 2023 day 8 part 1 & part 2
package d2308

import (
	"strings"
)

func parse(input string) ([]byte, map[string][]string) {
	parts := strings.Split(input, "\n\n")
	instructions := []byte(parts[0])

	nodes := make(map[string][]string, 10)
	for _, v := range strings.Split(parts[1], "\n") {
		if len(v) == 0 {
			break
		}
		nodes[v[:3]] = []string{v[7:10], v[12:15]}
	}
	return instructions, nodes
}

func Part1(input string) int {
	instructions, nodes := parse(input)

	current := "AAA"
	cnt := 0
	for {
		ip := cnt % len(instructions)

		switch instructions[ip] {
		case 'L':
			current = nodes[current][0]
		case 'R':
			current = nodes[current][1]
		}

		if current == "ZZZ" {
			break
		}

		cnt++

	}

	return cnt + 1

}

func Part2(input string) int {
	return 0
}
