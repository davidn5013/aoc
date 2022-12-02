// Package main https://adventofcode.com/2022/day/2
// --- Day 2: Rock Paper Scissors ---
package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const debugmode = false

// -- I find it easiest just to embed this one
//
//go:embed input.txt
var inputFile string

func main() {
	var ()
	desc2022day1()

	// fmt.Printf("%d\n", mostSol1([]string{
	// 	"A Y",
	// 	"B X",
	// 	"C Z",
	// }))

	input := strings.Split(strings.TrimSpace(inputFile), "\n")
	fmt.Printf("%d\n", mostSol1(input))

	// fmt.Printf("%d\n", sol2([]string{""}))
}

// Solution part 1 -
// X for Rock, Y for Paper, and Z for Scissors. A for Rock, B for Paper, and C for Scissors
// The score for a single round is the score for the shape you
// selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the
// score for the outcome of the round (0 if you lost, 3 if the round
// was a draw, and 6 if you won).
// In: A Y
// In the first round, your opponent will choose Rock (A), and
// you should choose Paper (Y). This ends in a win for you with a
// score of 8 (2 because you chose Paper + 6 because you won).
// 'A' | 'Y' Y winner
func mostSol1(lines []string) int {
	var (
		points int

		value = map[rune]int{
			'A': 1,
			'B': 2,
			'C': 3,
		}
		runelookup = map[rune]rune{
			'X': 'A',
			'Y': 'B',
			'Z': 'C',
		}
		win = map[string]struct{}{
			"AB": struct{}{},
			"BC": struct{}{},
			"CB": struct{}{},
		}
	)

	for i, line := range lines {

		opponent := rune(line[0])
		yourplay, ok := runelookup[rune(line[2])]
		if !ok {
			panic(fmt.Errorf("ERR can find letter for %s in line %d ", string(line[2]), i))
		}
		result := 'l'

		if opponent == yourplay {
			result = 'd'
		} else {
			lookup := string(opponent) + string(yourplay)
			// log.Printf("%s\n", string(lookup))
			_, ok := win[lookup]
			if ok {
				result = 'w'
			}
		}

		Debug(debugmode, "%s%s%s ", string(opponent), string(yourplay), string(result))

		point := value[yourplay]
		switch result {
		case 'w':
			point += 6
		case 'd':
			point += 3
		}

		Debug(debugmode, "point %d on play %s ", point, string(yourplay))
		Debug(debugmode, "match point : %d\n", point)
		points += point
	}

	return points
}

// // Solution part 2 -
// func sol2(numberString []string) int {
// 	var ()
// 	return 0
// }

// Helper functions

// Debug is fmt.Printf false or true
func Debug(b bool, format string, v ...any) {
	if b {
		fmt.Printf(format, v...)
	}
}

func desc2022day1() {
	fmt.Println(``)

}
