package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/mathy"
	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var realInput string

//go:embed test.txt
var testInput string

// global variable switch on extra prints
var flgVerbose bool

func main() {

	var flgPart int
	var flgFile bool

	flag.IntVar(&flgPart, 'part', 1, "part 1 or 2")
	flag.BoolVar(&flgFile, "td", false, "Test Data - Use test input instead for real")
	flag.BoolVar(&flgVerbose, "v", false, "Verbose output")
	flag.Parse()

	var input *string
	if !flgFile {
		input = &realInput
	} else {
		input = &testInput
	}

	shortest, longest := travelingSalesman(*input)

	switch flgPart {
	case 1:
		util.CopyToClipboard(fmt.Sprintf("%v", shortest))
		fmt.Println("Output part 1:", shortest)
	case 2:
		util.CopyToClipboard(fmt.Sprintf("%v", longest))
		fmt.Println("Output part 2:", longest)
	default:
		fmt.Println("Missing part")
		os.Exit(-1)
	}
}

type mapCitys map[string]map[string]int

func travelingSalesman(input string) (int, int) {
	graph := createCitys(input)

	short := math.MaxInt32
	long := 0
	for k := range graph {
		if flgVerbose {
			fmt.Println("> Running city: ", k)
		}
		dfsMin, dfsMax := dfsTotalDistance(graph, k, map[string]bool{k: true})
		short = min(short, dfsMin)
		long = max(long, dfsMax)
	}

	return short, long
}

func createCitys(input string) (citys mapCitys) {
	citys = make(mapCitys)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}

		var start, end, sVal string
		fmt.Sscanf(line, "%s to %s = %s", &start, &end, &sVal)

		val, err := strconv.Atoi(sVal)
		if err != nil {
			log.Printf("ERR read value from %s to %s , value %s", start, end, sVal)
			continue
		}

		if citys[start] == nil {
			citys[start] = make(map[string]int)
		}
		if citys[end] == nil {
			citys[end] = make(map[string]int)
		}

		citys[start][end] = val
		citys[end][start] = val
	}

	return citys
}

func dfsTotalDistance(graph mapCitys, entry string, visited map[string]bool) (min, max int) {
	// if all nodes have been visited, return a zero length
	if len(visited) == len(graph) {
		return 0, 0
	}

	minDistance := math.MaxInt32
	maxDistance := 0

	for k := range graph {
		if !visited[k] {
			visited[k] = true

			weight := graph[entry][k]
			minRecurse, maxRecurse := dfsTotalDistance(graph, k, visited)
			minDistance = mathy.MinInt(minDistance, weight+minRecurse)
			maxDistance = mathy.MaxInt(maxDistance, weight+maxRecurse)
			if flgVerbose {
				fmt.Printf("\t%s\t%d\t%d %d\n", k, weight, minDistance, maxDistance)
			}

			// backtrack
			// delete to so length of visited is accurate
			delete(visited, k)
		}
	}

	return minDistance, maxDistance

}
