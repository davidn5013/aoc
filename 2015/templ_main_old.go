// David Nilsson 2022 - asdvent of code
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// advent of code 20XX day X
func adv(input string) { // floor,  onestring in file
	var f int
	for i, c := range input {
		s := string(c)
	}
	Pause()
}

func main() {
	adv(string(Glowfile("input")))
}

// Read the whole file to a array
func Glowfile(s string) []byte {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)

	}
	return content
}

// Pause and wait for keypress
func Pause() {
	fmt.Println("press a key: ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
