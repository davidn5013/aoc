// Package cmd runnes aoc solutions
package cmd

// TODO: Run advent of code solution part 1 & 2
// TODO: Download input for solutions
// TODO: Create template for new solution

import (
	"flag"
	"fmt"
	"io"
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

func Execute() {
	f := flags()

	for _, v := range sols {
		if v.year == f.Year && v.day == f.Day {
			RunAoc(f.Year, f.Day, inputfilename, f.Cookie)
			os.Exit(1)
		}
	}

	inpfile(f.Year, f.Day, inputfilename, f.Cookie)
	// printMainSolutions(".")

	a := createAocInputPath(f.Year, f.Day, "main.go")
	if util.FileExists(a) {
		util.Run("go", "run", a)
	} else {
		fmt.Println(a, "does not exists")
	}

}

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

func inpfile(year, day, filename, cookie string) {
	inp := createAocInputPath(year, day, filename)
	fmt.Println(inp)
	if !util.FileExists(inp) {
		fmt.Println("Missing", filename, "file")
		if cookie != "" {
			aoc.GetInput(cast.ToInt(day), cast.ToInt(year), cookie)

			sourcename := inputfile(year, day, inputpath, filename)
			srcFile, err := os.Open(sourcename)
			if err != nil {
				log.Fatal("Cant open input file for copy", err)
			}
			defer srcFile.Close()

			destname := createAocInputPath(year, day, filename)
			destFile, err := os.Create(destname) // creates if file doesn't exist
			if err != nil {
				log.Fatal("ERR Opening destination file", err)
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
			if err != nil {
				log.Fatal("ERR Coping file", err)
			}

			err = destFile.Sync()
			if err != nil {
				log.Fatal("ERR Sync files", err)
			}
		}
	}

}

func createAocInputPath(year, day, filename string) string {
	d := "0" + day
	return filepath.Join(year, "day"+d[len(d)-2:], filename)
}
