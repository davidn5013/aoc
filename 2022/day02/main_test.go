package main

import (
	"testing"
)

func TestSol1_1(t *testing.T) {
	ans := sol1_1([]string{
		"A Y",
		"B X",
		"C Z",
	})
	if ans != 15 {
		t.Errorf("Sol1_1 =%d ;want 15", ans)
	}
}

func BenchmarkSol1_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1_1([]string{
			"A Y",
			"B X",
			"C Z",
		})
	}
}

func TestSol1_2(t *testing.T) {
	ans := sol1_2([]string{
		"A Y",
		"B X",
		"C Z",
	})
	if ans != 15 {
		t.Errorf("Sol1_2 =%d ;want 15", ans)
	}
}

func BenchmarkSol1_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1_2([]string{
			"A Y",
			"B X",
			"C Z",
		})
	}
}

func TestSol2(t *testing.T) {
	ans := sol2([]string{
		"A Y",
		"B X",
		"C Z",
	})
	if ans != 12 {
		t.Errorf("Sol2 =%d ;want 15", ans)
	}
}

func BenchmarkSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol2([]string{
			"A Y",
			"B X",
			"C Z",
		})
	}
}
