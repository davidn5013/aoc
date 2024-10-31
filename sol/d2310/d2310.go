// Package d2310 - Solution of Advent of code 2023 day 10 part 1 & part 2
package d2310

import (
	"fmt"
	"strings"

	"github.com/davidn5013/aoc/util"
)

type Pos struct {
	col, row int
}

func Part1(input string) int {
	previewInput(input)
	rows := strings.Split(input, "\n")

	maxRow := len(rows)
	maxCol := len(rows[0])

	visited := make(map[Pos]struct{}, maxRow)

	var start Pos

	for row, line := range rows {
		for col, cell := range line {
			if cell == 'S' {
				start := Pos{col, row}
				visited[start] = struct{}{}
			}
		}
	}
	util.Debugf("Start position S is :%v\n", start)

	moves := []Pos{
		{-1, 0}, // Col Upp
		{1, 0},  // Col Under
		{0, -1}, // Row Left
		{0, 1},  // Row Right
	}

	stepCount := 1
	var nextStack []Pos
	nextStack = append(nextStack, start)

	for {
		currPos := nextStack[0]
		visited[currPos] = struct{}{}
		for _, move := range moves {

			nextPos := Pos{
				row: currPos.row + move.row,
				col: currPos.col + move.col,
			}

			if nextPos.row < 0 || nextPos.row > maxRow || nextPos.col < 0 || nextPos.col > maxCol {
				break
			}

			if _, found := visited[nextPos]; found {
				break
			}

			val := rows[nextPos.row][nextPos.col]
			north := currPos.col > nextPos.col
			south := currPos.col < nextPos.col
			west := currPos.row > nextPos.row
			east := currPos.row < nextPos.row
			appendNxtPos := false

			/*
			   . is ground; there is no pipe in this tile.
			*/
			switch {
			case (north || south) && val == '|': // | is a vertical pipe connecting north and south.
				appendNxtPos = true
			case (west || east) && val == '-': // - is a horizontal pipe connecting east and west.
				appendNxtPos = true
			case (south || west) && val == 'L': // L is a 90-degree bend connecting north and east.
				appendNxtPos = true
			case (south || east) && val == 'J': // J is a 90-degree bend connecting north and west.
				appendNxtPos = true
			case (east || north) && val == '7': //  7 is a 90-degree bend connecting south and west.
				appendNxtPos = true
			case (west || north) && val == 'F': // F is a 90-degree bend connecting south and east.
				appendNxtPos = true
			}

			fmt.Printf("Add add next : %t, %v start:%t\n", appendNxtPos, nextPos, nextPos != start)

			if appendNxtPos && nextPos != start {
				stepCount++
				nextStack = append(nextStack, nextPos)
			}

		}

		if len(nextStack) > 0 {
			nextStack = nextStack[1:]
		} else {
			break
		}

	}

	return stepCount
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
