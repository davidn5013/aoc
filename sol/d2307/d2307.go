// Package d2307 - Solution of Advent of code 2023 day 7 part 1 & part 2
package d2307

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	previewInput(input)
	return 0
}

func Part2(input string) int {
	return 0
}

func previewInput(input string) {
	inputlines := strings.Split(input, "\n")
	fmt.Println("Numbers of lines:", len(inputlines))
	fmt.Println("First 10 lines")
	for i, v := range inputlines {
		fmt.Println(v)
		if i > 4 {
			break
		}
	}
}
