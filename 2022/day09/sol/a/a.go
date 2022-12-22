package a

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	row int
	col int
}

type worm struct {
	head pos
	tail pos
}

const boardrows = 80
const boardcols = 100

var (
	w = worm{
		// head: pos{row: 9, col: 0},
		// tail: pos{row: 9, col: 0},
		head: pos{row: boardrows / 2, col: boardcols / 2},
		tail: pos{row: boardrows / 2, col: boardcols / 2},
	}
	lastvisit = make(map[pos]struct{})
)

func Sol() int {
	moves := createmoves(movesString)
	lastvisit[pos{row: w.tail.row, col: w.tail.col}] = struct{}{}

	for _, m := range moves {
		doNextMove(pos{row: m.row, col: m.col}, &w)
	}

	return len(lastvisit)
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func doNextMove(nx pos, w *worm) {
	// move head
	w.head.row += nx.row
	w.head.col += nx.col

	// diff row col
	dr := w.head.row - w.tail.row
	dc := w.head.col - w.tail.col

	// skip moving tail if in touching distans of head
	if abs(dr) <= 1 && abs(dc) <= 1 {
		return
	}

	// move tail
	// if tail=9,0 head=7,0 (9-7)=-2/abs 2 = 1 move row on
	// if tail=9,3 head=7,4 row +0 col (3-4)=-1 / 1 =-1
	step := pos{}

	if dr == 0 {
		step.row = 0
	} else {
		step.row = (w.head.row - w.tail.row) / abs(w.head.row-w.tail.row)
	}

	if dc == 0 {
		step.col = 0
	} else {
		step.col = (w.head.col - w.tail.col) / abs(w.head.col-w.tail.col)
	}

	w.tail.col += step.col
	w.tail.row += step.row
	lastvisit[pos{row: w.tail.row, col: w.tail.col}] = struct{}{}
}

func createmoves(s string) (p []pos) {
	m := make(map[string]pos)
	m["U"] = pos{row: -1, col: 0}
	m["D"] = pos{row: 1, col: 0}
	m["R"] = pos{row: 0, col: 1}
	m["L"] = pos{row: 0, col: -1}

	for _, line := range strings.Split(s, "\n") {
		cell := strings.Fields(line)
		moveRune := cell[0]
		movesInRune := cell[1]
		newmoves, err := strconv.Atoi(movesInRune)
		if err != nil {
			fmt.Printf("failed to convert to int on %s %s\n", string(cell[0]), string(cell[1]))
			os.Exit(1)
		}
		for i := 0; i < newmoves; i++ {
			np := pos{row: m[moveRune].row, col: m[moveRune].col}
			p = append(p, np)
		}
	}
	return p
}

func createboard(w worm) (newboard string) {
	for i := 0; i <= boardrows; i++ {
		for j := 0; j <= boardcols; j++ {
			if _, ok := lastvisit[pos{row: i, col: j}]; ok {
				newboard += "#"
				continue
			}
			if w.head.row == i && w.head.col == j {
				newboard += "H"
				continue
			}
			if w.tail.row == i && w.tail.col == j {
				newboard += "T"
				continue
			}

			newboard += "."

			if j == boardcols {
				newboard += "\n"

			}
		}
	}
	return newboard
}
