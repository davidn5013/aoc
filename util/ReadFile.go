package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadFile is a wrapper over io/ioutil.ReadFile but also determines the dynamic
// absolute path to the file.
//
// Deprecated in favor of go:embed, refer to scripts/skeleton/tmpls
func ReadFile(pathFromCaller string) string {
	// Docs: https://golang.org/pkg/runtime/#Caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	// read the entire file & return the byte slice as a string
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}
	// trim off new lines and tabs at end of input files
	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

// ShardInputFile read shared (input files in input catalog) and return as a string
func ShardInputFile(inputpath string) string {
	b, err := os.ReadFile(inputpath)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// PathInputShared get the path from input files in the shared input catalog
func PathInputShared(year, day, catalog, filename string) string {
	d := "0" + day
	f := year + "day" + d[len(d)-2:] + filename
	return filepath.Join(catalog, f)
}

// PathInputStandalone return filepath for inputfiles used by standalone aoc solutions (solution with main package)
func PathInputStandalone(year, day, filename string) string {
	d := "0" + day
	return filepath.Join(year, "day"+d[len(d)-2:], filename)
}

// Dirname is a port of __dirname in node
func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

// FileExists check if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// FileCopy copies one file
func FileCopy(sourcename, destinationname string) error {
	srcFile, err := os.Open(sourcename)
	if err != nil {
		return fmt.Errorf("ERR Missing input files %s ", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destinationname) // creates if file doesn't exist
	if err != nil {
		return fmt.Errorf("ERR cant opening destination file %s ", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return fmt.Errorf("ERR copying file %s ", err)
	}

	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("ERR ending copying file %s ", err)
	}

	return nil

}

// PreviewInput a long input string
func PreviewInput(input string, lines int) string {
	inputlines := strings.Split(input, "\n")
	var s strings.Builder
	l := len(inputlines)
	outlines := lines
	if lines > l {
		outlines = l
	}
	s.WriteString(fmt.Sprintln("First", outlines, "lines of total", l))
	for i, v := range inputlines {
		if i >= lines {
			break
		}
		s.WriteString(v + "\n")
	}
	return s.String()
}
