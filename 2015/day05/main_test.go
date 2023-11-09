package main

import (
	"testing"
)

var example = ``

func Test_dubbelLetters(t *testing.T) {
	testing := "abccdea"
	gut := dubbelLetters(testing)
	want := true
	if want != gut {
		t.Errorf("wanted %t ; gute %t ", want, gut)
	}
}

func Test_vowelcount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		/*
			Santa needs help figuring out which strings in his text file are naughty or nice.
			A nice string is one with all of the following properties:
			It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
			It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
			It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
			For example:
			ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
			aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
			jchzalrnumimnmhp is naughty because it has no double letter.
			haegwjzuvuyypxyu is naughty because it contains the string xy.
			dvszwmarrgswjxmb is naughty because it contains only one vowel.
			How many strings are nice?
		*/
		{
			name:  "example vowels",
			input: "aeiouaeiouaeiou",
			want:  true,
		},
		{
			name:  "example dubbel",
			input: "aabbccdd",
			want:  false,
		},
		{
			name:  "example no double letter",
			input: "jchzalrnumimhp",
			want:  true,
		},
		{
			name:  "example has xy",
			input: "haegwjzuvuyypabui",
			want:  true,
		},
		{
			name:  "example has xy",
			input: "haegwjzuvuyypxyu",
			want:  true,
		},
		{
			name:  "example not three vowels",
			input: "dvszwmarrgswjxmb",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelcount(tt.input, 3); got != tt.want {
				t.Errorf("vowelcount() = %v, want %v", got, tt.want)
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
			name:  "example vowels",
			input: "ugknbfddgicrmopn",
			want:  1,
		},
		{
			name:  "example dubbel",
			input: "aaa",
			want:  1,
		},
		{
			name:  "example no dubbel",
			input: "jchzalrnumimnmhp",
			want:  0,
		},
		{
			name:  "example has xy",
			input: "haegwjzuvuyypxyu",
			want:  0,
		},
		{
			name:  "example not three vowels",
			input: "dvszwmarrgswjxmb",
			want:  0,
		},
		{
			name:  "actual",
			input: input,
			want:  236,
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
			name:  "three vowels",
			input: "qjhvhtzxzqqjkmpb",
			want:  1,
		},
		{
			name:  "secound nice",
			input: "xxyxx",
			want:  1,
		},
		{
			// no reapting singel letter between them
			name:  "not nice",
			input: "uurcxstgmygtbstg",
			want:  0,
		}, {
			name:  "not nice2",
			input: "ieodomkazucvgmuy",
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
