package main

import (
	"testing"
)

var example string = `123 -> a
456 -> b
a AND b -> d
a OR b -> e
a LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT a -> h
NOT b -> i`

func Test_someAssemblyRequired(t *testing.T) {
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
		// 	name:  "actual",
		// 	input: util.ReadFile("input.txt"),
		// 	want:  3376,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := someAssemblyRequired(tt.input, 1); got != tt.want {
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
		// 	input: util.ReadFile("input.txt"),
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := someAssemblyRequired(tt.input, 2); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
