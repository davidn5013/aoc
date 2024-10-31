// Package dYYDD - Solution of Advent of code 20YY day D part 1 & part 2
package dYYDD

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
	fmt.Println("First 5 lines or less:")
	for i, v := range inputlines {
		fmt.Println(v)
		if i > 3 {
			break
		}
	}
}
