package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/davidn5013/aoc/cast"
	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	rules := parseInput(input)
	_ = rules

	lowestloc := math.MaxInt
	ruleCnt := 6

	for _, seed := range rules.seeds {
		// seedstr := fmt.Sprint("Seed ", seed)
		for i := 0; i <= ruleCnt; i++ {
			// seedstr += fmt.Sprint(", ", strings.Split(rules.rulenames[i], "-")[2], " ")
			for _, v := range rules.rules[i] {
				dest := v[0]
				source := v[1]
				length := v[2]

				// fmt.Printf("%d seed %d (%t) %d\n", i, seed, seed >= source && seed <= source+length, seed-(source-dest))

				if seed >= source && seed <= source+length {
					seed = seed - (source - dest)
					break
				}
			}

			// seedstr += fmt.Sprint(seed)
		}
		// fmt.Println(seedstr)
		if seed < lowestloc {
			lowestloc = seed
		}
	}

	return lowestloc
}

func part2(input string) int {
	rules := parseInput(input)
	_ = rules

	lowestloc := math.MaxInt
	ruleCnt := 6

	// convert seeds to seed ranges
	for idx := 0; idx < len(rules.seeds)-1; idx += 2 {
		seedstart := rules.seeds[idx]
		seedend := rules.seeds[idx+1] + (rules.seeds[idx] - 1)

		// fmt.Printf("idx : \t %d %d\n", seedstart, seedend)

		for seedidx := seedstart; seedidx <= seedend; seedidx++ {
			seed := seedidx

			seedstr := fmt.Sprint("Seed ", seed)
			for i := 0; i <= ruleCnt; i++ {
				seedstr += fmt.Sprint(", ", strings.Split(rules.rulenames[i], "-")[2], " ")
				for _, v := range rules.rules[i] {
					dest := v[0]
					source := v[1]
					length := v[2]

					// fmt.Printf("%d seed %d (%t) %d\n", i, seed, seed >= source && seed <= source+length, seed-(source-dest))

					if seed >= source && seed <= source+length {
						seed = seed - (source - dest)
						break
					}
				}

				seedstr += fmt.Sprint(seed)
			}

			// fmt.Println(seedstr)

			if seed < lowestloc {
				lowestloc = seed
			}

		}

	}
	return lowestloc
}

type Rules struct {
	rulenames []string
	rules     map[int][][]int
	seeds     []int
}

func parseInput(input string) Rules {
	// parse seeds
	inputLine := strings.Split(input, "\n")

	var rules Rules
	rules.rules = make(map[int][][]int)
	for _, v := range strings.Fields(strings.Split(inputLine[0], ":")[1]) {
		rules.seeds = append(rules.seeds, cast.ToInt(v))
	}
	// fmt.Println("Seeds: ", rules.seeds)

	// parse unerule
	rulename := ""
	ruleorder := -1
	for _, v := range inputLine {

		if strings.Contains(v, "map:") {
			rulename = strings.Split(strings.Split(v, ":")[0], " ")[0]
			rules.rulenames = append(rules.rulenames, rulename)
			ruleorder++
		}

		if len(v) > 0 && len(rulename) > 0 && v[0] >= '0' && v[0] <= '9' { // rule
			r := []int{}
			for _, vv := range strings.Fields(v) {
				r = append(r, cast.ToInt(vv))
			}

			rules.rules[ruleorder] = append(rules.rules[ruleorder], r)
		}
	}

	// for i, v := range rules.rules {
	// 	fmt.Println(rules.rulenames[i], i, v)
	// }
	// fmt.Println("------------------------------------------------------------")

	return rules

}
