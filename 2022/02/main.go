// Package main https://adventofcode.com/2022/day/2
// --- Day 2: Rock Paper Scissors ---
package main

import (
	_ "embed"
	"fmt"
	"log"
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
	input := strings.Split(strings.TrimSpace(inputFile), "\n")
	fmt.Printf("Solution 1 %d\n", sol1_1([]string{
		"A Y",
		"B X",
		"C Z",
	}))
	fmt.Printf("Solution 1 %d\n", sol1_1(input))
	fmt.Printf("Solution 1 ver 2 %d\n", sol1_2([]string{
		"A Y",
		"B X",
		"C Z",
	}))
	fmt.Printf("Solution 1 ver 2 %d\n", sol1_2(input))
	fmt.Printf("Solution 2 %d\n", sol2([]string{
		"A Y",
		"B X",
		"C Z",
	}))
	fmt.Printf("Solution 2 %d\n", sol2(input))
}

// Solution part 1 -
// X for Rock, Y for Paper, and Z for Scissors. A for Rock, B for
// Paper, and C for Scissors
// The score for a single round is the score for the shape you
// selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the
// score for the outcome of the round (0 if you lost, 3 if the round
// was a draw, and 6 if you won).
// In: A Y
// In the first round, your opponent will choose Rock (A), and
// you should choose Paper (Y). This ends in a win for you with a
// score of 8 (2 because you chose Paper + 6 because you won).
// 'A' | 'Y' Y winner
// I took me the hole day to understand that it was 3+3 on draw and not
// the card score
// Is this lacking in English understanding. I hope Tsoding is going to
// do Aoc so I see how he does
// This was the first version and it just harder to read
func sol1_1(lines []string) int {
	var (
		totalPoints int

		translate = map[rune]rune{
			'X': 'A',
			'Y': 'B',
			'Z': 'C',
		}
		value = map[rune]int{
			'A': 1,
			'B': 2,
			'C': 3,
		}
		win = map[string]struct{}{
			"AB": struct{}{}, // R vs P win
			"BC": struct{}{}, // P vs S win
			"CA": struct{}{}, // S vs P win
		}
	)
	const winpoint = 6

	for nr, line := range lines {
		opponent := rune(line[0])
		yourplay := translate[rune(line[2])]
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

		point, ok := value[yourplay]
		if !ok {
			log.Printf("ERR: Unable to match string %s on line %d\n", line, nr+1)
			return 0
		}
		switch result {
		case 'w':
			point += winpoint
		case 'd':
			point += winpoint / 2
		case 'l':
			point += 0
		default:
			log.Printf("ERR: Result is = %s line %d\n", string(result), nr+1)
			return 0
		}
		Debug(debugmode, "point %d on play %s ", point, string(yourplay))
		Debug(debugmode, "match point : %d\n", point)
		totalPoints += point
	}
	return totalPoints
}

func sol1_2(lines []string) (totScore int) {
	var (
		getScore = map[string]int{
			"A X": 3 + 1, // R vs R draw 1
			"A Y": 6 + 2, // R vs P win +6
			"A Z": 0 + 3, // R vs S lose
			"B X": 0 + 1, // P vs R lose
			"B Y": 3 + 2, // P vs P draw 2
			"B Z": 6 + 3, // P vs S win +6
			"C X": 6 + 1, // S vs R win +6
			"C Y": 0 + 2, // S vs P lose
			"C Z": 3 + 3, // S vs S draw 3
		}
	)
	for nr, line := range lines {
		score, ok := getScore[line]
		if !ok {
			log.Printf("ERR: Unable to match string %s on line %d\n", line, nr+1)
			return 0
		}
		totScore += score
		Debug(debugmode, "Line %d with %s point %d\n", nr+1, line, score)
	}
	return totScore
}

func sol2(lines []string) (totScore int) {
	var (
		getScore = map[string]int{
			"A X": 0 + 3, // R vs S Need to lose
			"A Y": 3 + 1, // R vs R Need to draw
			"A Z": 6 + 2, // R vs P Need to win
			"B X": 0 + 1, // P vs R Need to lose
			"B Y": 3 + 2, // P vs P Need to draw
			"B Z": 6 + 3, // P vs S Need to win
			"C X": 0 + 2, // S vs P Need to lose
			"C Y": 3 + 3, // S vs S Need to draw
			"C Z": 6 + 1, // S vs R Need to win
		}
	)
	for nr, line := range lines {
		score, ok := getScore[line]
		if !ok {
			log.Printf("ERR: Unable to match string %s on line %d\n", line, nr+1)
			return 0
		}
		totScore += score
		Debug(debugmode, "Line %d with %s point %d\n", nr+1, line, score)
	}
	return totScore
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
