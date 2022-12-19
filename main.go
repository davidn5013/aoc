package main

import (
	"fmt"

	"github.com/davidn5013/aoc/matrix"
)

func main() {
	m := matrix.NewWithSize(10, 10)
	m.SetValueInPos(matrix.Pos{5, 5}, 5)
	fmt.Println(m)

	cur, _ := m.NewCursor(matrix.Pos{0, 0})

	//create moves
	moves := []matrix.Move{}
	for i := 1; i <= 8; i++ {
		moves = append(moves, matrix.DiaDownRight)
	}

	wormlength := 5
	tailvalue, err := m.WormGetValue(cur, moves, wormlength)
	if err != nil {
		fmt.Println("Ops hitta wall")
	}
	fmt.Println(cur, tailvalue)

	// fill matrix with values?
	for row := range m {
		for col := range m {
			m[row][col] = col
		}
	}
	fmt.Println(m)
	fmt.Println(m.Reverse().Transpose().Reverse())
}
