package main

import (
	"testing"
)

var example = ``

func Test_Squaref(t *testing.T) {
	want := 58
	gut := squaref(2, 3, 4)
	if want != gut {
		t.Errorf("squaref2 wanted %d gut %d", want, gut)
	}
	want = 43
	gut = squaref(1, 1, 10)
	if want != gut {
		t.Errorf("sqareef() 2 failed wanted %d gut %d", want, gut)
	}
}

func Test_Ribbon(t *testing.T) {
	want := 34
	gut := ribbonf(2, 3, 4)
	if want != gut {
		t.Errorf("ribbonf failed wanted %d gut %d", want, gut)
	}
}

func Test_part1(t *testing.T) {
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
