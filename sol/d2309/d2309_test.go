package d2309

import (
	"testing"

	"github.com/davidn5013/aoc/util"
)

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

var input = util.ShardInputFile(util.PathInputShared("2023", "9", "../../input", "input.txt"))

func Test_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  114,
		},
		{
			name:  "actual",
			input: input,
			want:  1842168671,
		},
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
	util.SetDebuglvl(10)
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  2,
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
