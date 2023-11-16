package main

import (
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "turn on 0,0 to 999,999",
			input: "turn on 0,0 through 999,999",
			want:  1000000,
		},
		{
			name:  "toggle 0,0 to 999,999",
			input: "toggle 0,0 through 999,999",
			want:  1000000,
		},
		{
			name:  "actual",
			input: input,
			want:  569999,
		},
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
			name:  "actual",
			input: input,
			want:  17836115,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
