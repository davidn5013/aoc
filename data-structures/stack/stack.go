/**
* Source :  https://github.com/lynerist/Advent-of-code-2022-golang.git
  All love to lynerist for work
*/

// Stack made of array of rune
package util

type Stack struct {
	elements []rune
}

func (s *Stack) Push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *Stack) PushMoves(r []rune) {
	s.elements = append(s.elements, r...)
}

func (s *Stack) Pop() (r rune) {
	r = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *Stack) PopMoves(n int) (r []rune) {
	r = s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return
}

func (s *Stack) AddToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s Stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}
