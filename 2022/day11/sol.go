package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davidn5013/aoc/utl"
)

// Aoc is a place for the main datastruct in this solution
type Aoc struct {
	monkeys []monkey
	round   int
}

type monkey struct {
	monkeyId   int
	itemWValue []int
	operType   rune
	operValue  string
	testValue  int
	toMonkey   []int
}

// NewAoc Aoc struct for cpu information
func NewAoc() *Aoc {
	a := Aoc{}
	return &a
}

func (a Aoc) String() (ret string) {
	ret = fmt.Sprintf("After round %d, the monkeys are holdning items with worry levels:\n", a.round)
	for idx, monkey := range a.monkeys {
		s := ""
		for _, itemV := range monkey.itemWValue {
			t := strconv.Itoa(itemV)
			s += t + ", "
		}
		ret += fmt.Sprintf("Monkey %d: %s\n", idx, s)
	}
	return ret
}

func (m monkey) String() (ret string) {
	ret = fmt.Sprintf("Monkey %d:\n", m.monkeyId)
	ret += fmt.Sprintf("  Starting items: %v\n", m.itemWValue)
	ret += fmt.Sprintf("  Operation codes: %c %s\n", m.operType, m.operValue)
	ret += fmt.Sprintf("  Test codes: %d\n", m.testValue)
	ret += fmt.Sprintf("  throw list: %v\n", m.toMonkey)
	return ret
}

func (a *Aoc) parse(filename string) {
	startmsg(utl.CurrFuncName())

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	sc := bufio.NewScanner(fp)

	for sc.Scan() {
		// monkeys id
		m := monkey{}
		fmt.Sscanf(sc.Text(), "Monkey %d:\n", &m.monkeyId)

		// items
		sc.Scan()
		for _, item := range strings.Split(sc.Text()[len("  Starting items:"):], ", ") {
			i, err := strconv.Atoi(strings.TrimSpace(item))
			if err != nil {
				log.Fatal("Failed to convert items wores values")
			}
			m.itemWValue = append(m.itemWValue, i)
		}

		// Operation
		sc.Scan()
		_, err := fmt.Sscanf(sc.Text(), "  Operation: new = old %c %s", &m.operType, &m.operValue)
		if err != nil {
			log.Fatal("Failed to read operation")
		}

		// Test
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &m.testValue)
		if err != nil {
			log.Fatal("Failed to read test div")
		}

		// throw if true
		sc.Scan()
		var m1 int
		_, err = fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &m1)
		if err != nil {
			log.Fatal("Failed to read monkey on true")
		}

		// monkey throw to if false or true
		sc.Scan()
		var m2 int
		_, err = fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &m2)
		if err != nil {
			log.Fatal("Failed to read monkey on false")
		}

		// store throw
		m.toMonkey = append(m.toMonkey, m1, m2)

		// skip empty line
		sc.Scan()

		// Fine store new monkey in Aoc struct
		a.monkeys = append(a.monkeys, m)

		utl.Debug(debuglvl >= 10, "%s", m) // Debug show every monkey like input
	}
	utl.Debug(debuglvl >= 5, "%s", a) // Summarize of monkey input
}

func (a *Aoc) sol1() (ret int) {
	startmsg(utl.CurrFuncName())

	// did we get it all and are monkey id same as index om monkeys?
	// Yes we have it all and index of monkeys array and monkey id is the same we don't
	// really need monkey id but whats the harm. That going to be a refactoring.
	// for idx, monkey := range a.monkeys {
	// 	fmt.Printf(" Monkey index %d has id %d\n", idx, monkey.monkeyId)
	// }

	// rounds 20
	for a.round = 1; a.round <= 1; a.round++ {
		for monkId, monkey := range a.monkeys {

			// update worry values add extra and dive by 3
			for ItemId, itemv := range monkey.itemWValue {
				var multiValue int
				if monkey.operValue == "old" {
					multiValue = itemv
				} else {
					var err error
					multiValue, err = strconv.Atoi(monkey.operValue)
					if err != nil {
						log.Fatal("Failed to convert operValue in sol1 calculation")
					}
				}

				if monkey.operType == '+' {
					a.monkeys[monkId].itemWValue[ItemId] += multiValue
					a.monkeys[monkId].itemWValue[ItemId] = a.monkeys[monkId].itemWValue[ItemId] / 3
				} else if monkey.operType == '*' {
					a.monkeys[monkId].itemWValue[ItemId] *= multiValue
					a.monkeys[monkId].itemWValue[ItemId] = a.monkeys[monkId].itemWValue[ItemId] / 3
				} else {
					log.Fatal("Failed to calculate item worry level")
				}
				// test worry value:167

				temp := monkey.itemWValue
				for iidx, itemv := range temp {
					if itemv%monkey.testValue == 0 {
						a.throwToMonkey(monkId, monkey.toMonkey[0], iidx)
					} else {
						a.throwToMonkey(monkId, monkey.toMonkey[1], iidx)
					}

				}

				a.monkeys[monkId].itemWValue = []int{}

				break
			}
		}
	}
	fmt.Printf("%s", a)

	return ret
}

func (a Aoc) throwToMonkey(fromMonkeyId, tomMonkeyId, itemId int) {
	utl.Debug(debuglvl >= 10, "throwToMonkey %d %d %d\n", fromMonkeyId, tomMonkeyId, itemId)

	// append to tom Monkey
	a.monkeys[tomMonkeyId].itemWValue = append(a.monkeys[tomMonkeyId].itemWValue, a.monkeys[fromMonkeyId].itemWValue[itemId])
}

func (a *Aoc) sol2() (ret int) {
	startmsg(utl.CurrFuncName())
	return ret
}
