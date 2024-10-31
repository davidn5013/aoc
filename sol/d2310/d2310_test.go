package d2310

import (
	"testing"

	"github.com/davidn5013/aoc/util"
)

var example = `.....
.S-7.
.|.|.
.L-J.
.....`

var input = util.ShardInputFile(util.PathInputShared("2023", "10", "../../input", "input.txt"))

func Test_Part1(t *testing.T) {
	util.SetDebuglvl(10)
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  5,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
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
			if got := Part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
