// David Nilsson 2022-08-01
// asdvent of code 2015 day 1
// 74
// 1795
// exmpel open smal file
// pause screen - down work with fresh

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// advent of code 2015 day 1
func adv2015d1(input string) { // floor,  onestring in file
	var f int
	for i, c := range input {
		s := string(c)
		switch s {
		case "(":
			f++
		case ")":
			f--
		default:
		}

		if f == -1 {
			fmt.Println(i+1, f)
		}
	}
	fmt.Println(f)
	// Pause()
}

func main() {
	adv2015d1(string(Glowfile("input")))
}

// Read the whole file to a array
func Glowfile(s string) []byte {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)

	}
	return content
}

//Standard error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Pause and wait for keypress
func Pause() {
	fmt.Println("press a key: ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
