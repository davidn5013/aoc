/*
 * Solution bye https://github.com/lynerist/
 *
 */

// Package main https://adventofcode.com/2022/day/6
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/davidn5013/aoc/utl"
)

// More debug text 0-10 : 10 Max
const debuglvl = 0

// main Run Solution for Advent of Code
func main() {
	fmt.Printf("Solution 1 = %d\n", sol1("input_test.txt"))
	fmt.Printf("Solution 1 = %d\n", sol1("input.txt"))
	fmt.Printf("Solution 1 = %d\n", sol2("input_test.txt"))
	fmt.Printf("Solution 1 = %d\n", sol2("input.txt"))
}

// sol1 Solution part 1
func sol1(inputFile string) (ret int) {
	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	input, _ := os.Open(inputFile)
	defer func() {
		err := input.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	sc := bufio.NewScanner(input)
	sc.Scan()
	for i := range sc.Text() {
		isUniq := make(map[byte]bool)

		for j := 0; j < 4; j++ {
			/* Fill a hashtabell byte */
			isUniq[sc.Text()[i+j]] = true
			utl.Debug(debuglvl >= 10, "Setting isUniq map to %#v\n", isUniq)
		}
		if len(isUniq) == 4 {
			ret = i + 4
			break
		}
	}
	return ret
}

// sol2
func sol2(inputFile string) (ret int) {
	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	input, _ := os.Open(inputFile)
	defer func() {
		err := input.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	sc := bufio.NewScanner(input)
	sc.Scan()
	for i := range sc.Text() {
		isUniq := make(map[byte]bool)

		for j := 0; j < 14; j++ {
			isUniq[sc.Text()[i+j]] = true
			utl.Debug(debuglvl >= 10, "Setting isUniq map to %#v\n", isUniq)
		}
		if len(isUniq) == 14 {
			ret = i + 14
			break
		}
	}
	return ret
}
