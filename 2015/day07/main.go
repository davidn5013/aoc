package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/davidn5013/aoc/cast"
	"github.com/davidn5013/aoc/util"
)

/*
func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "filename of wire connetion")
	flag.Parse()

	strOfRules := map[string]string{}

	for _, inst := range strings.Split(util.ReadFile(filename), "\n") {
		parts := strings.Split(inst, " -> ")
		strOfRules[parts[1]] = parts[0]
	}

	printRules(strOfRules, "a", map[string]string{})
}
*/

func printRules(strOfRules map[string]string, needle string, mem map[string]string) string {
	// if it's a number, return the casted value
	if regexp.MustCompile("[0-9]").MatchString(needle) {
		return needle
	}

	if _, ok := strOfRules[needle]; !ok {
		log.Fatal("Wires does not exists : ", needle)
	}

	if val, ok := mem[needle]; ok {
		return val
	}

	rule := strOfRules[needle]
	parts := strings.Split(rule, " ")

	var result string
	switch {

	case len(parts) == 1:
		result = printRules(strOfRules, parts[0], mem)
	case parts[0] == "NOT":
		result = " NOT " + printRules(strOfRules, parts[1], mem)
	case parts[1] == "AND":
		result = printRules(strOfRules, parts[0], mem) + " AND " + printRules(strOfRules, parts[2], mem)
	case parts[1] == "OR":
		result = printRules(strOfRules, parts[0], mem) + " OR " + printRules(strOfRules, parts[2], mem)
	case parts[1] == "LSHIFT":
		result = printRules(strOfRules, parts[0], mem) + " LSHIFT " + printRules(strOfRules, parts[2], mem)
	case parts[1] == "RSHIFT":
		result = printRules(strOfRules, parts[0], mem) + " RSHIFT " + printRules(strOfRules, parts[2], mem)
	}

	fmt.Printf("%s=%s\n", needle, result)
	return result
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := someAssemblyRequired(util.ReadFile("input.txt"), part)
	fmt.Println("Output:", ans)
}

func someAssemblyRequired(input string, part int) int {
	wireToRule := map[string]string{}

	// generate graph of wires to their source rule
	for _, inst := range strings.Split(input, "\n") {
		parts := strings.Split(inst, " -> ")
		wireToRule[parts[1]] = parts[0]
	}

	aSignal := memoDFS(wireToRule, "a", map[string]int{})
	if part == 1 {
		printRules(wireToRule, "a", map[string]string{})
		return aSignal
	}

	// for part 2, override the value sent to wire b, then get output to a again
	wireToRule["b"] = cast.ToString(aSignal)
	printRules(wireToRule, "b", map[string]string{})
	return memoDFS(wireToRule, "a", map[string]int{})
}

func memoDFS(graph map[string]string, entry string, memo map[string]int) int {
	if memoVal, ok := memo[entry]; ok {
		return memoVal
	}

	// if it's a number, return the casted value
	if regexp.MustCompile("[0-9]").MatchString(entry) {
		return cast.ToInt(entry)
	}

	sourceRule := graph[entry]
	parts := strings.Split(sourceRule, " ")

	var result int
	switch {
	case len(parts) == 1:
		result = memoDFS(graph, parts[0], memo)
	case parts[0] == "NOT":
		start := memoDFS(graph, parts[1], memo)
		result = (math.MaxUint16) ^ start
	case parts[1] == "AND":
		result = memoDFS(graph, parts[0], memo) & memoDFS(graph, parts[2], memo)
	case parts[1] == "OR":
		result = memoDFS(graph, parts[0], memo) | memoDFS(graph, parts[2], memo)
	case parts[1] == "LSHIFT":
		result = memoDFS(graph, parts[0], memo) << memoDFS(graph, parts[2], memo)
	case parts[1] == "RSHIFT":
		result = memoDFS(graph, parts[0], memo) >> memoDFS(graph, parts[2], memo)
	}

	memo[entry] = result
	return result
}
