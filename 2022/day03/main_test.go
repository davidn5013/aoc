package main

import "testing"

var (
	test = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
)

func TestSol1(t *testing.T) {
	ans := sol1(test)
	if ans != 157 {
		t.Errorf("Sol1_1 =%d ;want 157", ans)
	}
}

func BenchmarkSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1(test)
	}
}

func TestSol2(t *testing.T) {
	ans := sol2(test)
	if ans != 70 {
		t.Errorf("Sol1_1 =%d ;want 70", ans)
	}
}

func BenchmarkSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol2(test)
	}
}
