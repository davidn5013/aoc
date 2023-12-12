package main

import (
	"strings"
	"testing"

	"github.com/davidn5013/aoc/util"
)

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_hasSymbol(t *testing.T) {
	util.SetDebuglvl(10)
	var parsed [][]byte

	for _, row := range strings.Split(example, "\n") {
		var temp []byte
		for _, b := range row {
			temp = append(temp, byte(b))
		}
		parsed = append(parsed, temp)
	}

	tests := []struct {
		name string
		col  int
		num  number
		want bool
	}{
		{
			name: "Testing hasSymbol one first 35",
			col:  2,
			num:  number{35, 2, 4},
			want: true,
		},
		{
			name: "Testing hasSymbol one first 114",
			col:  0,
			num:  number{114, 5, 8},
			want: false,
		},
		{
			name: "Testing hasSymbol out of boucne",
			col:  0,
			num:  number{9999999, 5, 908},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSymbol(parsed, tt.col, tt.num); got != tt.want {
				t.Errorf("hasSymbol subtest %s wanted %v,gut %v", tt.name, got, tt.want)
			}
		})
	}

}

func Test_parseLine(t *testing.T) {
	util.SetDebuglvl(10)
	tests := []struct {
		name  string
		input []byte
		want  []number
	}{
		{
			name:  "Testing parseling of line ..35..633.",
			input: []byte{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
			want: []number{
				{
					nr:    35,
					start: 2,
					end:   4,
				},
				{
					nr:    633,
					start: 6,
					end:   9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := parseLine(tt.input)
			same := true
			for i := 0; i < len(ws); i++ {
				if ws[i].nr != tt.want[i].nr ||
					ws[i].start != tt.want[i].start ||
					ws[i].end != tt.want[i].end {
					same = false
				}

			}

			if !same {
				t.Errorf("ParseLine = %v, want %v", ws, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	util.SetDebuglvl(0)
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  4361,
		},
		{
			name:  "actual",
			input: input,
			want:  0,
		},
		/*
			wrong: 101491099 to high
			wrong: 538279
			wrong: 538121 to low
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  0,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
