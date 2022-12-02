// Package main https://adventofcode.com/2022/day/1
// From file with one number per line sum after empty line,
// This calories in food carried per elf
// First what elf carries the most calories?
// What the sum of three biggest of sums?
package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// -- I find it easiest just to embed this one
//
//go:embed input.txt
var inputFile string

// Solution part 1 - What Elf carries the most food calories?
func elfMosteFood(numberString []string) int {
	var (
		bigCalorieElf, sumCalcories = 0, 0
	)

	// Every line that can be convert to number
	// is prosumed to be the empty line for new elf
	for _, line := range numberString {
		calcorie, err := strconv.Atoi(line)
		if err != nil {
			if sumCalcories > bigCalorieElf {
				bigCalorieElf = sumCalcories
			}
			sumCalcories = 0
		} else {
			sumCalcories += calcorie
		}
	}
	return bigCalorieElf
}

// Solution 2 Sort all the sum of calories per elf
func caloriePerElfSort(numberString []string) []int {
	var (
		totCalcorSum []int
		sumCalcories = 0
	)

	// Every line that can be convert to number
	// is prosumed to be the empty line for new elf
	for _, line := range numberString {
		calcorie, err := strconv.Atoi(line)
		if err != nil {
			totCalcorSum = append(totCalcorSum, sumCalcories)
			sumCalcories = 0
		} else {
			sumCalcories += calcorie
		}
	}

	// Sort on biggest first
	sort.Slice(totCalcorSum, func(i, j int) bool {
		return totCalcorSum[i] > totCalcorSum[j]
	})

	return totCalcorSum
}

func main() {
	var (
		oneCellPerline = strings.Split(inputFile, "\n")
	)

	desc2022day1()

	sortedCalPerElf := caloriePerElfSort(oneCellPerline)

	sumOfTreeHigest := sortedCalPerElf[0] + sortedCalPerElf[1] + sortedCalPerElf[2]
	fmt.Printf("Elf carrying moste food two %d , %d\n", elfMosteFood(oneCellPerline), sortedCalPerElf[0])
	fmt.Printf("Three Elf with moste food totalte have %d\n", sumOfTreeHigest)
}

func desc2022day1() {
	fmt.Println(`--- Day 1: Calorie Counting ---

Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas. For that, their favorite snack is a special type of star fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.

To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

This list represents the Calories of the food carried by five Elves:

    The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
    The second Elf is carrying one food item with 4000 Calories.
    The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
    The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
    The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

Your puzzle answer was 69626.
--- Part Two ---

By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories. That way, even if one of those Elves runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories). The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?

Your puzzle answer was 206780.`)

}
