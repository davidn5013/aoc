package main

import (
	"testing"
)

func Test_wordToNum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  byte
	}{
		{
			name:  "example1-eightwothree",
			input: "eeightwothre",
			want:  '8',
		},
		{
			name:  "example2-eighwothree",
			input: "eighwothree",
			want:  '3',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordToNum(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
		{
			name:  "actual",
			input: input,
			want:  54450,
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
			name: "example",
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
		{
			name:  "actual",
			input: input,
			want:  54265,
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
