// Package matrix create [][]int matrix array with cursor methods
package matrix

import "fmt"

// Type pos = position in matrix, used for set and get value and for placing cursor
type Pos struct {
	Row    int
	Column int
}

// Type Move is move in matrix
/*
* Declare like so:
moves := []Move{
	matrix.Right,
	matrix.Down,
	matrix.Down,
}
*/
type Move struct {
	Row    int
	Column int
}

var (
	// Nop Move don't but read value
	Nop Move = Move{
		Row:    0,
		Column: 0,
	}
	// Move Right
	Right Move = Move{
		Row:    0,
		Column: 1,
	}
	// Move Left
	Left Move = Move{
		Row:    0,
		Column: -1,
	}
	// Move Up
	Up Move = Move{
		Row:    -1,
		Column: 0,
	}
	// Move Down
	Down Move = Move{
		Row:    1,
		Column: 0,
	}
	// Move Diagonal Up and to the Right
	DiaUpRight Move = Move{
		Row:    -1,
		Column: 1,
	}
	// Move Diagonal Up and to the Left
	DiaUpLeft Move = Move{
		Row:    -1,
		Column: -1,
	}
	// Move Diagonal Down and to the Right
	DiaDownRight Move = Move{
		Row:    1,
		Column: 1,
	}
	// Move Diagonal Down and to the Left
	DiaDownLeft Move = Move{
		Row:    1,
		Column: -1,
	}
)

type matrix [][]int

type cursor struct {
	row    int
	column int
}

// New create a new empty Matrix used for Moves
func New() *matrix {
	m := matrix{}
	return &m
}

// NewWithSize create a new Matrix used for Moves
func NewWithSize(rows, columns int) matrix {
	var res matrix
	for i := 0; i < rows; i++ {
		var j []int
		for i := 0; i < columns; i++ {
			j = append(j, 0)
		}
		res = append(res, j)
	}
	return res
}

// ExpandMatrix grow empty matrix
// to shrink make new and copy
func (m matrix) ExpandMatrixSize(rows, columns int) {
	for i := 0; i < rows; i++ {
		var j []int
		for i := 0; i < columns; i++ {
			j = append(j, 0)
		}
		m = append(m, j)
	}
}

// SetValueInPos takes a position and value and place in matrix
func (m matrix) SetValueInPos(p Pos, value int) bool {
	if p.Row >= len(m) || p.Row < 0 {
		return false
	}
	if p.Column >= len(m[0]) || p.Column < 0 {
		return false
	}
	m[p.Row][p.Column] = value
	return true
}

// GetValueInPos get the value for a position
func (m matrix) GetValueInPos(p Pos, value int) (int, bool) {
	if p.Row >= len(m) || p.Row < 0 {
		return 0, false
	}
	if p.Column >= len(m[0]) || p.Column < 0 {
		return 0, false
	}
	return m[p.Row][p.Column], true
}

// NewCursor create a new cursor for moves
func (m matrix) NewCursor(p Pos) (cur *cursor, err error) {
	r, c := p.Row, p.Column
	if (r >= 0 && r <= len(m)) && (c >= 0 && c < len(m[1])) {
		cur := cursor{
			row:    r,
			column: c,
		}
		return &cur, nil
	}
	return nil, fmt.Errorf("position outside of matrix")
}

// GetValueUnderCursor int
func (m matrix) GetValueUnderCursor(cur cursor) int {
	return m[cur.row][cur.column]
}

// SetValueUnderCursor int
func (m matrix) SetValueUnderCursor(cur cursor, value int) {
	m[cur.row][cur.column] = value
}

// CursorPosChange
func (m matrix) CursorPosChange(cur *cursor, p Pos) error {
	newRow := cur.row + p.Row
	newColumn := cur.column + p.Column
	if (newRow >= 0 && newRow < len(m)) && (newColumn >= 0 && newColumn < len(m[1])) {
		cur.row, cur.column = newRow, newColumn
		return nil
	}
	return fmt.Errorf("position outside of matrix")
}

// DoMovesSetValue follow moves in []Move and set one value
func (m matrix) DoMovesSetValue(cur *cursor, moves []Move, value int) error {
	for _, move := range moves {
		if err := m.CursorPosChange(cur, Pos(move)); err != nil {
			return err
		}
		m.SetValueUnderCursor(*cur, value)
	}
	return nil
}

// DoMoveGetValue follow moves in []Move and get all values
func (m matrix) DoMovesGetValue(cur *cursor, moves []Move) (ret []int, err error) {
	for _, move := range moves {
		if err := m.CursorPosChange(cur, Pos(move)); err != nil {
			return ret, err
		}
		ret = append(ret, m.GetValueUnderCursor(*cur))
	}
	return ret, nil
}

// Worm follow moves in []Move and return []int of x head moves. Reveres []int
func (m matrix) WormGetValue(cur *cursor, moves []Move, wormlength int) (ret []int, err error) {
	rev, err := m.DoMovesGetValue(cur, moves)
	ret = reverseArr(rev)

	if wormlength > len(ret)+1 {
		wormlength = len(ret) + 1
	}

	ret = ret[0:wormlength]

	if err != nil {
		return ret, err
	}
	return ret, err
}

func reverseArr(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func (m matrix) Transpose() (res matrix) {
	for row := range m {
		line := []int{}
		for col := range m {
			line = append(line, m[col][row])
		}
		res = append(res, line)
	}
	return res
}

func (m matrix) Reverse() (res matrix) {
	for row, rowvalue := range m {
		line := []int{}
		for i := len(rowvalue) - 1; i >= 0; i-- {
			line = append(line, m[row][i])
		}
		res = append(res, line)
	}
	return res
}

// Stringer for matrix, fmt.Println
func (m matrix) String() string {
	res := ""
	for _, v := range m {
		for _, i := range v {
			res += fmt.Sprintf("%d", i)
		}
		res += fmt.Sprintf("\n")
	}
	return res
}

// GoString for matrix, fmt.Printf("%#v\n",matrix)
func (m matrix) GoString() string {
	res := fmt.Sprintf("{\n")
	for _, v := range m {
		for _, i := range v {
			res += fmt.Sprintf("{")
			res += fmt.Sprintf("%d, ", i)
		}
		res += fmt.Sprintf("},\n")
	}
	res += fmt.Sprintf("}\n")
	return res
}
