package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var input string

var logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	Level: slog.LevelWarn,
}))

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	/*
		logger.Debug("Debug message")
		logger.Info("Info message")
		logger.Warn("Warning message")
		logger.Error("Error message")
	*/

	var flDebug bool
	flag.BoolVar(&flDebug, "debug", false, "Turn on debug text ")

	var flPart int
	flag.IntVar(&flPart, "part", 2, "part 1 or 2")

	flag.Parse()

	if flDebug == true {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	fmt.Println("Running part", flPart)

	if flPart == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func parseInput(input string) (ans []string) {
	logger.Info("Parsing input")
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

// vowelcount true if contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou
func vowelcount(s string, amount int) bool {
	vowels := 0
	for _, char := range s {
		if strings.ContainsRune("aeiou", char) {
			vowels++
			if vowels >= amount {
				break
			}
		}
	}
	logger.Debug(fmt.Sprintf("vowelcount true on string: %s", s))
	return vowels >= amount
}

// dubbelLetters true if contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
func dubbelLetters(s string) (hasDouble bool) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			logger.Debug(fmt.Sprint("dubbelletters true on string:", s))
			return true
		}

	}
	return false
}

func forbiddenLetters(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, fw := range forbidden {
		if strings.Contains(s, fw) {
			logger.Info(fmt.Sprint("forbiddenLetters true on string:", s, " with ", fw))
			return true
		}

	}
	return false
}

func nice(s string) bool {
	return dubbelLetters(s) && !forbiddenLetters(s) && vowelcount(s, 3)
}

func part1(input string) int {
	parsed := parseInput(input)
	cnt := 0
	logger.Info("Parsing part1")
	for _, line := range parsed {
		if nice(line) {
			cnt++
		}
	}
	return cnt
}

func passesRule1(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		// ** Remeberb 0:2 2 is exclusive so "abc"[0:2]="ab" not "ac"
		toMatch := line[i : i+2]
		for j := i + 2; j < len(line)-1; j++ {
			if line[j:j+2] == toMatch {
				return true
			}
		}
	}
	return false
}

func passesRule2(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func part2(input string) int {
	nice := 0
	for _, line := range strings.Split(input, "\n") {
		if passesRule1(line) && passesRule2(line) {
			nice++
		}
	}

	return nice
}
