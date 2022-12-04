package main

import "testing"

var (
	test = []string{
		"",
	}
)

func TestSol1(t *testing.T) {
	ans := sol1(test)
	if ans != -1 {
		t.Errorf("Sol1_1 =%d ;want XXX", ans)
	}
}

func TestSol2(t *testing.T) {
	ans := sol2(test)
	if ans != -1 {
		t.Errorf("Sol1_1 =%d ;want XXX", ans)
	}
}

func BenchmarkSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1(test)
	}
}

func BenchmarkSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol2(test)
	}
}
