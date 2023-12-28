package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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

// card number A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

// Every hand is exactly one type. From strongest to weakest, they are:
//
//     Five of a kind, where all five cards have the same label: AAAAA
//     Four of a kind, where four cards have the same label and one card has a different label: AA8AA
//     Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
//     Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
//     Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
//     One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
//     High card, where all cards' labels are distinct: 23456

func part1(input string) int {
	cards := parseInput(input)

	for _, card := range cards {
		fmt.Println(card.card, len(cardcnt(card.card, 0)), cardtype(card.card))
	}

	fmt.Println(cards[0].card, cards[1].card, firstCardBig(cards[0].card, cards[1].card))
	return 0
}

func cardtype(card string) string {
	switch {
	case len(cardcnt(card, 5)) >= 1:
		return "Five of a kind"
	case len(cardcnt(card, 4)) >= 1:
		return "Four of a kind"
	case len(cardcnt(card, 3)) >= 1 && len(cardcnt(card, 2)) >= 1:
		return "Full house"
	case len(cardcnt(card, 3)) >= 1:
		return "Three of a kind"
	case twopar(card):
		return "Two par"
	case len(cardcnt(card, 2)) >= 1:
		return "Par"
	}
	return "High card"
}

func firstCardBig(card1, card2 string) bool {
	cardval := []string{" ", "A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	if len(card1) != len(card2) {
		log.Fatal("ERR differen size of cards ", len(card1), len(card2))
	}

	var f, s int
	for i := 0; i < len(card1); i++ {
		for j := 0; j < len(cardval); j++ {
			if string(card1[i]) == cardval[j] {
				f = j
			}
			if string(card2[i]) == cardval[j] {
				s = j
			}

		}
		if f < s {
			fmt.Println(string(card1[i]), string(card2[i]), f, s)
			return true
		}
	}
	return false
}

func twopar(card string) bool {
	m := cardcnt(card, 2)
	ans := make(map[string]int)
	for i, v := range m {
		if v == 2 {
			ans[i] = v
		}
	}
	return len(ans) == 2
}

func cardcnt(card string, filter int) (ans map[string]int) {
	t := make(map[string]int, len(card))
	for i := 0; i < len(card); i++ {
		t[string(card[i])]++
	}
	ans = make(map[string]int, len(t))
	for i, v := range t {
		if v == filter || filter == 0 {
			ans[string(i)] = v
		}
	}
	return ans
}

func part2(input string) int {
	return 0
}

type Cards struct {
	sortvalue int
	card      string
	rankvalue int
}

func parseInput(input string) (ans []Cards) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		ans = append(ans, Cards{
			card:      parts[0],
			rankvalue: cast.ToInt(parts[1]),
		})
	}
	return ans
}
