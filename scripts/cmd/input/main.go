package main

import "github.com/davidn5013/aoc/scripts/aoc"

func main() {
	day, year, cookie := aoc.ParseFlags()
	aoc.GetInput(day, year, cookie)
}
