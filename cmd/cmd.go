// Package cmd runnes aoc solutions
package cmd

// TODO: Create template for new solution

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/davidn5013/aoc/cast"
	"github.com/davidn5013/aoc/scripts/aoc"
	"github.com/davidn5013/aoc/util"
)

type flagType struct {
	Year     string
	Day      string
	Cookie   string
	DebugLvl int
}

func flags() (f flagType) {
	flag.StringVar(&f.Year, "year", "2023", "AOC Year")
	flag.StringVar(&f.Day, "day", "7", "AOC Day")
	flag.StringVar(&f.Cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.IntVar(&f.DebugLvl, "debuglvl", 0, "Set debug level 0=none, 5=info, 10=info&debug")
	flag.Parse()
	util.SetDebuglvl(f.DebugLvl)
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
			os.Exit(0) // Note exit here
		}
	}

	// no solution in sol catalog run standalone
	GetCreateStandAloneInput(f.Year, f.Day, inputfilename, f.Cookie)
	// printMainSolutions(".")

	input := util.PathInputStandalone(f.Year, f.Day, "main.go")
	if util.FileExists(input) {
		util.Run("go", "run", input)
	} else {
		fmt.Println(input, "does not exists")
	}

}

// RunAoc execute aoc solutions in sol catalog
func RunAoc(year, day, filename, cookie string) {
	inputfile := util.PathInputShared(year, day, inputpath, filename)

	if !util.FileExists(inputfile) {
		if cookie != "" {
			aoc.GetInput(cast.ToInt(day), cast.ToInt(year), cookie)
		} else {
			log.Fatal("Need aoc cookie for download of missing inputfile")
		}
	}

	for _, s := range sols {
		if s.year == year && s.day == day {
			input := util.ShardInputFile(inputfile)
			fmt.Printf("Part1: %d\nPart2: %d\n", s.part1(input), s.part2(input))
		}
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

// GetCreateStandAloneInput get inputfiles for stand alone aoc solutions (solution with main package), download and copy if missing
func GetCreateStandAloneInput(year, day, filename, cookie string) {
	inpInFold := util.PathInputStandalone(year, day, filename)
	if !util.FileExists(inpInFold) {

		src := util.PathInputShared(year, day, inputpath, inputfilename)
		dst := util.PathInputStandalone(year, day, inputfilename)

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
