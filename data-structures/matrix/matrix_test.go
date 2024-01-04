package matrix

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/davidn5013/aoc/matrix"
)

// Create 3x3 matrix set middel value
// set value in u
// read middel row
func TestNew(t *testing.T) {
	m := matrix.NewWithSize(3, 3)
	m.SetValueInPos(matrix.Pos{1, 1}, 1)
	output := strings.TrimSpace(fmt.Sprintln(m))

	cur, err := m.NewCursor(matrix.Pos{0, 0})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	move := []matrix.Move{
		matrix.Nop,
		matrix.Down,
		matrix.Down,
		matrix.Right,
		matrix.Right,
		matrix.Up,
		matrix.Up,
	}
	if err := m.DoMovesSetValue(cur, move, 5); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output += strings.TrimSpace(fmt.Sprintln(m))

	cur1, _ := m.NewCursor(matrix.Pos{0, 1})
	move2 := []matrix.Move{
		matrix.Nop,
		matrix.Down,
		matrix.Down,
	}
	value, _ := m.DoMovesGetValue(cur1, move2)
	output += strings.TrimSpace(fmt.Sprintln(value))

	want := `000
010
000505
515
555[0 1 5]`

	if output != want {
		t.Errorf("Did not get the want result testing DoMoveSetValue :\n %s\n", output)
	}
}

func TestWorm(t *testing.T) {
	m := matrix.NewWithSize(10, 10)
	m.SetValueInPos(matrix.Pos{5, 5}, 5)
	output := strings.TrimSpace(fmt.Sprintln(m))

	cur, _ := m.NewCursor(matrix.Pos{0, 0})

	//create moves
	moves := []matrix.Move{}
	for i := 1; i <= 8; i++ {
		moves = append(moves, matrix.DiaDownRight)
	}

	wormlength := 5
	tailvalue, err := m.WormGetValue(cur, moves, wormlength)
	if err != nil {
		output += strings.TrimSpace(fmt.Sprintln("Ops hitta wall"))
	}
	output += strings.TrimSpace(fmt.Sprintln(tailvalue))

	want := `0000000000
0000000000
0000000000
0000000000
0000000000
0000050000
0000000000
0000000000
0000000000
0000000000[0 0 0 5 0]`
	if output != want {
		t.Errorf("Did not get the want result testing Worm :\n %s\n", output)
	}
}

func TestTransposeReverse(t *testing.T) {
	m := matrix.NewWithSize(10, 10)
	for row := range m {
		for col := range m {
			m[row][col] = col
		}
	}
	output := strings.TrimSpace(fmt.Sprintln(m))
	output += strings.TrimSpace(fmt.Sprintln(m.Reverse().Transpose().Reverse()))

	want := `0123456789
0123456789
0123456789
0123456789
0123456789
0123456789
0123456789
0123456789
0123456789
01234567899999999999
8888888888
7777777777
6666666666
5555555555
4444444444
3333333333
2222222222
1111111111
0000000000`

	if output != want {
		t.Errorf("Did not get the want result testing Transpose and Reverse:\n %s\n", output)
	}
}
