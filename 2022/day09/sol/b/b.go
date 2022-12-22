package b

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/davidn5013/snickerboa"
)

type pos struct {
	row int
	col int
}

// rope has ten knots
type rope struct {
	head pos // knot 1 9 tails
	tail [9]pos
}

const (
	boardrows = 10
	boardcols = 40
)

var (
	w = rope{
		// head: pos{row: 9, col: 0},
		// tail: pos{row: 9, col: 0},
		head: pos{row: boardrows, col: boardcols / 2},
	}

	lastvisit = make(map[pos]struct{})
)

func Sol() int {
	// import moves in string to struct moves
	// smalle exampel
	// moves := createmoves(`R 4
	// U 4
	// L 3
	// D 1
	// R 4
	// D 1
	// L 5
	// R 2`)

	// exampel 2
	// moves := createmoves(`R 5
	// U 8
	// L 8
	// D 3
	// R 17
	// D 10
	// L 25
	// U 20`)

	moves := createmoves(movesString)

	// set visit to first position
	lastvisit[pos{row: w.head.row, col: w.head.col}] = struct{}{}

	// tail to head position
	for idx := range w.tail {
		w.tail[idx].row = w.head.row
		w.tail[idx].col = w.head.col
	}

	// This just to visualization
	// area, _ := pterm.DefaultArea.Start() // Start the Area printer, with the Center option.

	snickerboa.ClearScreen()

	for idx, m := range moves {
		// newBoard := createboard(w)

		// No need just for visualization
		// area.Update(newBoard)

		doNextMove(pos{row: m.row, col: m.col}, &w)
		fmt.Printf("Move %d\n", idx)
		// time.Sleep(100 * time.Millisecond)
	}

	// fmt.Println(lastvisit)
	return len(lastvisit)
}

// doNextMove is the import function here this where all import stuff gets done
func doNextMove(nx pos, w *rope) {
	// move head
	w.head.row += nx.row
	w.head.col += nx.col

	// first tail follow head, next tail follow the tail position before
	prevPos := w.head
	dr := prevPos.row - w.tail[0].row
	dc := prevPos.col - w.tail[0].col
	for idx := range w.tail {
		if idx > 0 {
			prevPos = w.tail[idx-1]
			dr = w.tail[idx-1].row - w.tail[idx].row
			dc = w.tail[idx-1].col - w.tail[idx].col
		}

		// skip moving tail if in touching distans of head
		if abs(dr) <= 1 && abs(dc) <= 1 {
			return
		}

		step := pos{}
		step.row = setStep(dr, prevPos.row, w.tail[idx].row)
		step.col = setStep(dc, prevPos.col, w.tail[idx].col)
		w.tail[idx].col += step.col
		w.tail[idx].row += step.row

	}

	// for idx := range w.tail {
	// 	lastvisit[pos{row: w.tail[idx].row, col: w.tail[idx].col}] = struct{}{}
	// }

	lastvisit[pos{row: w.tail[8].row, col: w.tail[8].col}] = struct{}{}

}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func setStep(distens int, comp, thispos int) (step int) {
	if distens != 0 {
		// Compile time assertion of comp-thispos !=0
		i := comp - thispos
		var _ uint = uint(i)/uint(abs(i)) - 1
		return (i) / abs(i)
	}
	return 0
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

func createboard(w rope) (newboard string) {
	const (
		xtrRows = 10
		xtrCols = 0
	)
	var b strings.Builder
	b.Grow((boardrows + xtrRows) * (boardcols + xtrCols + 1))
	for i := 0; i <= boardrows+xtrRows; i++ {
		for j := 0; j <= boardcols+xtrCols; j++ {
			if w.head.row == i && w.head.col == j {
				b.WriteString("H")
				continue
			}
			for idx, tailpos := range w.tail {
				if tailpos.row == i && tailpos.col == j {
					b.WriteString(strconv.Itoa(idx + 1))
					goto nextLoop
				}
			}

			if _, ok := lastvisit[pos{row: i, col: j}]; ok {
				b.WriteString("#")
				continue
			}
			b.WriteString(".")

			if j == boardcols {
				b.WriteString("\n")
			}

		nextLoop:
		}
	}
	return b.String()
}
