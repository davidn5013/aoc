package main

import "testing"

var ( // Test is used here and YES I'am testing two times...
	Test = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
)

func TestMostSol1(t *testing.T) {
	ans := mostSol(Test, 1)
	if ans != 2 {
		t.Errorf("Sol1 =%d ;want 2", ans)
	}
}

func TestMostSol2(t *testing.T) {
	ans := mostSol(Test, 2)
	if ans != 4 {
		t.Errorf("Sol2 =%d ;want 4", ans)
	}
}

func BenchmarkMostSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mostSol(Test, 1)
	}
}

func BenchmarkMostSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mostSol(Test, 2)
	}
}
