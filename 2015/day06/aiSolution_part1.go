// ChatGPT solution for day 6
// prompt:
// work solution in go for advent of code 2015 day 6 part 1 using idiomatic go with comments. The correct answer is 569999.
//
// https://chat.openai.com/c/9e464bc5-ba1d-4b28-b460-d9b4afa146c9
//
//go:build:ignore
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Action represents the possible actions for each instruction.
type Action int

const (
	TurnOn Action = iota
	TurnOff
	Toggle
)

// Instruction represents a single instruction for manipulating lights.
type Instruction struct {
	Action Action
	Start  Point
	End    Point
}

// Point represents a point in 2D space.
type Point struct {
	X, Y int
}

// Grid represents the 1000x1000 grid of lights.
type Grid [][]bool

// NewGrid creates a new grid with all lights turned off.
func NewGrid(rows, cols int) Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}
	return grid
}

// ApplyInstruction applies the given instruction to the grid.
func (grid Grid) ApplyInstruction(inst Instruction) {
	for i := inst.Start.X; i <= inst.End.X; i++ {
		for j := inst.Start.Y; j <= inst.End.Y; j++ {
			switch inst.Action {
			case TurnOn:
				grid[i][j] = true
			case TurnOff:
				grid[i][j] = false
			case Toggle:
				grid[i][j] = !grid[i][j]
			}
		}
	}
}

// CountLightsOn counts the number of lights that are turned on.
func (grid Grid) CountLightsOn() int {
	count := 0
	for i := range grid {
		for _, on := range grid[i] {
			if on {
				count++
			}
		}
	}
	return count
}

// ParseInstruction parses a single line of input into an Instruction.
func ParseInstruction(line string) (Instruction, error) {
	// I made this smaller
	inst := Instruction{}

	// the order is very
	commands := []string{"turn on", "turn off", "toggle"}
	for idx, command := range commands {
		if strings.HasPrefix(line, command) {
			switch idx {
			case 0:
				inst.Action = TurnOn
			case 1:
				inst.Action = TurnOff
			case 2:
				inst.Action = Toggle
			default:
				return inst, fmt.Errorf("unknown instruction: %s", line)
			}

		}
		line = strings.TrimPrefix(line, command)
	}

	fmt.Sscanf(line, "%d,%d through %d,%d",
		&inst.Start.X,
		&inst.Start.Y,
		&inst.End.X,
		&inst.End.Y,
	)
	return inst, nil
}

// atoi is a helper function to convert a string to an integer.
func atoi(s string) int {
	val := 0
	fmt.Sscanf(s, "%d", &val)
	return val
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := NewGrid(1000, 1000)

	// Process each line of input
	for scanner.Scan() {
		line := scanner.Text()
		inst, err := ParseInstruction(line)
		if err != nil {
			fmt.Println("Error parsing instruction:", err)
			return
		}
		grid.ApplyInstruction(inst)
	}

	// Count and print the number of lights that are turned on
	fmt.Println("Number of lights on:", grid.CountLightsOn())
}
