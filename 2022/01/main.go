// Package main https://adventofcode.com/2022/day/1
package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

// Solution part 1
func sumStringArrOnEmptyLineBigest(numberString []string) int {
	var (
		bigCalorieElf, sumCalcories = 0, 0
	)
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

func sumStringOnLineReturnSortSum(numberString []string) []int {
	var (
		allCalcoriseSum []int
		sumCalcories    = 0
	)
	for _, line := range numberString {
		calcorie, err := strconv.Atoi(line)
		if err != nil {
			allCalcoriseSum = append(allCalcoriseSum, sumCalcories)
			sumCalcories = 0
		} else {
			sumCalcories += calcorie
		}
	}

	// Sort on bigest first
	sort.Slice(allCalcoriseSum, func(i, j int) bool {
		return allCalcoriseSum[i] > allCalcoriseSum[j]
	})

	return allCalcoriseSum
}
func main() {
	// find empty lines number of elfs
	var (
		caloriesSubLists = strings.Split(inputFile, "\n")
	)

	fmt.Println(sumStringArrOnEmptyLineBigest(caloriesSubLists))
	sortCalories := sumStringOnLineReturnSortSum(caloriesSubLists)

	sumOfTreeHigest := sortCalories[0] + sortCalories[1] + sortCalories[2]
	fmt.Println(sumOfTreeHigest)
}
