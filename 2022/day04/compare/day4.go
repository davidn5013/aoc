// aoc 2022 day 4
// https://raw.githubusercontent.com/mengkeat/AOC-2022/main/day04.go
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/davidn5013/aoc/utl"
)

//go:embed ..\input.txt
var inputFile string

func main() {
	start := utl.SetTimer()
	// bytes, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(inputFile)), "\n")

	count, count2 := 0, 0
	for _, s := range lines {
		tok := strings.Split(strings.ReplaceAll(s, "-", ","), ",")
		a1, _ := strconv.Atoi(tok[0])
		a2, _ := strconv.Atoi(tok[1])
		b1, _ := strconv.Atoi(tok[2])
		b2, _ := strconv.Atoi(tok[3])

		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			count++
		}
		if !(a2 < b1 || a1 > b2) {
			count2++
		}
	}
	fmt.Println("Part A:", count)
	fmt.Println("Part B:", count2)
	fmt.Println("took  ", start()*time.Millisecond)
}
