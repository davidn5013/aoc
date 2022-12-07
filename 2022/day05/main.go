// NOT My solution kredit goes to:
// https://github.com/lynerist/Advent-of-code-2022-golang.git
/*
* Thank to lynerist nice solution, all love for help me keep on pratics
* my golang skills
 */
// Package main https://adventofcode.com/2022/day/5
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davidn5013/aoc/utl"
)

// More debug text 0-10 : 10 Max
const debuglvl = 0

// main Run Solution for Advent of Code
func main() {
	fmt.Println("Advent of code day 5")

	timer := utl.SetTimer()
	fmt.Printf("Solution 1 (test) = %s\n", sol1("input_test.txt", 3))
	fmt.Printf("Solution 1        = %s\n", sol1("input.txt", 9))
	log.Printf("took %s", timer())

	// timer = utl.SetTimer()
	fmt.Printf("Solution 2 (test) = %s\n", sol2("input_test.txt", 3))
	fmt.Printf("Solution 2        = %s\n", sol2("input.txt", 9))
	log.Printf("took %s", timer())
}

// sol1 Solution part 1
func sol1(inputFile string, s int) (res string) {
	//Read input file
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	//create slice of stacks
	stacks := make([]utl.Stack, s)

	//Parsing the input
	sc.Scan()
	// for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
	for !strings.Contains(sc.Text(), " 1   2   3 ") {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].AddToBottom(r)
			}
		}
		sc.Scan()
	}
	//Read empty line
	sc.Scan()

	for sc.Scan() {
		var toMove, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

		//Move elements one by one
		for move := 0; move < toMove; move++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	for _, s := range stacks {
		res += string(s.Pop())
	}
	return res

}

// sol2 Solution part 2
func sol2(inputFile string, s int) (res string) {
	utl.Debug(debuglvl >= 1, "Starting %s\n", utl.CurrFuncName())

	var ()

	//Read input file
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	//create slice of stacks
	stacks := make([]utl.Stack, s)

	//Parsing the input
	sc.Scan()
	// for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
	for !strings.Contains(sc.Text(), " 1   2   3 ") {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].AddToBottom(r)
			}
		}
		sc.Scan()
	}
	//Read empty line
	sc.Scan()
	for sc.Scan() {
		var toMove, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)
		//Move a block of elements
		stacks[to-1].PushMoves(stacks[from-1].PopMoves(toMove))
	}

	for _, s := range stacks {
		res += string(s.Pop())
	}

	return res
}
