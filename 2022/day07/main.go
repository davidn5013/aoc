// Package main https://adventofcode.com/2022/day/
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/utl"
)

// node catalog struct for storing file of catalog information
type node struct {
	name   string
	size   int
	isFile bool
	sons   map[string]*node
	father *node
}

// More debug text 0-10 : 10 Max
const debuglvl = 0

// main Run Solution for Advent of Code
func main() {
	fmt.Printf("Solution 1 (test case)\t= %d\n", sol1("input_test.txt"))
	stopTimer := utl.SetTimer()
	fmt.Printf("Solution 1\t= %d\ttook %s\n", sol1("input.txt"), stopTimer())

	fmt.Printf("Solution 2 (test case)\t= %d\n", sol2("input_test.txt"))
	stopTimer = utl.SetTimer()
	fmt.Printf("Solution 2\t= %d\ttook %s\n", sol2("input.txt"), stopTimer())
}

// sol1 Solution part 1
func sol1(inputFile string) (ret int) {
	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	input, _ := os.Open(inputFile)
	// Closing with check for closing failer
	defer func() {
		err := input.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	// open file split on ScanLines by default
	sc := bufio.NewScanner(input)

	var (
		currentDirectory *node
		dirs             = []*node{}
		totalSize        int
	)

	/*
		* Parsing input file like this:
				* $ cd /
				* $ ls
				* dir a
				* 14848514 b.txt
				* 8504156 c.dat
				* dir d
				* $ cd a
				* $ ls
				* dir e
				* ...
	*/

	// fail in here -- it was extra newline or something. direct download for file worked
	for sc.Scan() {
		line := strings.Fields(sc.Text())
		if len(line) > 2 {
			if line[2] == ".." {
				currentDirectory = currentDirectory.father
			} else if line[2] == "/" {
				currentDirectory = &node{"/", 0, false, make(map[string]*node), nil}
				dirs = append(dirs, currentDirectory)
			} else {
				currentDirectory = currentDirectory.sons[line[2]]
			}
		} else if line[0] == "dir" {
			currentDirectory.sons[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
			dirs = append(dirs, currentDirectory.sons[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			if err := sc.Err(); err != nil {
				err = fmt.Errorf("convert to number error line after dir not number %s %s", utl.CurrFuncName(), err)
				log.Fatal(err)
			}
			currentDirectory.sons[line[1]] = &node{line[1], size, true, nil, currentDirectory}
		}
	}

	if err := sc.Err(); err != nil {
		err = fmt.Errorf("error closing scanner in %s %s,", utl.CurrFuncName(), err)
		log.Fatal(err)
	}

	for _, dir := range dirs {
		size := calcSize(*dir)
		if size <= 100000 {
			totalSize += size
		}
	}

	// fmt.Println(totalSize)

	/*
		* Only for part two

			 toFree:= 30000000 - (70000000 - calcSize(*dirs[0]))
			 var smallestEnaugthSize int = calcSize(*dirs[0])
				for _, dir := range dirs {
					size := calcSize(*dir)
					if size > toFree && size-toFree < smallestEnaugthSize-toFree {
						smallestEnaugthSize = size
					}
				}

				fmt.Println(smallestEnaugthSize)
	*/

	return totalSize
}

// sol2 Solution part 2
func sol2(inputFile string) (ret int) {
	utl.Debug(debuglvl >= 10, "Starting %s\n", utl.CurrFuncName())

	input, _ := os.Open(inputFile)
	// Closing with check for closing failer
	defer func() {
		err := input.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	// open file split on ScanLines by default
	sc := bufio.NewScanner(input)

	var (
		currentDirectory    *node
		dirs                = []*node{}
		smallestEnaugthSize int
	)

	/*
		* Parsing input file like this:
				* $ cd /
				* $ ls
				* dir a
				* 14848514 b.txt
				* 8504156 c.dat
				* dir d
				* $ cd a
				* $ ls
				* dir e
				* ...
	*/

	// fail in here -- it was extra newline or something. direct download for file worked
	for sc.Scan() {
		line := strings.Fields(sc.Text())
		if len(line) > 2 {
			if line[2] == ".." {
				currentDirectory = currentDirectory.father
			} else if line[2] == "/" {
				currentDirectory = &node{"/", 0, false, make(map[string]*node), nil}
				dirs = append(dirs, currentDirectory)
			} else {
				currentDirectory = currentDirectory.sons[line[2]]
			}
		} else if line[0] == "dir" {
			currentDirectory.sons[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
			dirs = append(dirs, currentDirectory.sons[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			currentDirectory.sons[line[1]] = &node{line[1], size, true, nil, currentDirectory}
		}
	}

	// if err := sc.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	/*
		* Part one
			for _, dir := range dirs {
				size := calcSize(*dir)
				if size <= 100000 {
					totalSize += size
				}
			}
	*/

	// fmt.Println(totalSize)

	/* Only for part two*/

	toFree := 30000000 - (70000000 - calcSize(*dirs[0]))
	smallestEnaugthSize = calcSize(*dirs[0])
	for _, dir := range dirs {
		size := calcSize(*dir)
		if size > toFree && size-toFree < smallestEnaugthSize-toFree {
			smallestEnaugthSize = size
		}
	}

	return smallestEnaugthSize
}

/*
* Helper function - method for type node
 */
// calcSize recursive sum of size in node
func calcSize(root node) (size int) {
	if root.isFile {
		return root.size
	}
	for _, d := range root.sons {
		size += calcSize(*d)
	}
	return
}
