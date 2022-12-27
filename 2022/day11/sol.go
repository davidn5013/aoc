// Sol 2 day 11
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

// aoc is a place for the main datastruct in this solution
type aoc struct {
	monkeylist []monkey // monkeys array
	// solution variables
	round        int // number of rounds}
	bigLimitSol2 int
}

// newAoc Aoc struct for cpu information
func newAoc() *aoc {
	a := aoc{}
	return &a
}

type monkey struct { // Monkey
	items []int               // worry value for every item that the monkey have
	op    func(int, bool) int // func for changing value of items for this monkey
	throw func(int) int       // func for throw to a new monkey for this monkey
}

func (m *monkey) newOp(operator rune, value int) func(int, bool) int {
	operation := func(n int, sol2 bool) int {
		var valueToUse int = value
		if valueToUse == 0 {
			valueToUse = n
		}
		if operator == '+' {
			if sol2 == true {
				return (n + valueToUse)
			} else {
				return (n + valueToUse) / 3
			}
		}
		if sol2 == true {
			return (n * valueToUse)
		} else {
			return (n * valueToUse) / 3
		}
	}
	return operation
}

func (m *monkey) newThrow(testingValue, toThrowIfTrue, toThrowIfFalse int) func(int) int {

	testAndThrow := func(n int) int {
		if n%testingValue == 0 {
			return toThrowIfTrue
		}
		return toThrowIfFalse
	}
	return testAndThrow
}

func (a *aoc) parse(filename string) {
	startmsg(utl.CurrFuncName())
	a.bigLimitSol2 = 1

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
		var id int
		// monkeys id is set in array index Aoc.m
		m := monkey{}

		// items
		sc.Scan()
		for _, item := range strings.Split(sc.Text()[len("  Starting items:"):], ", ") {
			i, _ := strconv.Atoi(strings.TrimSpace(item))
			m.items = append(m.items, i)
			utl.Debug(debuglvl >= 10, "Monkey id %d new item %d\n", id, i)
		}

		// Operation
		sc.Scan()
		var op rune
		var opValue int
		if strings.Contains(sc.Text(), "old * old") || strings.Contains(sc.Text(), "old + old") {
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c old", &op)
		} else {
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c %d", &op, &opValue)
		}
		m.op = m.newOp(op, opValue)
		utl.Debug(debuglvl >= 10, "Monkey id %d op %c %d\n", id, op, opValue)

		// Test and throw
		sc.Scan()
		var tValue int
		fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &tValue)
		a.bigLimitSol2 *= tValue

		sc.Scan()
		var toThrowIfTrue int
		fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &toThrowIfTrue)

		sc.Scan()
		var toThrowIfFalse int
		fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &toThrowIfFalse)
		m.throw = m.newThrow(tValue, toThrowIfTrue, toThrowIfFalse)

		utl.Debug(debuglvl >= 10, "Monkey %d test value %d throw %d %d\n",
			id, tValue, toThrowIfTrue, toThrowIfFalse)

		// skip empty line
		sc.Scan()

		// Fine store new monkey in Aoc struct
		a.monkeylist = append(a.monkeylist, m)
		utl.Debug(debuglvl >= 10, "New monkey list %v\n", m)
		id++
	}
}

func (a *aoc) sol1(filename string) (ret int) {
	startmsg(utl.CurrFuncName())
	a.parse(filename)
	counts := make([]int, len(a.monkeylist))

	// rounds 20
	for a.round = 1; a.round <= 20; a.round++ {
		for mId, m := range a.monkeylist {

			// update worry values add extra and dive by 3
			for _, item := range m.items {
				nV := m.op(item, false)
				a.monkeylist[m.throw(nV)].items =
					append(a.monkeylist[m.throw(nV)].items, nV)
			}
			counts[mId] += len(a.monkeylist[mId].items)
			a.monkeylist[mId].items = []int{}
		}
		utl.Debug(debuglvl >= 5, "%#v\n", counts)

		var highestCount, secondHighest int
		for _, count := range counts {
			if count > secondHighest {
				secondHighest = count
			}
			if secondHighest > highestCount {
				highestCount, secondHighest =
					secondHighest, highestCount
			}
		}
		ret = highestCount * secondHighest
	}
	return ret
}

func (a *aoc) sol2(filename string) (ret int) {
	startmsg(utl.CurrFuncName())
	a.parse(filename)
	counts := make([]int, len(a.monkeylist))

	// rounds 10_000
	for a.round = 1; a.round <= 10_000; a.round++ {
		for mId, m := range a.monkeylist {
			// update worry values add extra and dive by 3
			for _, item := range m.items {
				nV := (m.op(item, true) % a.bigLimitSol2)

				a.monkeylist[m.throw(nV)].items =
					append(a.monkeylist[m.throw(nV)].items, nV)

			}
			counts[mId] += len(a.monkeylist[mId].items)
			a.monkeylist[mId].items = []int{}
		}
		utl.Debug(debuglvl >= 5, "%#v\n", counts)
	}

	var highestCount, secondHighest int
	for _, count := range counts {
		if count > secondHighest {
			secondHighest = count
		}
		if secondHighest > highestCount {
			highestCount, secondHighest =
				secondHighest, highestCount
		}
	}
	utl.Debug(debuglvl >= 3, "Two moste active monkeys %d,%d\n", secondHighest, highestCount)
	ret = highestCount * secondHighest
	return ret
}
