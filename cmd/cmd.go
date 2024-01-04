// Package cmd runnes aoc solutions
package cmd

// TODO: Run advent of code solution part 1 & 2
// TODO: Download input for solutions
// TODO: Create template for new solution

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidn5013/aoc/cast"
	"github.com/davidn5013/aoc/scripts/aoc"
	"github.com/davidn5013/aoc/util"
)

type flagType struct {
	Year   string
	Day    string
	Cookie string
}

func flags() (f flagType) {
	flag.StringVar(&f.Year, "year", "2023", "AOC Year")
	flag.StringVar(&f.Day, "day", "7", "AOC Day")
	flag.StringVar(&f.Cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.Parse()
	return f
}

const inputfilename = "input.txt"
const inputpath = "input"

// Execute is the main for running aoc
func Execute() {
	f := flags()

	for _, v := range sols {
		if v.year == f.Year && v.day == f.Day {
			RunAoc(f.Year, f.Day, inputfilename, f.Cookie)
			os.Exit(1) // Not exit here
		}
	}

	// no solution in sol catalog run standalone
	inputfileCreate(f.Year, f.Day, inputfilename, f.Cookie)
	// printMainSolutions(".")

	a := createAocInputPath(f.Year, f.Day, "main.go")
	if util.FileExists(a) {
		util.Run("go", "run", a)
	} else {
		fmt.Println(a, "does not exists")
	}

}

// TODO no nead for this function
// printMainSolutions list aoc solution that standalone main files
func printMainSolutions(path string) {
	fileSystem := os.DirFS(path)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if len(path) > 2 && path[:2] == "20" && strings.Contains(path, "day") && strings.Contains(path, "main.go") {
			fmt.Println(path)
		}
		return nil
	})

}

// TODO move check for input catalog so aoc sol catalog can check for input

// inputfileCreate get inputfiles for stand alone aoc solutions, download and copy if missing
func inputfileCreate(year, day, filename, cookie string) {
	inpInFold := createAocInputPath(year, day, filename)
	if !util.FileExists(inpInFold) {

		src := inputPath(year, day, inputpath, inputfilename)
		dst := createAocInputPath(year, day, inputfilename)

		if !util.FileExists(src) && cookie != "" {
			aoc.GetInput(cast.ToInt(day), cast.ToInt(year), cookie)
		} else {
			log.Fatal("ERR Missing aoc cookie")
		}

		if !util.FileExists(dst) {
			err := util.FileCopy(src, dst)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}

// createAocInputPath return path and file for standalone aoc solutions
func createAocInputPath(year, day, filename string) string {
	d := "0" + day
	return filepath.Join(year, "day"+d[len(d)-2:], filename)
}
