// matrix walker
/*
┌──────────────────────┐
│ │                    │
│ │                    │
│ │                    │
│ └──────┐             │
│        │             │
│        │             │
│        └───────┐     │
│                │     │
│        ┌───────┘     │
│        │             │
│        ▼───────►     │
│                      │
└──────────────────────┘
*/
package mat

import (
	"fmt"
	"strconv"
)

type heltal int

type Matris [][]heltal

func (m Matris) New(storlek int) *Matris {
	m = make([][]heltal, storlek)
	return &m
}

func (m Matris) Add(rad, kolumn, tal int) {
	m[rad][kolumn] = heltal(tal)
}

func (m Matris) Get(rad, kolumn int) int {
	return int(m[rad][kolumn])
}

var (
	Left  [2]int = [2]int{0, -1}
	Right [2]int = [2]int{0, 1}
	Up    [2]int = [2]int{-1, 0}
	Down  [2]int = [2]int{1, 0}

	cursor [2]heltal = [2]heltal{0, 0}
)

func (m Matris) SetCursorPos(rad, kolumn int) {
	cursor[0] = heltal(rad)
	cursor[1] = heltal(kolumn)
}

func (m Matris) GetCursorPos() (rad, kolumn int) {
	return int(cursor[0]), int(cursor[1])
}

func (m Matris) curValue() heltal {
	return m[cursor[0]][cursor[1]]
}

func (m Matris) MoveCurser(Direction [2]int, Moves int) {
	R, K := m.GetCursorPos()
	for Steps := 1; Steps <= Moves; Steps++ {
		K += Direction[1]
		R += Direction[0]
		if (R >= 0 && R <= len(m)-1) && (K >= 0 && K <= len(m[1])-1) {
			m.SetCursorPos(R, K)
		} else {
			break
		}
	}
}

func (m Matris) GetDirectionValue(Direction [2]int, Moves int) (ret []int) {
	R, K := m.GetCursorPos()
	for Steps := 1; Steps <= Moves; Steps++ {
		K += Direction[1]
		R += Direction[0]
		if (R >= 0 && R <= len(m)-1) && (K >= 0 && K <= len(m[1])-1) {
			m.SetCursorPos(R, K)
			ret = append(ret, int(m.curValue()))
		} else {
			break
		}
	}
	return ret
}

func (m Matris) String() string {
	res := ""
	for _, v := range m {
		for _, r := range v {
			res += strconv.Itoa(int(r))
		}
		res += "\n"
	}
	return res
}

func (m Matris) GoString() string {
	res := ""
	for _, v := range m {
		for _, i := range v {
			res += fmt.Sprintf("[%d]", int(i))
		}
		res += fmt.Sprintln()
	}
	return res
}
