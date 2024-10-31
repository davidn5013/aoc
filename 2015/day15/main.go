package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

// go:embed inp.txt
//
//disabled_ go:embed input.txt
// var input string

// INFO print all [INFO] lines
const INFO = false

var input = "3113322113"
var iterations = 40

func main() {
	fmt.Println("input =", input)
	for i := 0; i < iterations; i++ {
		input = lookAndSay(input)
	}
	fmt.Println(len(input))
	//fmt.Println(input)

}

func lookAndSay(in string) (out string) {
	left := 0
	right := 0
	digcnt := 0
	for {
		if right >= len(in) {
			out = out + string(strconv.Itoa(digcnt)) + string(in[left])
			break
		}

		if in[left] == in[right] {
			digcnt++
			right++
		} else {
			out = out + string(strconv.Itoa(digcnt)) + string(in[left])
			left = right
			digcnt = 0
		}

	}

	return out
}
